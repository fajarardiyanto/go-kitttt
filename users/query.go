package users

var (
	getUserQuery = `SELECT * FROM users WHERE id=$1`
	createUserQuery = `INSERT INTO users (id, name, email, username, password) VALUES ($1, $2, $3, $4, $5)`
)