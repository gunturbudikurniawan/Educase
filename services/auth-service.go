package services

import (
	"edufunds/models"
	"edufunds/repository"
	"edufunds/validations"
	"log"

	"github.com/mashingan/smapping"
	"golang.org/x/crypto/bcrypt"
)

// AuthService is a contract about something that this service can do
type AuthService interface {
	VerifyCredential(username string, password string) interface{}
	CreateUser(user validations.RegisterValidation) models.User
	FindByEmail(username string) models.User
	IsDuplicateEmail(username string) bool
}

type authService struct {
	userRepository repository.UserRepository
}

// NewAuthService creates a new instance of AuthService
func NewAuthService(userRep repository.UserRepository) AuthService {
	return &authService{
		userRepository: userRep,
	}
}

func (service *authService) VerifyCredential(username string, password string) interface{} {
	res := service.userRepository.VerifyCredential(username, password)
	if v, ok := res.(models.User); ok {
		comparedPassword := comparePassword(v.Password, []byte(password))
		if v.Username == username && comparedPassword {
			return res
		}
		return false
	}
	return false
}

func (service *authService) CreateUser(user validations.RegisterValidation) models.User {
	userToCreate := models.User{}
	err := smapping.FillStruct(&userToCreate, smapping.MapFields(&user))
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}
	res := service.userRepository.InsertUser(userToCreate)
	return res
}

func (service *authService) FindByEmail(username string) models.User {
	return service.userRepository.FindByEmail(username)
}

func (service *authService) IsDuplicateEmail(username string) bool {
	res := service.userRepository.IsDuplicateEmail(username)
	return !(res.Error == nil)
}

func comparePassword(hashedPwd string, plainPassword []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPassword)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
