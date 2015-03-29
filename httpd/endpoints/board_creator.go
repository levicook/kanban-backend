package endpoints

import (
	"io"
	"net/http"

	"github.com/levicook/todo-api/models"
	"github.com/levicook/todo-api/repos"
)

type boardCreator struct {
	boardRepo repos.BoardRepo
	board     models.Board
}

func (c *boardCreator) notAcceptable(http.Header) bool {
	return true
}

func (c *boardCreator) unauthorized(http.Header) bool {
	return true
}

func (c *boardCreator) badRequest(io.ReadCloser) bool {
	return true
}

func (c *boardCreator) forbidden() bool {
	return true
}

func (c *boardCreator) create() (entity interface{}, errors models.Errors) {
	return
}
