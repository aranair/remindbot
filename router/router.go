package router

import (
	"net/http"

	"github.com/gorilla/context"
	"github.com/julienschmidt/httprouter"
)

type router struct {
	*httprouter.Router
}

func New() *router {
	return &router{httprouter.New()}
}

func (r *router) OPTIONS(path string, h http.Handler) {
	r.Handle("OPTIONS", path, wrapHandler(h))
}

func (r *router) GET(path string, h http.Handler) {
	r.Handle("GET", path, wrapHandler(h))
}

func (r *router) POST(path string, h http.Handler) {
	r.Handle("POST", path, wrapHandler(h))
}

func (r *router) PUT(path string, h http.Handler) {
	r.Handle("PUT", path, wrapHandler(h))
}

func (r *router) PATCH(path string, h http.Handler) {
	r.Handle("PATCH", path, wrapHandler(h))
}

func (r *router) DELETE(path string, h http.Handler) {
	r.Handle("DELETE", path, wrapHandler(h))
}

func (r *router) HEAD(path string, h http.Handler) {
	r.Handle("HEAD", path, wrapHandler(h))
}

func wrapHandler(h http.Handler) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		context.Set(r, "params", ps)
		h.ServeHTTP(w, r)
	}
}
