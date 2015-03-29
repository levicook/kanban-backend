package routes

import (
	"net/http"

	"github.com/dimfeld/httptreemux"
)

type routeSet []route

func (rs routeSet) Handler() http.Handler {
	return rs.treeMux()
}

func (rs routeSet) treeMux() *httptreemux.TreeMux {
	mux := httptreemux.New()

	for _, route := range rs {
		mux.Handle(route.verb, route.path, func(
			w http.ResponseWriter,
			r *http.Request,
			params map[string]string,
		) {

			route.handlerFunc(w, r, params)

		})
	}

	return mux
}
