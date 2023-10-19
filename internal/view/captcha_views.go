package view

import (
	"captcha_example/internal/domain"
	"captcha_example/internal/use_case/captcha"
	"encoding/json"
	"net/http"
)

func GetCaptchaView(w http.ResponseWriter, r *http.Request) {
	captchaResponse := captcha.GenerateCaptchaBySecretKey()

	jsonData, err := json.Marshal(captchaResponse)

	if err != nil {
		http.Error(w, "Failed to create JSON response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func ValidateCaptchaSolution(w http.ResponseWriter, r *http.Request) {
	var captchaRequest domain.CaptchaRequest
	err := json.NewDecoder(r.Body).Decode(&captchaRequest)

	if err != nil {
		http.Error(w, "Failed to decode JSON request", http.StatusBadRequest)
		return
	}

	validationResponse := captcha.ValidateCaptchaSolution(captchaRequest)

	jsonData, err := json.Marshal(validationResponse)

	if err != nil {
		http.Error(w, "Failed to create JSON response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}
