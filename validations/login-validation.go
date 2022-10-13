package validations

// LoginValidation is a model that used by client when POST from /login url
type LoginValidation struct {
	Username    string `json:"username" form:"username" binding:"required,email"`
	Password string `json:"password" form:"password" binding:"required,min=12"`
}
