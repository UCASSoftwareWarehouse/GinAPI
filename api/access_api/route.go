package access_api

import (
	"GinAPI/api/api_format"
	"github.com/gin-gonic/gin"
)

const (
	accessPrefix = "access"
	localPrefix  = "local"
)

func Register(r *gin.Engine) {
	accessRouter := r.Group(accessPrefix)

	localRouter := accessRouter.Group(localPrefix)
	localAPI := localAPI{}
	localRouter.POST("/upload/source_code", api_format.Wrap(localAPI.localUploadSourceCode()))
	localRouter.POST("/upload/binary", api_format.Wrap(localAPI.localUploadBinary()))
}
