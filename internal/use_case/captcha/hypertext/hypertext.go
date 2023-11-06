package hypertext

import (
	"bytes"
	"captcha_example/internal/use_case/captcha"
	"html/template"
)

func GenerateCaptchaHyperTextBySecretKey() []byte {
	buffer := bytes.Buffer{}

	response := captcha.GenerateCaptchaBySecretKey()

	component := `<img src="data:image/jpeg;base64,{{.Base64CaptchaImage}}" />`

	tmpl, _ := template.New("captcha").Parse(component)

	tmpl.Execute(&buffer, response)

	return buffer.Bytes()
}
