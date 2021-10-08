package err

import (
	"fmt"
	"net/http"
)

var (
	InternalErr = &APIErr{
		Message: "Server Internal Error",
		Status:  http.StatusInternalServerError,
		Stable:  false,
	}
	NotFoundErr = &APIErr{
		Message: "Not Found",
		Status:  http.StatusNotFound,
		Stable:  true,
	}
	BadRequestErr = &APIErr{
		Message: "Bad Request",
		Status:  http.StatusBadRequest,
		Stable:  true,
	}
	ForbiddenErr = &APIErr{
		Message: "Forbidden",
		Status:  http.StatusForbidden,
		Stable:  true,
	}
)

type APIErr struct {
	Message string `json:"message"`
	Status  int `json:"status"`
	Stable  bool `json:"stable"`
}

func (A *APIErr) Error() string {
	return A.Message
}

func (A *APIErr) CustomMessage(message string) *APIErr {
	return &APIErr{
		Message: message,
		Status:  A.Status,
		Stable:  A.Stable,
	}
}

func (A *APIErr) CustomMessageF(msg string, formatter ...interface{}) *APIErr {
	return &APIErr{
		Message: fmt.Sprintf(msg, formatter...),
		Status:  A.Status,
		Stable:  A.Stable,
	}
}
