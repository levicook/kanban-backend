package repos

import (
	"github.com/jmoiron/sqlx"
	"github.com/levicook/kanban-backend/models"
)

type BoardRepo interface {
	CanCreate(models.Board, models.Identity) bool
	CanUpdate(models.BoardId, models.Identity) bool

	OneById(models.BoardId) models.Errors
	Create(*models.Board) models.Errors
}

type boardRepo struct{ *sqlx.Tx }

func NewBoardRepo(t Transaction) BoardRepo {
	return newBoardRepo(t.(*sqlx.Tx))
}

func newBoardRepo(tx *sqlx.Tx) BoardRepo {
	return boardRepo{tx}
}

func (r boardRepo) CanCreate(m models.Board, as models.Identity) bool {
	return false
}

func (r boardRepo) CanUpdate(m models.Board, as models.Identity) bool {
	return false
}

func (r boardRepo) Create(*models.Board) models.Errors {
	return noErrors
}
