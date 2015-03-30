package endpoints

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/levicook/kanban-backend/httpd/status"
	"github.com/levicook/kanban-backend/models"
	"github.com/levicook/kanban-backend/repos"
	"github.com/stretchr/testify/assert"
)

func Test_createHandler(t *testing.T) {

	testNotAcceptable := func() {
		var (
			w = httptest.NewRecorder()
			r http.Request
			p map[string]string
			m = mockTransaction{}
			h = createHandler(
				func() repos.Transaction { return &m },
				func(repos.Transaction) creator {
					return fakeCreator{
						notAcceptableStub: func(http.Header) bool { return true },
					}
				},
			)
		)

		h(w, &r, p)

		assert.Equal(t, 1, m.rollbackCalls)
		assert.Equal(t, 0, m.commitCalls)

		assert.Equal(t, status.NotAcceptable, w.Code)
		assert.Equal(t, "application/json; charset=UTF-8", w.Header().Get("Content-Type"))
		assert.Equal(t, "", w.Header().Get("ETag"))
	}

	testUnauthorized := func() {
		var (
			w = httptest.NewRecorder()
			r http.Request
			p map[string]string
			m = mockTransaction{}
			h = createHandler(
				func() repos.Transaction { return &m },
				func(repos.Transaction) creator {
					return fakeCreator{
						notAcceptableStub: func(http.Header) bool { return false },
						unauthorizedStub:  func(http.Header) bool { return true },
					}
				},
			)
		)

		h(w, &r, p)

		assert.Equal(t, 1, m.rollbackCalls)
		assert.Equal(t, 0, m.commitCalls)

		assert.Equal(t, status.Unauthorized, w.Code)
		assert.Equal(t, "application/json; charset=UTF-8", w.Header().Get("Content-Type"))
		assert.Equal(t, "", w.Header().Get("ETag"))
	}

	testBadRequest := func() {
		var (
			w = httptest.NewRecorder()
			r http.Request
			p map[string]string
			m = mockTransaction{}
			h = createHandler(
				func() repos.Transaction { return &m },
				func(repos.Transaction) creator {
					return fakeCreator{
						notAcceptableStub: func(http.Header) bool { return false },
						unauthorizedStub:  func(http.Header) bool { return false },
						processBodyStub:   func(io.ReadCloser) bool { return false },
					}
				},
			)
		)

		h(w, &r, p)

		assert.Equal(t, 1, m.rollbackCalls)
		assert.Equal(t, 0, m.commitCalls)

		assert.Equal(t, status.BadRequest, w.Code)
		assert.Equal(t, "application/json; charset=UTF-8", w.Header().Get("Content-Type"))
		assert.Equal(t, "", w.Header().Get("ETag"))
	}

	testForbidden := func() {
		var (
			w = httptest.NewRecorder()
			r http.Request
			p map[string]string
			m = mockTransaction{}
			h = createHandler(
				func() repos.Transaction { return &m },
				func(repos.Transaction) creator {
					return fakeCreator{
						notAcceptableStub: func(http.Header) bool { return false },
						unauthorizedStub:  func(http.Header) bool { return false },
						processBodyStub:   func(io.ReadCloser) bool { return true },
						forbiddenStub:     func() bool { return true },
					}
				},
			)
		)

		h(w, &r, p)

		assert.Equal(t, 1, m.rollbackCalls)
		assert.Equal(t, 0, m.commitCalls)

		assert.Equal(t, status.Forbidden, w.Code)
		assert.Equal(t, "application/json; charset=UTF-8", w.Header().Get("Content-Type"))
		assert.Equal(t, "", w.Header().Get("ETag"))
	}

	testUnprocessableEntity := func() {
		var (
			w = httptest.NewRecorder()
			r http.Request
			p map[string]string
			m = mockTransaction{}
			h = createHandler(
				func() repos.Transaction { return &m },
				func(repos.Transaction) creator {
					return fakeCreator{
						notAcceptableStub: func(http.Header) bool { return false },
						unauthorizedStub:  func(http.Header) bool { return false },
						processBodyStub:   func(io.ReadCloser) bool { return true },
						forbiddenStub:     func() bool { return false },
						createStub: func() (interface{}, models.Errors) {
							return nil, models.Errors{"base": "is invalid"}
						},
					}
				},
			)
		)

		h(w, &r, p)

		assert.Equal(t, 1, m.rollbackCalls)
		assert.Equal(t, 0, m.commitCalls)

		assert.Equal(t, status.UnprocessableEntity, w.Code)
		assert.Equal(t, "application/json; charset=UTF-8", w.Header().Get("Content-Type"))
		assert.Equal(t, "", w.Header().Get("ETag"))

		var e models.Errors
		assert.Nil(t, json.Unmarshal(w.Body.Bytes(), &e))
		assert.Equal(t, models.Errors{"base": "is invalid"}, e)
	}

	testCreated := func() {
		var (
			w = httptest.NewRecorder()
			r http.Request
			p map[string]string
			m = mockTransaction{}
			h = createHandler(
				func() repos.Transaction { return &m },
				func(repos.Transaction) creator {
					return fakeCreator{
						notAcceptableStub: func(http.Header) bool { return false },
						unauthorizedStub:  func(http.Header) bool { return false },
						processBodyStub:   func(io.ReadCloser) bool { return true },
						forbiddenStub:     func() bool { return false },
						createStub: func() (interface{}, models.Errors) {
							return "hello", models.Errors{}
						},
					}
				},
			)
		)

		h(w, &r, p)

		assert.Equal(t, 0, m.rollbackCalls)
		assert.Equal(t, 1, m.commitCalls)

		assert.Equal(t, status.Created, w.Code)
		assert.Equal(t, "application/json; charset=UTF-8", w.Header().Get("Content-Type"))
		assert.Equal(t, "77f088caf896bc2be55af5cdf8e860ae", w.Header().Get("ETag"))

		var b string
		assert.Nil(t, json.Unmarshal(w.Body.Bytes(), &b))
		assert.Equal(t, "hello", b)
	}

	testNotAcceptable()
	testUnauthorized()
	testBadRequest()
	testForbidden()
	testUnprocessableEntity()
	testCreated()
}
