package model

import (
	"github.com/gofiber/fiber/v2"
	"github.com/muhfahmia/pkg/enum"
)

type httpResponse struct {
	httpCode int       `json:"-"`
	appError AppError  `json:"-"`
	Success  bool      `json:"success"`
	Message  string    `json:"message"`
	Err      *string   `json:"error,omitempty"`
	Data     fiber.Map `json:"data"`
}

type httpResponseBuilder struct {
	ctx      *fiber.Ctx
	response *httpResponse
}

func NewHttpResponseBuilder(c *fiber.Ctx) *httpResponseBuilder {
	b := &httpResponseBuilder{}
	b.ctx = c
	b.response = b.Build()
	return b
}

// WithSuccess sets the success status
func (b *httpResponseBuilder) WithRequestParameter(request any) error {
	if errB := b.ctx.BodyParser(&request); errB != nil {
		appError := NewAppError(errB, enum.ErrorUnprocessableEntity, "Request parameters are invalid or missing")
		b.WithError(appError)
		b.WithHttpCode(fiber.StatusUnprocessableEntity)
		return errB
	}
	return nil
}

// WithSuccess sets the success status
func (b *httpResponseBuilder) WithSuccess(success bool) *httpResponseBuilder {
	b.response.Success = success
	return b
}

// WithHttpCode sets the HTTP status code
func (b *httpResponseBuilder) WithHttpCode(code int) *httpResponseBuilder {
	b.response.httpCode = code
	return b
}

// WithMessage sets the response message
func (b *httpResponseBuilder) WithMessage(message string) *httpResponseBuilder {
	b.response.Message = message
	return b
}

// WithError sets the error information
func (b *httpResponseBuilder) WithError(err AppError) *httpResponseBuilder {
	b.response.Message = err.GetErrorMessage()
	if err != nil {
		if err.GetErrorType() == enum.ErrorValidation {
			b.WithData("validationErrors", err.GetErrorDetail())
		} else {
			errStr := err.GetError().Error()
			b.response.Err = &errStr
		}
		b.response.appError = err
	}
	return b
}

// WithData adds a key-value pair to the response data
func (b *httpResponseBuilder) WithData(key string, value any) *httpResponseBuilder {
	if b.response.Data == nil {
		b.response.Data = make(fiber.Map)
	}
	b.response.Data[key] = value
	return b
}

// WithDataMap merges a map into the response data
func (b *httpResponseBuilder) WithDataMap(data map[string]any) *httpResponseBuilder {
	if b.response.Data == nil {
		b.response.Data = make(fiber.Map)
	}
	for k, v := range data {
		b.response.Data[k] = v
	}
	return b
}

func (b *httpResponseBuilder) Build() *httpResponse {
	return &httpResponse{}
}

// Send sends the response using Fiber context
func (b *httpResponseBuilder) Send() error {
	if b.response.httpCode == 0 {
		// Set default status code based on success
		if b.response.Success {
			b.response.httpCode = fiber.StatusOK
		} else {
			b.response.httpCode = fiber.StatusBadRequest
		}
	}
	return b.ctx.Status(b.response.httpCode).JSON(b.response)
}
