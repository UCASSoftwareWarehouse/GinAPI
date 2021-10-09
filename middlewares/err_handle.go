package middlewares

import (
	"GinAPI/api/api_format"
	AErr "GinAPI/err"
	"github.com/gin-gonic/gin"
)

func ErrHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		if length := len(c.Errors); length > 0 {
			e := c.Errors[length-1]
			err := e.Err
			if err != nil {
				var aErr *AErr.APIErr
				if e, ok := err.(*AErr.APIErr); ok {
					aErr = e
				} else if e, ok := err.(error); ok {
					aErr = AErr.ForbiddenErr.CustomMessage(e.Error())
				} else {
					aErr = AErr.InternalErr
				}
				// 记录一个错误的日志
				api_format.UnwrapErr(c, aErr)
				return
			}
		}

	}
}

func NotFoundHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		api_format.UnwrapErr(c, AErr.NotFoundErr)
	}
}
