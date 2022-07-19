package group

import (
	"mux/internal/models"
)

type Repository interface {
	Create(*models.Group) error
	CheckGroupByTitle(string) (bool, error)
}
