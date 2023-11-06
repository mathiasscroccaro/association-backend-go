package domain

type CaptchaResponse struct {
	Base64CaptchaImage string `json:"base64CaptchaImage"`
	HashedSolution     string `json:"hashedSolution"`
}

type CaptchaRequest struct {
	HashedSolution string `json:"hashedSolution"`
	Solution       string `json:"solution"`
}

type ValidationResponse struct {
	IsValid bool `json:"isValid"`
}

type CaptchaData struct {
	Base64CaptchaImage string `json:"base64CaptchaImage"`
	HashedSolution     string `json:"hashedSolution"`
}

type CaptchaSolutionData struct {
	HashedSolution string `json:"hashedSolution"`
	Solution       string `json:"solution"`
}
