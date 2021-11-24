package middleware

import (
	"fmt"
	"net/http"
)

type LogMiddleware struct {
	HttpHandler http.Handler
}

func (middleware *LogMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fmt.Println(request.Method, request.URL)
	middleware.HttpHandler.ServeHTTP(writer, request)
}
