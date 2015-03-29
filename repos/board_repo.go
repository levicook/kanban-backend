package repos

import "github.com/levicook/todo-api/models"

type BoardRepo interface {
	Create(*models.Board) models.Errors
}
