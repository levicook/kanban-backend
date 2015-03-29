package endpoints

import (
	"net/http"

	"github.com/levicook/todo-api/httpd/send"
)

func notImplemented(w http.ResponseWriter, _ *http.Request, _ map[string]string) {
	send.NotImplemented(w)
}
