package search_api

import (
	"GinAPI/err"
	format2 "GinAPI/format"
	"GinAPI/internal/search"
	"github.com/gin-gonic/gin"
)

type SearchAPI struct {}

const prefix = "/search"

func Register(r *gin.Engine) {
	api := SearchAPI{}
	searchRouter := r.Group(prefix)
	searchRouter.GET("/ping", api.wrap(api.pingCodeSim()))
}

func (SearchAPI) wrap(handler interface{}) func(c *gin.Context) {
	switch handler.(type) {
	case format2.SimpleJSONHandler:
		return format2.UnwrapSimpleJSONHandler(handler.(format2.SimpleJSONHandler))
	case format2.NormalHandler:
		return handler.(func(ctx *gin.Context))
	default:
		panic("Unsupported Search API type")
	}
}

func (SearchAPI) pingCodeSim() format2.SimpleJSONHandler {
	return func(c *gin.Context) (*format2.JSONRespFormat, *err.APIErr) {
		return search.GetHandler().Ping(c)
	}
}