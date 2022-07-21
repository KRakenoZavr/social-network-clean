package repository

const (
	createGroupQuery = `
		INSERT INTO groups 
			(group_id,user_id,created_at,title,body) 
		VALUES ($1,$2,$3,$4,$5);`
	createGroupUserQuery = `
		INSERT INTO group_user 
			(group_id,user_id,created_at,invite) 
		VALUES ($1,$2,$3,$4);`

	getGroupsQuery = `
		SELECT * FROM groups;`

	getUserGroupInvites = `
		SELECT u.user_id,u.email,u.first_name,u.last_name,u.avatar FROM users u
		JOIN 
			(SELECT gu.id,gu.group_id, gu.user_id,gu.created_at,gu.invite FROM group_user gu
			JOIN groups g USING(group_id)
			WHERE g.user_id=$1 AND gu.invite=2) 
		USING(user_id);`

	getGroupByTitleQuery = `
		SELECT group_id FROM groups WHERE title=$1;`
	getGroupByIDQuery = `
		SELECT group_id FROM groups WHERE group_id=$1;`
)
