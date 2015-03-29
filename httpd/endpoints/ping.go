package endpoints

import (
	"fmt"
	"net/http"
)

func Ping(w http.ResponseWriter, _ *http.Request, _ map[string]string) {
	fmt.Fprintf(w, "pong")
}
