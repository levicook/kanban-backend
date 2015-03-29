package endpoints

import (
	"net/http"

	"github.com/dimfeld/httptreemux"
	"github.com/levicook/todo-api/httpd/send"
	"github.com/levicook/todo-api/httpd/status"
)

func listHandler(newLister listerFactory) httptreemux.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, _ map[string]string) {
		l := newLister()

		if l.notAcceptable(r.Header) {
			send.NotAcceptable(w)
			return
		}

		if l.unauthorized(r.Header) {
			send.Unauthorized(w)
			return
		}

		if l.badRequest(r.Body) {
			send.BadRequest(w)
			return
		}

		if l.forbidden() {
			send.Forbidden(w)
			return
		}

		list := l.list()
		etag := etagFor(list)

		if etag == r.Header.Get("If-None-Match") {
			send.NotModified(w)
			return
		}

		w.Header().Set("ETag", etag)
		send.Json(w, status.OK, list)
	}
}
