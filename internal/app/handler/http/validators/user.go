package validators

type SignUp struct {
	Name     string `validate:"required"`
	Email    string `validate:"required,email"`
	Password string `validate:"required"`
}

type SignIn struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required"`
}
