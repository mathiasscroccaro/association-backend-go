package domain

type IError interface {
	Error() string
}

type ApiError struct {
	Message string
	Code    int
}

type UseCaseError struct {
	Message string
}

func (ce ApiError) Error() string {
	return ce.Message
}

func (ce ApiError) StatusCode() int {
	return ce.Code
}

func (ce UseCaseError) Error() string {
	return ce.Message
}

var (
	ErrPreRegisterNotFound           IError = UseCaseError{"register not found"}
	ErrPreRegisterDocumentDuplicated IError = UseCaseError{"document already registered"}

	Err400 IError = ApiError{"invalid request", 400}
	Err404 IError = ApiError{"record not found", 404}
	Err500 IError = ApiError{"internal server error", 500}
)
