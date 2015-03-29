package routes

import "github.com/dimfeld/httptreemux"

type route struct {
	name        string
	verb        string
	path        string
	handlerFunc httptreemux.HandlerFunc
}
