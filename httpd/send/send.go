package send

import (
	"encoding/json"
	"net/http"

	"github.com/levicook/slog"
	"github.com/levicook/todo-api/httpd/status"
	"github.com/levicook/todo-api/models"
)

type errDoc struct {
	Code int    `json:"code"`
	Text string `json:"text"`
}

func BadRequest(w http.ResponseWriter) {
	Json(w, status.BadRequest, errDoc{
		Code: status.BadRequest,
		Text: "Bad Request",
	})
}

func Forbidden(w http.ResponseWriter) {
	Json(w, status.Forbidden, errDoc{
		Code: status.Forbidden,
		Text: "Forbidden",
	})
}

func Json(w http.ResponseWriter, status int, v interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(status)
	slog.PanicIf(json.NewEncoder(w).Encode(v))
}

func NotAcceptable(w http.ResponseWriter) {
	Json(w, status.NotAcceptable, errDoc{
		Code: status.NotAcceptable,
		Text: "Not Acceptable",
	})
}

func NotImplemented(w http.ResponseWriter) {
	Json(w, status.NotImplemented, errDoc{
		Code: status.NotImplemented,
		Text: "Not Implemented",
	})
}

func NotModified(w http.ResponseWriter) {
	http.Error(w, "", status.NotModified)
}

func Unauthorized(w http.ResponseWriter) {
	Json(w, status.Unauthorized, errDoc{
		Code: status.Unauthorized,
		Text: "Unauthorized",
	})
}

func UnprocessableEntity(w http.ResponseWriter, errors models.Errors) {
	Json(w, status.UnprocessableEntity, errors)
}
