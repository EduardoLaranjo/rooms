package web

import (
	"dottime.dev/room/cmd/rooms/handlers"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
)

type routerMux struct {
	log *log.Logger
	mux *http.ServeMux
}

func NewServer(log *log.Logger, db *sqlx.DB) *routerMux {

	s := &routerMux{
		log: log,
		mux: http.NewServeMux(),
	}

	c := handlers.NewChat(log, db)

	s.registerHandler("/", c.Homepage(), logMiddleware(log))

	return s
}

func (s *routerMux) registerHandler(pattern string, handler http.HandlerFunc, middlewares ...Middleware) {

	for _, middleware := range middlewares {
		handler = middleware(handler)
	}

	s.mux.HandleFunc(pattern, handler)
}

func (s *routerMux) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	s.mux.ServeHTTP(writer, request)
}
