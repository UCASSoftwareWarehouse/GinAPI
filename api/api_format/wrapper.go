package api_format

import (
	"GinAPI/err"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Wrap(handler interface{}) func(c *gin.Context) {
	var target func(c *gin.Context)
	switch handler.(type) {
	case SimpleJSONHandler:
		target = wrapSimpleJSONHandler(handler.(SimpleJSONHandler))
		break
	case NormalHandler:
		target = wrapNormalHandler(handler.(NormalHandler))
		break
	default:
		panic("Unsupported Search API type")
	}
	return target
}

func wrapSimpleJSONHandler(handler SimpleJSONHandler) func(c *gin.Context) {
	return func(c *gin.Context) {
		c.Set("some_key", "some_value")
		resp, e := handler(c)
		if e != nil {
			UnwrapErr(c, e)
			return
		}
		if resp == nil {
			c.JSON(http.StatusOK, SimpleOKResp(nil))
			return
		}
		c.JSON(resp.Status, NewJSONResp(resp.Status, resp.Message, resp.Data))
	}
}

func wrapNormalHandler(handler NormalHandler) func(c *gin.Context) {
	return func(c *gin.Context) {
		handler(c)
	}
}

func UnwrapErr(c *gin.Context, e *err.APIErr) {
	c.JSON(e.Status, NewJSONResp(e.Status, e.Message, nil))
}
