package pre_register

import (
	"captcha_example/internal/use_case/captcha/hash"
	"captcha_example/internal/use_case/env"
)

func ValidateCaptchaSolution(hashedSolution, solution string) bool {
	secretKey := env.GetEnvVariableOrDefault("SECRET_KEY", "secret")
	return hash.IsHashedKeyEqualToSolution(hashedSolution, solution, secretKey)
}
