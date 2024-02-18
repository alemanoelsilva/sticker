package validators

type SignUp struct {
	ID       int    `validate:"optional"`
	Name     string `validate:"required"`
	Email    string `validate:"required,email"`
	Password string `validate:"required"`
}

type SignIn struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required"`
}
