package repository

import (
	"database/sql"
	"log"

	"mux/internal/group"
	"mux/internal/group/dto"
	"mux/internal/models"

	"github.com/satori/uuid"
)

type groupRepo struct {
	db     *sql.DB
	logger *log.Logger
}

func NewRepository(db *sql.DB, logger *log.Logger) group.Repository {
	return &groupRepo{db: db, logger: logger}
}

func (r *groupRepo) Create(group *models.Group) error {
	id := uuid.NewV4()

	query, err := r.db.Prepare(createGroupQuery)
	if err != nil {
		return err
	}

	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Stmt(query).Exec(id, group.UserID,
		group.CreatedAt, group.Title, group.Body)
	if err != nil {
		r.logger.Println("doing rollback")
		r.logger.Println(err.Error())
		tx.Rollback()
	} else {
		tx.Commit()
	}

	return nil
}

func (r *groupRepo) JoinRequest(gUser *models.GroupUser) error {
	query, err := r.db.Prepare(createGroupUserQuery)
	if err != nil {
		return err
	}

	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Stmt(query).Exec(gUser.GroupID, gUser.UserID, gUser.CreatedAt, gUser.Invite)
	if err != nil {
		r.logger.Println("doing rollback")
		r.logger.Println(err.Error())
		tx.Rollback()
	} else {
		tx.Commit()
	}

	return nil
}

func (r *groupRepo) GetAllGroups() ([]models.Group, error) {
	var groups []models.Group
	rows, err := r.db.Query(getGroupsQuery)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		var dbGroup models.Group
		err = rows.Scan(&dbGroup.GroupID, &dbGroup.UserID, &dbGroup.CreatedAt, &dbGroup.Title, &dbGroup.Body)
		if err != nil {
			return nil, err
		}
		groups = append(groups, dbGroup)
	}

	err = rows.Err()
	if err != nil {
		r.logger.Println(err)
	}

	return groups, nil
}

func (r *groupRepo) GetRequests(user models.User) ([]dto.ModelJoinRequest, error) {
	var models []dto.ModelJoinRequest
	rows, err := r.db.Query(getUserGroupInvites, user.UserID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		var dbModel dto.ModelJoinRequest
		err = rows.Scan(&dbModel.UserID, &dbModel.Email, &dbModel.FName,
			&dbModel.LName, &dbModel.Avatar, &dbModel.GroupID, &dbModel.Title)
		if err != nil {
			return nil, err
		}
		models = append(models, dbModel)
	}

	err = rows.Err()
	if err != nil {
		r.logger.Println(err)
	}

	return models, nil
}

func (r *groupRepo) CheckGroupByTitle(title string) (bool, error) {
	var groupId string
	row := r.db.QueryRow(getGroupByTitleQuery, title)
	err := row.Scan(&groupId)

	switch err {
	case sql.ErrNoRows:
		return false, nil
	case nil:
		return true, nil
	default:
		return false, err
	}
}

func (r *groupRepo) CheckGroupByID(id uuid.UUID) (bool, error) {
	var groupId string
	row := r.db.QueryRow(getGroupByIDQuery, id)
	err := row.Scan(&groupId)

	switch err {
	case sql.ErrNoRows:
		return false, nil
	case nil:
		return true, nil
	default:
		return false, err
	}
}
