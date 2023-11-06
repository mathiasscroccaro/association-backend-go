package view

import (
	"captcha_example/internal/domain"
	"captcha_example/internal/use_case/pre_register"
	"captcha_example/internal/use_case/repository"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func SavePreRegisterView(w http.ResponseWriter, r *http.Request) {

	fmt.Println(r)
	var preRegistrationRequest domain.PreRegistrationRequest
	err := json.NewDecoder(r.Body).Decode(&preRegistrationRequest)

	if err != nil {
		http.Error(w, "Failed to decode JSON request", http.StatusBadRequest)
		return
	}

	// isValid := pre_register.ValidateCaptchaSolution(preRegistrationRequest.HashedSolution, preRegistrationRequest.Solution)

	// if !isValid {
	// 	http.Error(w, "Invalid captcha solution", http.StatusBadRequest)
	// 	return
	// }

	repositoryInstance := repository.GetDB()

	apiErr := pre_register.SavePreRegister(repositoryInstance, preRegistrationRequest.FormData)

	w.Header().Set("Content-Type", "application/json")

	switch apiErr.StatusCode() {
	case http.StatusBadRequest:
		log.Print(apiErr.Error())
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(domain.PreRegistrationResponse{
			Detail: apiErr.Detail(),
		})
		return
	case http.StatusInternalServerError:
		log.Print(apiErr.Error())
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(domain.PreRegistrationResponse{
			Detail: apiErr.Detail(),
		})
		return
	}

	jsonData, err := json.Marshal(domain.PreRegistrationResponse{
		Detail: "Pre-registration saved successfully",
	})

	w.WriteHeader(http.StatusCreated)
	w.Write(jsonData)
}
