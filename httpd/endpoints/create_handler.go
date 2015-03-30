package endpoints

import (
	"net/http"

	"github.com/dimfeld/httptreemux"
	"github.com/levicook/kanban-backend/httpd/send"
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
			send.NotAcceptable(w)
			return
		}

		if c.unauthorized(r.Header) {
			panicIf(t.Rollback())
			send.Unauthorized(w)
			return
		}

		if !c.processBody(r.Body) {
			panicIf(t.Rollback())
			send.BadRequest(w)
			return
		}

		if c.forbidden() {
			panicIf(t.Rollback())
			send.Forbidden(w)
			return
		}

		entity, errors := c.create()

		if errors.Present() {
			panicIf(t.Rollback())
			send.UnprocessableEntity(w, errors)
			return
		}

		panicIf(t.Commit())
		w.Header().Set("ETag", etagFor(entity))
		send.Json(w, status.Created, entity)
	}
}
