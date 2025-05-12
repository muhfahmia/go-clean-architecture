package enum

type ErrorType string

const (
	// Client Errors (4xx)
	ErrorBadRequest          ErrorType = "bad_request"          // 400 - Malformed request
	ErrorUnauthorized        ErrorType = "unauthorized"         // 401 - Authentication needed
	ErrorForbidden           ErrorType = "forbidden"            // 403 - No permission
	ErrorNotFound            ErrorType = "not_found"            // 404 - Resource not found
	ErrorNotAllowed          ErrorType = "method_not_allowed"   // 405 - Wrong HTTP method
	ErrorNotAcceptable       ErrorType = "not_acceptable"       // 406 - Can't satisfy Accept header
	ErrorConflict            ErrorType = "conflict"             // 409 - Resource state conflict
	ErrorGone                ErrorType = "gone"                 // 410 - Resource permanently unavailable
	ErrorValidation          ErrorType = "validation"           // 400 - Input validation failed
	ErrorTooMany             ErrorType = "too_many_requests"    // 429 - Rate limited
	ErrorUnprocessableEntity ErrorType = "unprocessable_entity" // 422 - Request was well-formed but unable to be followed due to semantic errors

	// Server Errors (5xx)
	ErrorInternal       ErrorType = "internal"        // 500 - Generic server error
	ErrorNotImplemented ErrorType = "not_implemented" // 501 - Unsupported functionality
	ErrorUnavailable    ErrorType = "unavailable"     // 503 - Maintenance/downtime
	ErrorTimeout        ErrorType = "timeout"         // 504 - Upstream timeout
)

func (e ErrorType) String() string {
	switch e {
	// Client Errors
	case ErrorBadRequest:
		return "bad request"
	case ErrorUnauthorized:
		return "unauthorized"
	case ErrorForbidden:
		return "forbidden"
	case ErrorNotFound:
		return "not found"
	case ErrorNotAllowed:
		return "method not allowed"
	case ErrorNotAcceptable:
		return "not acceptable"
	case ErrorConflict:
		return "conflict"
	case ErrorGone:
		return "gone"
	case ErrorValidation:
		return "validation failed"
	case ErrorTooMany:
		return "too many requests"

	// Server Errors
	case ErrorInternal:
		return "internal server error"
	case ErrorNotImplemented:
		return "not implemented"
	case ErrorUnavailable:
		return "service unavailable"
	case ErrorTimeout:
		return "gateway timeout"
	default:
		return "unknown error"
	}
}

func (e ErrorType) HttpCode() int {
	switch e {
	// Client Errors
	case ErrorBadRequest, ErrorValidation:
		return 400
	case ErrorUnauthorized:
		return 401
	case ErrorForbidden:
		return 403
	case ErrorNotFound:
		return 404
	case ErrorNotAllowed:
		return 405
	case ErrorNotAcceptable:
		return 406
	case ErrorConflict:
		return 409
	case ErrorGone:
		return 410
	case ErrorUnprocessableEntity:
		return 422
	case ErrorTooMany:
		return 429
	// Server Errors
	case ErrorInternal:
		return 500
	case ErrorNotImplemented:
		return 501
	case ErrorUnavailable:
		return 503
	case ErrorTimeout:
		return 504
	default:
		return 500 // Default to Internal Server Error
	}
}

// Additional helpful methods
func (e ErrorType) IsClientError() bool {
	code := e.HttpCode()
	return code >= 400 && code < 500
}

func (e ErrorType) IsServerError() bool {
	code := e.HttpCode()
	return code >= 500 && code < 600
}
