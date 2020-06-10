package web

import (
	"dottime.dev/room/internal/business"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
)

type RouterMux struct {
	log *log.Logger
	mux *http.ServeMux
	// router *someRouter
	// email  EmailSender
}

func NewRouter(log *log.Logger, db *sqlx.DB) *RouterMux {

	s := &RouterMux{
		log: log,
		mux: http.NewServeMux(),
	}

	c := business.NewChat(log, db)

	s.registerHandler("/", c.Homepage())

	return s
}

func (s *RouterMux) registerHandler(pattern string, handlerFunc http.HandlerFunc, middleware ...http.HandlerFunc) {
	s.mux.HandleFunc(pattern, handlerFunc)
}

func (s *RouterMux) ServeHTTP(writer http.ResponseWriter, reader *http.Request) {
	log.Println(reader)
	s.mux.ServeHTTP(writer, reader)
}
