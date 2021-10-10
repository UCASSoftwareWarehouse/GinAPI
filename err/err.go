package err

import (
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
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
	Status  int    `json:"status"`
	Stable  bool   `json:"stable"`
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

func WrapRPCError(err error) *APIErr {
	// TODO 使用status透传下游错误，之后需要统一处理rpc错误。
	errStatus, ok := status.FromError(err)
	if !ok {
		return InternalErr.CustomMessageF(err.Error())
	}
	log.Printf("wrap rpc err, message=[%s], code=[%d]", errStatus.Message(), errStatus.Code())
	switch errStatus.Code() {
	case codes.InvalidArgument, codes.NotFound, codes.PermissionDenied:
		return BadRequestErr.CustomMessage(errStatus.Message())
	case codes.Unimplemented:
		return ForbiddenErr.CustomMessage(errStatus.Message())
	default:
		return InternalErr.CustomMessage(errStatus.Message())
	}
}
