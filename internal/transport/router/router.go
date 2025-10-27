package router

import (
	"net/http"
	"strings"
)

type Middleware func(http.Handler) http.Handler

type Router struct {
	mux        *http.ServeMux
	prefix     string
	middleware []Middleware
}

func New() *Router {
	return &Router{
		mux: http.NewServeMux(),
	}
}

func (r *Router) Group(prefix string) *Router {
	return &Router{
		mux:        r.mux,
		prefix:     r.prefix + prefix,
		middleware: append([]Middleware{}, r.middleware...),
	}
}

func (r *Router) Use(mws ...Middleware) {
	r.middleware = append(r.middleware, mws...)
}

func (r *Router) handle(method, pattern string, handler http.Handler) {

	for i := len(r.middleware) - 1; i >= 0; i-- {
		handler = r.middleware[i](handler)
	}

	fullPattern := strings.TrimRight(r.prefix+pattern, "/")
	r.mux.Handle(method+" "+fullPattern, handler)
}

func (r *Router) Get(pattern string, handler http.HandlerFunc) {
	r.handle("GET", pattern, handler)
}

func (r *Router) Post(pattern string, handler http.HandlerFunc) {
	r.handle("POST", pattern, handler)
}

func (r *Router) Put(pattern string, handler http.HandlerFunc) {
	r.handle("PUT", pattern, handler)
}

func (r *Router) Delete(pattern string, handler http.HandlerFunc) {
	r.handle("DELETE", pattern, handler)
}

func (r *Router) Handler() http.Handler {
	return r.mux
}
