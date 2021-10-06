package search_api

import (
	"GinAPI/api/api_format"
	"GinAPI/err"
	"GinAPI/internal/search"
	"github.com/gin-gonic/gin"
)

type API struct{}

const prefix = "/search"

func Register(r *gin.Engine) {
	api := API{}
	searchRouter := r.Group(prefix)
	searchRouter.GET("/ping", api_format.Wrap(api.pingCodeSim()))
	searchRouter.POST("/query", api_format.Wrap(api.query()))
}

func (API) pingCodeSim() api_format.SimpleJSONHandler {
	return func(c *gin.Context) (*api_format.JSONRespFormat, *err.APIErr) {
		return search.GetHandler().Ping(c)
	}
}

func (API) query() api_format.SimpleJSONHandler {
	return func(c *gin.Context) (*api_format.JSONRespFormat, *err.APIErr) {
		return search.GetHandler().Query(c)
	}
}
