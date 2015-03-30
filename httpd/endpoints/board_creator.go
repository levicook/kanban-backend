package endpoints

import (
	"io"
	"net/http"

	"github.com/levicook/kanban-backend/models"
	"github.com/levicook/kanban-backend/repos"
)

type boardCreator struct {
	boardRepo repos.BoardRepo
	board     models.Board
}

func (c *boardCreator) notAcceptable(http.Header) bool {
	return false
}

func (c *boardCreator) unauthorized(http.Header) bool {
	return true
}

func (c *boardCreator) processBody(body io.ReadCloser) bool {
	err := readJSON(body, &c.board)
	// todo log err
	return err == nil
}

func (c *boardCreator) forbidden() bool {
	return true
}

func (c *boardCreator) create() (entity interface{}, errors models.Errors) {
	return c.board, c.boardRepo.Create(&c.board)
}
