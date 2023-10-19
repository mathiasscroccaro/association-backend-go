package api

import (
	"bytes"
	"captcha_example/internal/domain"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sync"
	"testing"
	"time"
)

func TestGenerateCaptchaAndValidate(t *testing.T) {
	os.Setenv("SECRET_KEY", "secret")

	var wg sync.WaitGroup
	wg.Add(1)

	server := CreateServer()

	go func() {
		defer wg.Done()
		fmt.Printf("Server is running on %s...\n", server.Addr)
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			fmt.Printf("Error: %s\n", err)
		}
	}()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := http.Get("http://localhost:8080/captcha")
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, resp.StatusCode)
	}

	resBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("server: could not read request body: %s\n", err)
	}
	var captchaResponse domain.CaptchaResponse
	err = json.Unmarshal(resBody, &captchaResponse)

	if err != nil {
		fmt.Printf("server: could not unmarshal request body: %s\n", err)
	}

	if captchaResponse.Base64CaptchaImage == "" {
		t.Errorf("Expected base64CaptchaImage, but got %s", captchaResponse.Base64CaptchaImage)
	}
	if captchaResponse.HashedSolution == "" {
		t.Errorf("Expected hashedSolution, but got %s", captchaResponse.HashedSolution)
	}

	resp.Body.Close()
	server.Shutdown(ctx)
}

func TestInvalidCaptchaSolution(t *testing.T) {
	os.Setenv("SECRET_KEY", "secret")
	var wg sync.WaitGroup
	wg.Add(1)

	server := CreateServer()

	go func() {
		defer wg.Done()
		fmt.Printf("Server is running on %s...\n", server.Addr)
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			fmt.Printf("Error: %s\n", err)
		}
	}()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := http.Post("http://localhost:8080/captcha", "application/json", bytes.NewBuffer([]byte(`{}`)))
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, resp.StatusCode)
	}

	resBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("server: could not read request body: %s\n", err)
	}
	var validationResponse domain.ValidationResponse
	err = json.Unmarshal(resBody, &validationResponse)

	if err != nil {
		fmt.Printf("server: could not unmarshal request body: %s\n", err)
	}

	if validationResponse.IsValid == true {
		t.Errorf("Expected false, but got true")
	}

	resp.Body.Close()
	server.Shutdown(ctx)
}
