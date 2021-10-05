package api

import (
	"GinAPI/api/search_api"
	"GinAPI/api/swagger"
	"github.com/gin-gonic/gin"
)

func Register(r *gin.Engine) {
	r.GET("/ping", ping)
	swagger.Register(r)
	search_api.Register(r)
}

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}