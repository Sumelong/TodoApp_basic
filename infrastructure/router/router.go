package router

import (
	"TodoApp_basic/adapters/logger"
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"regexp"
	"syscall"
	"time"
)

type Route struct {
	Method  string
	Pattern string
	Handler http.Handler
}

type Router struct {
	log    logger.Logger
	port   Port
	routes []Route
}

func NewRouter(logger logger.Logger, port Port) *Router {
	return &Router{
		log:  logger,
		port: port,
	}
}

func (r *Router) AddRoute(method, path string, handler http.Handler) {
	r.routes = append(r.routes, Route{Method: method, Pattern: path, Handler: handler})
}

func (r *Router) GET(path string, handler http.Handler) {
	r.AddRoute("GET", path, handler)
}

func (r *Router) POST(path string, handler http.Handler) {
	r.AddRoute("POST", path, handler)
}

func (r *Router) PUT(path string, handler http.Handler) {
	r.AddRoute("PUT", path, handler)
}

func (r *Router) DELETE(path string, handler http.Handler) {
	r.AddRoute("DELETE", path, handler)
}

func (r *Router) getHandler(method, path string) http.Handler {
	for _, route := range r.routes {
		re := regexp.MustCompile(route.Pattern)
		if route.Method == method && re.MatchString(path) {
			return route.Handler
		}
	}
	return http.NotFoundHandler()
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	method := req.Method

	handler := r.getHandler(method, path)

	// handler middlewares go here

	handler.ServeHTTP(w, req)
}

func (r *Router) Listen() {

	var req *http.Request
	path := req.URL.Path
	method := req.Method

	server := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 15 * time.Second,
		Addr:         fmt.Sprintf(":%d", r.port),
		Handler:      r.getHandler(method, path),
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		r.log.WithFields(logger.Fields{"port": r.port}).Info("Starting HTTP Server")
		if err := server.ListenAndServe(); err != nil {
			r.log.WithError(err).Error("Error starting HTTP server")
		}
	}()

	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
	}()

	if err := server.Shutdown(ctx); err != nil {
		r.log.WithError(err).Error("Server Shutdown Failed")
	}

	r.log.Info("Service down")
}
