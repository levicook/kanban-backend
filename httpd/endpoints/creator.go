package endpoints

import (
	"io"
	"net/http"

	"github.com/levicook/todo-api/models"
)

type creatorFactory func() creator

type creator interface {
	notAcceptable(http.Header) bool
	unauthorized(http.Header) bool
	badRequest(io.ReadCloser) bool
	forbidden() bool
	create() (interface{}, models.Errors)
}
