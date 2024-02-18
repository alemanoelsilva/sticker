package query

const (
	AddUserQuery        = "INSERT INTO users (name, email, password) VALUES (:name, :email, :password)"
	GetUserByIdQuery    = "SELECT * FROM users WHERE id = '%d' LIMIT 1"
	GetUserByEmailQuery = "SELECT id, name, email, password FROM users WHERE email = '%s' LIMIT 1"
)
