package routes

import (
	"encoding/json"
	"net/http"
)

type Route struct {
	Method  string
	Pattern string
	Handler http.Handler
}

type Router struct {
	routes []Route
}

func NewRouter() *Router {
	return &Router{}
}

func (r Router) AddRoute(method, path string, handler http.Handler) {
	r.routes = append(r.routes, Route{
		Method:  method,
		Pattern: path,
		Handler: handler,
	})
}

func (r Router) GET(path string, handler http.Handler) {
	r.AddRoute("GET", path, handler)
}

func (r Router) POST(path string, handler http.Handler) {
	r.AddRoute("POST", path, handler)
}

func (r Router) PUT(path string, handler http.Handler) {
	r.AddRoute("PUT", path, handler)
}

func (r Router) PATCH(path string, handler http.Handler) {
	r.AddRoute("PATCH", path, handler)
}

func (r Router) DELETE(path string, handler Handler) {
	r.AddRoute("DELETE", path, handler)
}

type Handler func(r *http.Request) (statusCode int, data map[string]interface{})

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	statusCode, data := h(r)
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}
