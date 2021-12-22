package project

import (
	"GinAPI/api/api_format"
	"GinAPI/err"
	"GinAPI/internal/search/model"
	"GinAPI/pb_gen"
	"GinAPI/rpc_cli"
	"github.com/gin-gonic/gin"
)

// Handler 对于Project的公共操作应该使用Project包作为入口。
// 例如upload, delete等。在local, remote下面应该定义的是service实例。
type Handler struct{}

var handler = &Handler{}

func GetHandler() *Handler {
	return handler
}

// Delete
// @Summary 删除Project
// @Tags project
// @Produce json
// @Router /api/v1/project/{projectName}/{tag} [delete]
// @param projectName path string true "projectName"
// @param tag path string true "tag"
// @Success 200 {object} model.ProjectDeleteResponse
func (h *Handler) Delete(c *gin.Context) (*api_format.JSONRespFormat, *err.APIErr) {
	projectName := c.Param("projectName")
	tag := c.Param("tag")
	if projectName == "" || tag == "" {
		return nil, err.BadRequestErr.CustomMessageF("删除请求参数异常，projectName=[%s], tag=[%s]", projectName, tag)
	}
	// 尝试向code_sim删除。
	_, e := rpc_cli.CodeSimCli.Delete(c, &pb_gen.CodeSimDeleteRequest{
		ProjectName: projectName,
		Tag:         tag,
	})
	if e != nil {
		return nil, err.WrapRPCError(e)
	}
	return api_format.SimpleOKResp(&model.DeleteSourceCodeResponse{}), nil
}