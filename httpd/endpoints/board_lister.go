package endpoints

import (
	"io"
	"net/http"

	"github.com/levicook/kanban-backend/models"
	"github.com/levicook/kanban-backend/repos"
)

type boardLister struct {
	boardRepo repos.BoardRepo
	boards    models.Boards
}

func (c *boardLister) notAcceptable(http.Header) bool {
	return true
}

func (c *boardLister) unauthorized(http.Header) bool {
	return true
}

func (c *boardLister) badRequest(io.ReadCloser) bool {
	return true
}

func (c *boardLister) forbidden() bool {
	return true
}

func (c *boardLister) list() interface{} {
	return nil
}
