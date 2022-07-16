package repository

const (
	createUser = `INSERT INTO users (user_id,email,password,first_name,last_name,
																	date_of_birth,is_private,avatar,nickname,about) 
									values ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)`
)
