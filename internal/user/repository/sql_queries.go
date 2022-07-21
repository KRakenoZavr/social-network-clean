package repository

const (
	createUserQuery = `
		INSERT INTO users 
			(user_id,email,password,first_name,last_name,
			date_of_birth,is_private,avatar,nickname,about) 
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10);`
	getUserByEmailQuery = `
		SELECT * FROM users WHERE email=$1;`
	getUserIDByEmailQuery = `
		SELECT user_id FROM users WHERE email=$1;`
	getUserByIDQuery = `
		SELECT * FROM users WHERE user_id=$1;`
)

const (
	createUserAuthQuery = `
		INSERT INTO user_auth
			(user_id,expires,session)
		VALUES ($1,$2,$3);`
	getUserAuthQuery = `
		SELECT * FROM user_auth WHERE session=$1;`
)
