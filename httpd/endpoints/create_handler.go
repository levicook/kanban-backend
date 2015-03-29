package endpoints

import (
	"net/http"

	"github.com/dimfeld/httptreemux"
	"github.com/levicook/todo-api/httpd/send"
	"github.com/levicook/todo-api/httpd/status"
)

func createHandler(newCreator creatorFactory) httptreemux.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, _ map[string]string) {
		c := newCreator()

		if c.notAcceptable(r.Header) {
			send.NotAcceptable(w)
			return
		}

		if c.unauthorized(r.Header) {
			send.Unauthorized(w)
			return
		}

		if c.badRequest(r.Body) {
			send.BadRequest(w)
			return
		}

		if c.forbidden() {
			send.Forbidden(w)
			return
		}

		entity, errors := c.create()

		if errors.Present() {
			send.UnprocessableEntity(w, errors)
			return
		}

		w.Header().Set("ETag", etagFor(entity))
		send.Json(w, status.Created, entity)
	}
}
