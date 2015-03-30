package endpoints

import (
	"net/http"

	"github.com/dimfeld/httptreemux"
	"github.com/levicook/kanban-backend/httpd/status"
)

func createHandler(
	newTransaction transactionFactory,
	newCreator creatorFactory,
) httptreemux.HandlerFunc {

	return func(
		w http.ResponseWriter,
		r *http.Request,
		_ map[string]string,
	) {
		var (
			t = newTransaction()
			c = newCreator(t)
		)

		if c.notAcceptable(r.Header) {
			panicIf(t.Rollback())
			sendNotAcceptable(w)
			return
		}

		if c.unauthorized(r.Header) {
			panicIf(t.Rollback())
			sendUnauthorized(w)
			return
		}

		if !c.processBody(r.Body) {
			panicIf(t.Rollback())
			sendBadRequest(w)
			return
		}

		if c.forbidden() {
			panicIf(t.Rollback())
			sendForbidden(w)
			return
		}

		entity, errors := c.create()

		if errors.Present() {
			panicIf(t.Rollback())
			sendUnprocessableEntity(w, errors)
			return
		}

		panicIf(t.Commit())
		w.Header().Set("ETag", etagFor(entity))
		sendJSON(w, status.Created, entity)
	}
}
