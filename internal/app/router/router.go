package router

import (
	"net/http"

	"github.com/gorilla/mux"

	"golang.org/x/time/rate"
)

type Router struct {
	router     *mux.Router
	controller controller
}

type Route struct {
	Name    string
	Method  string
	Path    string
	Handler http.Handler
}

type controller interface {
	GetBooks(w http.ResponseWriter, r *http.Request)
	DeleteBook(w http.ResponseWriter, r *http.Request)
}

func (r *Router) GetRouter() http.Handler {
	return r.router
}

func limit(next http.Handler) http.Handler {
	limiter := rate.NewLimiter(10, 1)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if limiter.Allow() == false {
			http.Error(w, http.StatusText(429), http.StatusTooManyRequests)

			return
		}

		next.ServeHTTP(w, r)
	})
}

func New() *Router {
	router := mux.NewRouter()

	r := Router{
		router: router,
	}

	return &r
}

func (r *Router) AppendRoutes(ctr controller) {
	routes := []Route{
		{
			Name:    "books",
			Path:    "/books",
			Method:  http.MethodGet,
			Handler: limit(http.HandlerFunc(ctr.GetBooks)),
		},
		{
			Name:    "books",
			Path:    "/books",
			Method:  http.MethodDelete,
			Handler: limit(http.HandlerFunc(ctr.DeleteBook)),
		},
	}

	for _, route := range routes {
		r.router.
			Methods(route.Method).
			Name(route.Name).
			Path(route.Path).
			Handler(route.Handler)
	}
}
