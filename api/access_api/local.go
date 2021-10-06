package access_api

import (
	"GinAPI/api/api_format"
	"GinAPI/err"
	"GinAPI/internal/local"
	"github.com/gin-gonic/gin"
)

type localAPI struct{}

func (localAPI) localUploadSourceCode() api_format.SimpleJSONHandler {
	return func(c *gin.Context) (*api_format.JSONRespFormat, *err.APIErr) {
		return local.GetHandler().UploadSourceCode(c)
	}
}

func (localAPI) localUploadBinary() api_format.NormalHandler {
	return func(c *gin.Context) {

	}
}
