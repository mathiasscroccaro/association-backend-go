package api

import (
	"captcha_example/internal/route"
	"net/http"
)

func CreateServer() *http.Server {
	var server *http.Server

	server = &http.Server{
		Addr:    ":8080",
		Handler: route.Router(),
	}

	return server
}
