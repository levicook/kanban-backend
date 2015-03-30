package endpoints

import (
	"net/http"

	"github.com/dimfeld/httptreemux"
	"github.com/levicook/kanban-backend/httpd/status"
)

func listHandler(
	newTransaction transactionFactory,
	newLister listerFactory,
) httptreemux.HandlerFunc {

	return func(
		w http.ResponseWriter,
		r *http.Request,
		p map[string]string,
	) {
		var (
			t = newTransaction()
			l = newLister(t)
		)

		if l.notAcceptable(r.Header) {
			panicIf(t.Rollback())
			sendNotAcceptable(w)
			return
		}

		if l.unauthorized(r.Header) {
			panicIf(t.Rollback())
			sendUnauthorized(w)
			return
		}

		if l.badRequest(r.Body) {
			panicIf(t.Rollback())
			sendBadRequest(w)
			return
		}

		if l.forbidden() {
			panicIf(t.Rollback())
			sendForbidden(w)
			return
		}

		list := l.list(p)
		t.Commit()

		etag := etagFor(list)

		if etag == r.Header.Get("If-None-Match") {
			sendNotModified(w)
			return
		}

		w.Header().Set("ETag", etag)
		sendJSON(w, status.OK, list)
	}
}
