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
	createFollowQuery = `
		INSERT INTO user_follow
			(user_id1,user_id2,created_at,invite)
		VALUES ($1,$2,$3,$4)`

	getInvites = `
		SELECT u.user_id,u.email,u.first_name,u.last_name,u.date_of_birth,u.is_private,u.avatar FROM user_follow uf
		JOIN users u ON uf.user_id1=u.user_id
		WHERE uf.user_id2=$1;`
	acceptInvite = `
		UPDATE user_follow
		SET invite=0
		WHERE user_id1=$1 AND user_id2=$2;`
	declineInvite = `
		DELETE FROM user_follow
		WHERE user_id1=$1 AND user_id2=$2;`
)

const (
	createUserAuthQuery = `
		INSERT INTO user_auth
			(user_id,expires,session)
		VALUES ($1,$2,$3);`
	getUserAuthQuery = `
		SELECT * FROM user_auth WHERE session=$1;`
)
