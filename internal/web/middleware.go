package web

import (
	"context"
	"dottime.dev/room/cmd/rooms/handlers"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

type Middleware func(handlerFunc http.HandlerFunc) http.HandlerFunc

func logMiddleware(log *log.Logger) Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(writer http.ResponseWriter, request *http.Request) {
			//gen id
			requestId := strconv.FormatUint(rand.Uint64(), 10)
			//get context
			ctx := request.Context()
			//set
			request = request.WithContext(context.WithValue(ctx, handlers.ContextId, requestId))

			log.Printf("id: %s method: %s path: %s ip:%s",
				requestId, request.Method, request.URL.Path, request.RemoteAddr,
			)

			next(writer, request)
		}
	}

}
