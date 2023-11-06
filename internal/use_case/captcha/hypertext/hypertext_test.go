package hypertext

import (
	"os"
	"testing"
)

func TestGenerateCaptchaHyperTextBySecretKey(t *testing.T) {
	response := GenerateCaptchaHyperTextBySecretKey()
	if string(response) == "" {
		t.Errorf("Failed to generate captcha hyper text by secret key")
	}
	os.WriteFile("test.html", response, 0644)
}
