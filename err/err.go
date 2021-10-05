package err

import "net/http"

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
)

type APIErr struct {
	Message string
	Status  int
	Stable  bool
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