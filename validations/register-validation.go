package validations

// RegisterValidation is used when client post from /register url
type RegisterValidation struct {
	Username    string `json:"username" form:"username" binding:"required,email"`
	Password string `json:"password" form:"password" binding:"required,min=12"`
}
