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
	getGroupByTitleQuery = `
		SELECT group_id FROM groups WHERE title=$1;`
	getGroupByIDQuery = `
		SELECT group_id FROM groups WHERE group_id=$1;`
)
