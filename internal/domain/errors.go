package domain

type IApiError interface {
	Error() string
}

type ApiError struct {
	Message string
	Code    int
}

func (ce ApiError) Error() string {
	return ce.Message
}

func (ce ApiError) StatusCode() int {
	return ce.Code
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

	Err400 IApiError = ApiError{"invalid request", 400}
	Err404 IApiError = ApiError{"record not found", 404}
	Err500 IApiError = ApiError{"internal server error", 500}
)
