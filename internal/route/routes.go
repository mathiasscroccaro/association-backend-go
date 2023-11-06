package route

import (
	"captcha_example/internal/view"

	"github.com/go-chi/chi/v5"
)

func Router() *chi.Mux {
	router := chi.NewRouter()

	router.Get("/captcha", view.GetCaptchaView)
	router.Post("/captcha", view.ValidateCaptchaSolution)

	router.Post("/pre-register", view.SavePreRegisterView)

	return router
}
