package enum

type ErrorType string

const (
	ErrorValidation    ErrorType = "validation"
	ErrorNotFound      ErrorType = "not_found"
	ErrorConflict      ErrorType = "conflict"
	ErrorUnauthorized  ErrorType = "unauthorized"
	ErrorInternal      ErrorType = "internal"
	ErrorBadRequest    ErrorType = "bad_request"
	ErrorForbidden     ErrorType = "forbidden"
	ErrorNotAcceptable ErrorType = "not_acceptable"
)

func (e ErrorType) String() string {
	switch e {
	case ErrorValidation:
		return "validation"
	case ErrorNotFound:
		return "not found"
	case ErrorConflict:
		return "conflict"
	case ErrorUnauthorized:
		return "unauthorized"
	case ErrorInternal:
		return "internal"
	case ErrorBadRequest:
		return "bad request"
	case ErrorForbidden:
		return "forbidden"
	case ErrorNotAcceptable:
		return "not_acceptable"
	default:
		return "unknown error type"
	}
}
