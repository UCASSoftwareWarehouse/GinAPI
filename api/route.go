package api

import (
	"GinAPI/api/api_format"
	"GinAPI/err"
	"GinAPI/internal/local"
	"GinAPI/internal/search"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func Register(r *gin.Engine) {
	r.GET("/ping", ping)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	rg := r.Group("/api/v1")

	sourceCodeRouter := rg.Group(prefixSourceCode)
	binaryRouter := rg.Group(prefixBinary)
	testRouter := rg.Group(prefixTest)

	sourceCodeAPI := sourceCodeAPI{}
	sourceCodeRouter.PATCH("local", api_format.Wrap(sourceCodeAPI.localPatchSourceCode()))
	sourceCodeRouter.GET("search", api_format.Wrap(sourceCodeAPI.search()))

	binaryAPI := binaryAPI{}
	binaryRouter.PATCH("local", api_format.Wrap(binaryAPI.localPatchBinary()))

	testAPI := testAPI{}
	testRouter.GET("ping_code_sim", api_format.Wrap(testAPI.pingCodeSim()))
	testRouter.GET("error_handler", api_format.Wrap(testAPI.testErrorHandler()))
}

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

const (
	prefixSourceCode = "source_code"
	prefixBinary     = "binary"
	prefixProject = "project"
	prefixTest    = "test"
)

type sourceCodeAPI struct{}

func (sourceCodeAPI) localPatchSourceCode() api_format.JSONHandler {
	return func(c *gin.Context) (*api_format.JSONRespFormat, *err.APIErr) {
		return local.GetHandler().PatchSourceCode(c)
	}
}

func (sourceCodeAPI) search() api_format.JSONHandler {
	return func(c *gin.Context) (*api_format.JSONRespFormat, *err.APIErr) {
		return search.GetHandler().SearchSourceCode(c)
	}
}

type binaryAPI struct{}

func (binaryAPI) localPatchBinary() api_format.JSONHandler {
	return func(c *gin.Context) (*api_format.JSONRespFormat, *err.APIErr) {
		return local.GetHandler().PatchBinary(c)
	}
}

type testAPI struct{}

func (testAPI) pingCodeSim() api_format.JSONHandler {
	return func(c *gin.Context) (*api_format.JSONRespFormat, *err.APIErr) {
		return search.GetHandler().Ping(c)
	}
}

// Ping
// @Summary test error handler
// @Tags test
// @Produce json
// @Router /api/v1/test/error_handler [get]
// @Success 200
// @Fail err.APIErr
func (testAPI) testErrorHandler() api_format.JSONHandler {
	return func(c *gin.Context) (*api_format.JSONRespFormat, *err.APIErr) {
		_ = c.AbortWithError(err.BadRequestErr.Status, err.BadRequestErr)
		return nil, nil
	}
}
