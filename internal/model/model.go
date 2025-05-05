package model

import "github.com/gofiber/fiber/v2"

type response struct {
	HttpCode int       `json:"-"`
	Success  bool      `json:"success"`
	Message  string    `json:"message"`
	Err      *string   `json:"error,omitempty"`
	Data     fiber.Map `json:"data"`
}

func NewAppResponse() *response {
	return &response{}
}

func (r *response) SetError(err error) {
	if err != nil {
		errStr := err.Error() // Get the string from error
		r.Err = &errStr       // Assign its address to the *string field
	} else {
		r.Err = nil // Explicitly set nil if no error
	}
}

func (r *response) SetSuccess(success bool) {
	r.Success = success
}

func (r *response) SetHttpCode(httpCode int) {
	r.HttpCode = httpCode
}

func (r *response) SetMessage(message string) {
	r.Message = message
}

func (r *response) SetData(key string, value any) {
	if r.Data == nil {
		r.Data = fiber.Map{}
	}
	r.Data[key] = value
}
