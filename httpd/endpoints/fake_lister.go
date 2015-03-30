package endpoints

import (
	"io"
	"net/http"
)

type fakeLister struct {
	notAcceptableStub func(http.Header) bool
	unauthorizedStub  func(http.Header) bool
	badRequestStub    func(io.ReadCloser) bool
	forbiddenStub     func() bool
	listStub          func(routeParams) interface{}
}

func (l fakeLister) notAcceptable(h http.Header) bool { return l.notAcceptableStub(h) }
func (l fakeLister) unauthorized(h http.Header) bool  { return l.unauthorizedStub(h) }
func (l fakeLister) badRequest(rc io.ReadCloser) bool { return l.badRequestStub(rc) }
func (l fakeLister) forbidden() bool                  { return l.forbiddenStub() }
func (l fakeLister) list(p routeParams) interface{}   { return l.listStub(p) }
