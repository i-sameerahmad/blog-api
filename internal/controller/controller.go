package controller

import "github.com/ImranZahoor/blog-api/internal/service"

type (
	Controller struct {
		service service.Service
	}
)

func NewController(svc service.Service) Controller {
	return Controller{service: svc}
}

const (
	ErrorInvalidPayload   = "Invalid Payload"
	ErrorInvalidSearchKey = "Invalid Search Key"
	MessageSuccess        = "Success"
	MessageFailure        = "Failed"
	ErrorNotFound         = "Not Found"
)
