package endpoints

import (
	"io"
	"net/http"

	"github.com/levicook/kanban-backend/models"
)

type creator interface {
	notAcceptable(http.Header) bool
	unauthorized(http.Header) bool
	processBody(io.ReadCloser) bool
	forbidden() bool
	create() (interface{}, models.Errors)
}
