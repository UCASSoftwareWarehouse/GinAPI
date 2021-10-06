package api_format

import (
	"GinAPI/err"
	"github.com/gin-gonic/gin"
)

type SimpleJSONHandler func(*gin.Context) (*JSONRespFormat, *err.APIErr)
type NormalHandler func(*gin.Context)
