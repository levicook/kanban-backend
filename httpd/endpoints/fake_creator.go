package endpoints

import (
	"io"
	"net/http"

	"github.com/levicook/todo-api/models"
)

type fakeCreator struct {
	notAcceptableStub func(http.Header) bool
	unauthorizedStub  func(http.Header) bool
	badRequestStub    func(io.ReadCloser) bool
	forbiddenStub     func() bool
	createStub        func() (interface{}, models.Errors)
}

func (l fakeCreator) notAcceptable(h http.Header) bool     { return l.notAcceptableStub(h) }
func (l fakeCreator) unauthorized(h http.Header) bool      { return l.unauthorizedStub(h) }
func (l fakeCreator) badRequest(rc io.ReadCloser) bool     { return l.badRequestStub(rc) }
func (l fakeCreator) forbidden() bool                      { return l.forbiddenStub() }
func (l fakeCreator) create() (interface{}, models.Errors) { return l.createStub() }
