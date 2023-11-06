package api

import (
	"captcha_example/internal/route"
	"captcha_example/internal/use_case/repository"
	"net/http"
)

func CreateServer() *http.Server {
	var server *http.Server

	repository.InitDBConnectionByEnvConfig()

	server = &http.Server{
		Addr:    ":8080",
		Handler: route.Router(),
	}

	return server
}
