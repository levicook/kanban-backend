package endpoints

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/levicook/todo-api/httpd/status"
	"github.com/stretchr/testify/assert"
)

func Test_listHandler(t *testing.T) {

	testNotAcceptable := func() {
		var (
			w = httptest.NewRecorder()
			r http.Request
			p map[string]string
			h = listHandler(func() lister {
				return fakeLister{
					notAcceptableStub: func(http.Header) bool { return true },
				}
			})
		)

		h(w, &r, p)

		assert.Equal(t, status.NotAcceptable, w.Code)
		assert.Equal(t, "application/json; charset=UTF-8", w.Header().Get("Content-Type"))
		assert.Equal(t, "", w.Header().Get("ETag"))
	}

	testUnauthorized := func() {
		var (
			w = httptest.NewRecorder()
			r http.Request
			p map[string]string
			h = listHandler(func() lister {
				return fakeLister{
					notAcceptableStub: func(http.Header) bool { return false },
					unauthorizedStub:  func(http.Header) bool { return true },
				}
			})
		)

		h(w, &r, p)

		assert.Equal(t, status.Unauthorized, w.Code)
		assert.Equal(t, "application/json; charset=UTF-8", w.Header().Get("Content-Type"))
		assert.Equal(t, "", w.Header().Get("ETag"))
	}

	testBadRequest := func() {
		var (
			w = httptest.NewRecorder()
			r http.Request
			p map[string]string
			h = listHandler(func() lister {
				return fakeLister{
					notAcceptableStub: func(http.Header) bool { return false },
					unauthorizedStub:  func(http.Header) bool { return false },
					badRequestStub:    func(io.ReadCloser) bool { return true },
				}
			})
		)

		h(w, &r, p)

		assert.Equal(t, status.BadRequest, w.Code)
		assert.Equal(t, "application/json; charset=UTF-8", w.Header().Get("Content-Type"))
		assert.Equal(t, "", w.Header().Get("ETag"))
	}

	testForbidden := func() {
		var (
			w = httptest.NewRecorder()
			r http.Request
			p map[string]string
			h = listHandler(func() lister {
				return fakeLister{
					notAcceptableStub: func(http.Header) bool { return false },
					unauthorizedStub:  func(http.Header) bool { return false },
					badRequestStub:    func(io.ReadCloser) bool { return false },
					forbiddenStub:     func() bool { return true },
				}
			})
		)

		h(w, &r, p)

		assert.Equal(t, status.Forbidden, w.Code)
		assert.Equal(t, "application/json; charset=UTF-8", w.Header().Get("Content-Type"))
		assert.Equal(t, "", w.Header().Get("ETag"))
	}

	testNotModified := func() {
		var (
			w = httptest.NewRecorder()
			r http.Request
			p map[string]string
			h = listHandler(func() lister {
				return fakeLister{
					notAcceptableStub: func(http.Header) bool { return false },
					unauthorizedStub:  func(http.Header) bool { return false },
					badRequestStub:    func(io.ReadCloser) bool { return false },
					forbiddenStub:     func() bool { return false },
					listStub:          func() interface{} { return []string{} },
				}
			})
		)

		r.Header = http.Header{}
		r.Header.Set("If-None-Match", "58e0494c51d30eb3494f7c9198986bb9")

		h(w, &r, p)

		assert.Equal(t, status.NotModified, w.Code)
		assert.Equal(t, "text/plain; charset=utf-8", w.Header().Get("Content-Type"))
		assert.Equal(t, "", w.Header().Get("ETag"))
	}

	testOK := func() {
		var (
			w = httptest.NewRecorder()
			r http.Request
			p map[string]string
			h = listHandler(func() lister {
				return fakeLister{
					notAcceptableStub: func(http.Header) bool { return false },
					unauthorizedStub:  func(http.Header) bool { return false },
					badRequestStub:    func(io.ReadCloser) bool { return false },
					forbiddenStub:     func() bool { return false },
					listStub:          func() interface{} { return []string{"hello"} },
				}
			})
		)

		h(w, &r, p)

		assert.Equal(t, status.OK, w.Code)
		assert.Equal(t, "application/json; charset=UTF-8", w.Header().Get("Content-Type"))
		assert.Equal(t, "ce7a4e769b92caf9f8b46c84625ae0fc", w.Header().Get("ETag"))

		var b []string
		assert.Nil(t, json.Unmarshal(w.Body.Bytes(), &b))
		assert.Equal(t, []string{"hello"}, b)
	}

	testNotAcceptable()
	testUnauthorized()
	testBadRequest()
	testForbidden()
	testNotModified()
	testOK()
}
