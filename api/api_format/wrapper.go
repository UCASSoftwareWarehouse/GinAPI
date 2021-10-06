package api_format

import (
	"github.com/gin-gonic/gin"
)

func Wrap(handler interface{}) func(c *gin.Context) {
	var target func(c *gin.Context)
	switch handler.(type) {
	case SimpleJSONHandler:
		target = UnwrapSimpleJSONHandler(handler.(SimpleJSONHandler))
		break
	case NormalHandler:
		target = handler.(func(ctx *gin.Context))
		break
	default:
		panic("Unsupported Search API type")
	}
	return target
}
