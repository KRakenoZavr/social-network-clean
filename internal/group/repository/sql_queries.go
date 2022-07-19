package repository

const (
	createGroupQuery = `
		INSERT INTO groups 
			(group_id,user_id,created_at,title,body) 
		VALUES ($1,$2,$3,$4,$5);`
	getGroupByTitleQuery = `
		SELECT group_id FROM groups WHERE title=$1`
)
