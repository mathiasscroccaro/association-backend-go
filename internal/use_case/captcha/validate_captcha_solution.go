package captcha

import (
	"captcha_example/internal/domain"
	"captcha_example/internal/use_case/captcha/hash"
	"captcha_example/internal/use_case/env"
)

func ValidateCaptchaSolution(arguments domain.CaptchaRequest) domain.ValidationResponse {
	secretKey := env.GetEnvVariableOrDefault("SECRET_KEY", "secret")

	return domain.ValidationResponse{
		IsValid: hash.IsHashedKeyEqualToSolution(arguments.HashedSolution, arguments.Solution, secretKey),
	}
}
