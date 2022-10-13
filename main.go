package main

import (
	"edufunds/config"
	"edufunds/controllers"
	"edufunds/repository"
	"edufunds/services"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)


var (
		//Database
		db *gorm.DB = config.SetUpDatabaseConnection()
		//Repositories
		userRepository repository.UserRepository = repository.NewUserRepository(db)

		//Services
		jwtService  services.JWTService  = services.NewJWTService()
		authService services.AuthService = services.NewAuthService(userRepository)
	
		//Controllers
		authController controllers.AuthController = controllers.NewAuthController(authService, jwtService)
		
)

func main()  {
	defer config.CloseDatabaseConnection(db)
	server := gin.New()
	authRoutes := server.Group("/api/auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}
	port := os.Getenv("APP_PORT")
	server.Run(fmt.Sprintf(":%s", port))
}