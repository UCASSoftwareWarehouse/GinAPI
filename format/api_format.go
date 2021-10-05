package format

import (
	"GinAPI/err"
	"github.com/gin-gonic/gin"
	"net/http"
)

type SimpleJSONHandler func(*gin.Context) (*JSONRespFormat, *err.APIErr)
type NormalHandler func(*gin.Context)

func UnwrapSimpleJSONHandler(handler SimpleJSONHandler) func(c *gin.Context) {
	return func(c *gin.Context) {
		resp, e := handler(c)
		if e != nil {
			c.JSON(e.Status, NewJSONResp(e.Status, e.Message, nil))
			return
		}
		if resp == nil {
			c.JSON(http.StatusOK, SimpleOKResp(nil))
			return
		}
		c.JSON(resp.Status, NewJSONResp(resp.Status, resp.Message, resp.Data))
	}
}