package handlers

import (
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
)

type Chat struct {
	*log.Logger
	*sqlx.DB
}

const ContextId string = "id"

func NewChat(log *log.Logger, db *sqlx.DB) *Chat {
	return &Chat{
		Logger: log,
		DB:     db,
	}
}

func (c *Chat) Homepage() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		id := request.Context().Value(ContextId).(string)
		writer.Write([]byte(id))
	}
}
