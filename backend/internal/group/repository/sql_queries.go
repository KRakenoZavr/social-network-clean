package repository

const (
	createGroupQuery = `
		INSERT INTO groups 
			(group_id,user_id,created_at,title,body) 
		VALUES ($1,$2,$3,$4,$5);`

	acceptInvite = `
		UPDATE group_user
		SET invite=0
		WHERE group_id=$1 AND user_id=$2;`
	declineInvite = `
		DELETE FROM group_user
		WHERE group_id=$1 AND user_id=$2;`

	getGroupsQuery = `
		SELECT * FROM groups;`
	getGroupByTitleQuery = `
		SELECT group_id FROM groups WHERE title=$1;`
	getGroupByIDQuery = `
		SELECT group_id FROM groups WHERE group_id=$1;`
	checkIfAdmin = `
		SELECT group_id FROM groups WHERE group_id=$1 AND user_id=$2;`
)

const (
	createGroupUserQuery = `
		INSERT INTO group_user 
			(group_id,user_id,created_at,invite) 
		VALUES ($1,$2,$3,$4);`
	createGroupInviteQuery = `
		INSERT INTO group_user
			(group_id,user_id,created_at,invite)
		VALUES ($1,$2,$3,$4);`

	getUserInvites = `
		SELECT gu.group_id,gu.user_id,g.created_at,g.title,g.body FROM group_user gu
		JOIN groups g USING(group_id)
		WHERE gu.user_id=$1 AND gu.invite=1;`
	getUserGroupInvites = `
		SELECT u.user_id,u.email,u.first_name,u.last_name,u.avatar,j.group_id,j.title 
		FROM users u
		JOIN 
			(SELECT gu.group_id,gu.user_id,g.title FROM group_user gu
			JOIN groups g USING(group_id)
			WHERE g.user_id=$1 AND gu.invite=2) j
		USING(user_id);`
)
