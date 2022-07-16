package repository

const (
	createUser = "INSERT INTO user (user_id,username,name,password,age) values ($1,$2,$3,$4,$5)"
)
