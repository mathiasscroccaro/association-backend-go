package domain

type IApiError interface {
	Error() string
	StatusCode() int
	Detail() string
}

func NewApiError(message string, code int, detail string) IApiError {
	return ApiError{
		Message:       message,
		Code:          code,
		DetailMessage: detail,
	}
}

type ApiError struct {
	Message       string
	Code          int
	DetailMessage string
}

func (ce ApiError) Error() string {
	return ce.Message
}

func (ce ApiError) StatusCode() int {
	return ce.Code
}

func (ce ApiError) Detail() string {
	return ce.DetailMessage
}

type IRepositoryError interface {
	Error() string
}

type RepositoryError struct {
	Message string
}

func (ce RepositoryError) Error() string {
	return ce.Message
}

var (
	ErrPreRegisterNotFound           IRepositoryError = RepositoryError{"register not found"}
	ErrPreRegisterDocumentDuplicated IRepositoryError = RepositoryError{"document already registered"}
	ErrOtherRepositoryError          IRepositoryError = RepositoryError{"other repository error"}

	Err400 IApiError = ApiError{"invalid request", 400, ""}
	Err404 IApiError = ApiError{"record not found", 404, ""}
	Err500 IApiError = ApiError{"internal server error", 500, ""}
)
