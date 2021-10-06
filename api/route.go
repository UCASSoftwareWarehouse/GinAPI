package api

import (
	"GinAPI/api/access_api"
	"GinAPI/api/search_api"
	"GinAPI/api/swagger_api"
	"github.com/gin-gonic/gin"
)

func Register(r *gin.Engine) {
	r.GET("/ping", ping)
	swagger_api.Register(r)
	search_api.Register(r)
	access_api.Register(r)
}

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
