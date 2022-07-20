package repository

import (
	"database/sql"
	"log"

	"mux/internal/group"
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
