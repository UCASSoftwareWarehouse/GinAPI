package api

import (
	"GinAPI/api/api_format"
	"GinAPI/err"
	"GinAPI/internal/local"
	"GinAPI/internal/ms_local"
	"GinAPI/internal/project"
	"GinAPI/internal/remote"
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
	projectRouter := rg.Group(prefixProject)

	sourceCodeAPI := sourceCodeAPI{}
	sourceCodeRouter.PATCH("local", api_format.Wrap(sourceCodeAPI.localPatchSourceCode()))
	sourceCodeRouter.GET("search", api_format.Wrap(sourceCodeAPI.search()))
	sourceCodeRouter.GET("pypi", api_format.Wrap(sourceCodeAPI.downloadPypi()))

	binaryAPI := binaryAPI{}
	binaryRouter.PATCH("local", api_format.Wrap(binaryAPI.localPatchBinary()))
	binaryRouter.GET("apt", api_format.Wrap(binaryAPI.downloadApt()))

	testAPI := testAPI{}
	testRouter.GET("ping_code_sim", api_format.Wrap(testAPI.pingCodeSim()))
	testRouter.GET("error_handler", api_format.Wrap(testAPI.testErrorHandler()))

	projectAPI := projectAPI{}
	projectRouter.DELETE(":projectName/:tag", api_format.Wrap(projectAPI.delete()))

	//ms_local
	localProjectRouter := rg.Group(prefixLocalProject)
	localProjectAPI := LocalProjectAPI{}
	localProjectRouter.POST("create", api_format.Wrap(localProjectAPI.CreateProject()))
	localProjectRouter.GET("/:id", api_format.Wrap(localProjectAPI.GetProject()))
	localProjectRouter.POST("/:id/upload", api_format.Wrap(localProjectAPI.Upload()))
	localProjectRouter.GET("/:id/download", api_format.Wrap(localProjectAPI.Download()))
	localProjectRouter.GET("/search", api_format.Wrap(localProjectAPI.Search()))
	localProjectRouter.GET("/:id/codes", api_format.Wrap(localProjectAPI.GetCodes()))
	localProjectRouter.DELETE("/:id/delete", api_format.Wrap(localProjectAPI.Delete()))
}

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

const (
	prefixSourceCode   = "source_code"
	prefixBinary       = "binary"
	prefixProject      = "project"
	prefixTest         = "test"
	prefixLocalProject = "projects"
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

func (sourceCodeAPI) downloadPypi() api_format.JSONHandler {
	return func(c *gin.Context) (*api_format.JSONRespFormat, *err.APIErr) {
		return remote.GetHandler().DownloadPypi(c)
	}
}

type binaryAPI struct{}

func (binaryAPI) localPatchBinary() api_format.JSONHandler {
	return func(c *gin.Context) (*api_format.JSONRespFormat, *err.APIErr) {
		return local.GetHandler().PatchBinary(c)
	}
}

func (binaryAPI) downloadApt() api_format.JSONHandler {
	return func(c *gin.Context) (*api_format.JSONRespFormat, *err.APIErr) {
		return remote.GetHandler().DownloadApt(c)
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

type projectAPI struct{}

func (projectAPI) delete() api_format.JSONHandler {
	return func(ctx *gin.Context) (*api_format.JSONRespFormat, *err.APIErr) {
		return project.GetHandler().Delete(ctx)
	}
}

type LocalProjectAPI struct{}

func (LocalProjectAPI) CreateProject() api_format.JSONHandler {
	return func(ctx *gin.Context) (*api_format.JSONRespFormat, *err.APIErr) {
		return ms_local.GetHnadler().CreateProject(ctx)
	}
}

func (LocalProjectAPI) GetProject() api_format.JSONHandler {
	return func(ctx *gin.Context) (*api_format.JSONRespFormat, *err.APIErr) {
		return ms_local.GetHnadler().GetProject(ctx)
	}
}

func (LocalProjectAPI) Delete() api_format.JSONHandler {
	return func(ctx *gin.Context) (*api_format.JSONRespFormat, *err.APIErr) {
		return ms_local.GetHnadler().Delete(ctx)
	}
}

func (LocalProjectAPI) Upload() api_format.JSONHandler {
	return func(ctx *gin.Context) (*api_format.JSONRespFormat, *err.APIErr) {
		return ms_local.GetHnadler().Upload(ctx)
	}
}

func (LocalProjectAPI) Download() api_format.NormalHandler {
	return func(ctx *gin.Context) {
		ms_local.GetHnadler().Download(ctx)
	}
}

func (LocalProjectAPI) Search() api_format.JSONHandler {
	return func(ctx *gin.Context) (*api_format.JSONRespFormat, *err.APIErr) {
		return ms_local.GetHnadler().SearchProjects(ctx)
	}
}

func (LocalProjectAPI) GetCodes() api_format.JSONHandler {
	return func(ctx *gin.Context) (*api_format.JSONRespFormat, *err.APIErr) {
		return ms_local.GetHnadler().GetCodes(ctx)
	}
}
