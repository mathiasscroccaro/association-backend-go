package captcha

import (
	"bytes"
	"captcha_example/internal/domain"
	"captcha_example/internal/use_case/captcha/hash"
	"captcha_example/internal/use_case/env"
	"encoding/base64"

	"github.com/steambap/captcha"
)

func GenerateBase64CaptchaImageAndSolution() (base64Image, solution string) {
	var binaryBuffer bytes.Buffer

	data, _ := captcha.New(150, 50)
	data.WriteImage(&binaryBuffer)

	base64Image = base64.StdEncoding.EncodeToString(binaryBuffer.Bytes())
	solution = data.Text

	return base64Image, solution
}

func GenerateCaptchaBySecretKey() domain.CaptchaResponse {
	secretKey := env.GetEnvVariableOrDefault("SECRET_KEY", "secret")

	base64Image, solution := GenerateBase64CaptchaImageAndSolution()

	hashedSolution, _ := hash.HashSolutionBySecretKey(solution, secretKey)

	return domain.CaptchaResponse{
		Base64CaptchaImage: base64Image,
		HashedSolution:     hashedSolution,
	}
}
