package endpoints

import (
	"io"
	"net/http"
)

type lister interface {
	notAcceptable(http.Header) bool
	unauthorized(http.Header) bool
	badRequest(io.ReadCloser) bool
	forbidden() bool
	list(routeParams) interface{}
}
