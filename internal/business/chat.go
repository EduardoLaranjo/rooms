package business

import (
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
)

type Chat struct {
	*log.Logger
	*sqlx.DB
}

func NewChat(log *log.Logger, db *sqlx.DB) *Chat {
	return &Chat{
		Logger: log,
		DB:     db,
	}
}

func (c *Chat) Homepage() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

	}
}
