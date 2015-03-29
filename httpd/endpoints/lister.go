package endpoints

import (
	"io"
	"net/http"
)

type listerFactory func() lister

type lister interface {
	notAcceptable(http.Header) bool
	unauthorized(http.Header) bool
	badRequest(io.ReadCloser) bool
	forbidden() bool
	list() interface{}
}
