package endpoints

import "net/http"

func Panic(_ http.ResponseWriter, _ *http.Request, _ map[string]string) {
	panic("panic - endpoints.Panic panicked on purpose")
}
