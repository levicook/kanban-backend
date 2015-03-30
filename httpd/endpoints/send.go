package endpoints

import (
	"encoding/json"
	"net/http"

	"github.com/levicook/kanban-backend/httpd/status"
	"github.com/levicook/kanban-backend/models"
	"github.com/levicook/slog"
)

type errDoc struct {
	Code int    `json:"code"`
	Text string `json:"text"`
}

func sendBadRequest(w http.ResponseWriter) {
	sendJSON(w, status.BadRequest, errDoc{
		Code: status.BadRequest,
		Text: "Bad Request",
	})
}

func sendForbidden(w http.ResponseWriter) {
	sendJSON(w, status.Forbidden, errDoc{
		Code: status.Forbidden,
		Text: "Forbidden",
	})
}

func sendJSON(w http.ResponseWriter, status int, v interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(status)
	slog.PanicIf(json.NewEncoder(w).Encode(v))
}

func sendNotAcceptable(w http.ResponseWriter) {
	sendJSON(w, status.NotAcceptable, errDoc{
		Code: status.NotAcceptable,
		Text: "Not Acceptable",
	})
}

func sendNotImplemented(w http.ResponseWriter) {
	sendJSON(w, status.NotImplemented, errDoc{
		Code: status.NotImplemented,
		Text: "Not Implemented",
	})
}

func sendNotModified(w http.ResponseWriter) {
	http.Error(w, "", status.NotModified)
}

func sendUnauthorized(w http.ResponseWriter) {
	sendJSON(w, status.Unauthorized, errDoc{
		Code: status.Unauthorized,
		Text: "Unauthorized",
	})
}

func sendUnprocessableEntity(w http.ResponseWriter, errors models.Errors) {
	sendJSON(w, status.UnprocessableEntity, errors)
}
