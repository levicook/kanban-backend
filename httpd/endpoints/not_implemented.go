package endpoints

import "net/http"

func notImplemented(w http.ResponseWriter, _ *http.Request, _ map[string]string) {
	sendNotImplemented(w)
}
