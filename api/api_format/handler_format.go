package api_format

import (
	"GinAPI/err"
	"github.com/gin-gonic/gin"
)

type JSONHandler func(*gin.Context) (*JSONRespFormat, *err.APIErr)
type NormalHandler func(*gin.Context)
