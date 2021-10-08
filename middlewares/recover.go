package middlewares

import (
	"GinAPI/api/api_format"
	AErr "GinAPI/err"
	"github.com/gin-gonic/gin"
	"log"
	"runtime/debug"
)

func Recover() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("panic: %v\n", r)
				debug.PrintStack()
				api_format.UnwrapErr(c, AErr.InternalErr)
				c.Abort()
			}
		}()
		c.Next()
	}
}

