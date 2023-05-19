package router

import (
	"fmt"
	"net/http"
	"path"
)

type Router struct {
	handlers map[string]http.HandlerFunc
}

func New() *Router{
	return &Router{
		make(map[string]http.HandlerFunc),
	}
}

// register path and its handler function
func (r *Router) Add(path string, handler http.HandlerFunc) {
	r.handlers[path] = handler
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// combine request method and path to check
	check := req.Method + " " + req.URL.Path

	// loop through registered path
	for pattern, handlerFunc := range r.handlers {
		// match given path with registered one
		ok, err := path.Match(pattern, check)
		if ok && err == nil {
			// call function to handle mathed path
			handlerFunc(w, req)
			return
		} else if err != nil {
			fmt.Fprint(w, err)
		}
	}
	http.NotFound(w, req)
}