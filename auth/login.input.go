package auth

type LoginInput struct {
	Email    string `form:"email"`
	Password string `form:"password"`
}
