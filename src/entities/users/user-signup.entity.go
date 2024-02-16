package users_entity

type UserSignup struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
