package remote

import (
	"GinAPI/api/api_format"
	AErr "GinAPI/err"
	"GinAPI/pb_gen"
	"GinAPI/rpc_cli"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Handler struct{}

var handler = &Handler{}

func GetHandler() *Handler {
	return handler
}

/*
MS_RemoteCode here
*/
func (h *Handler) HelloWorld(c *gin.Context) (*api_format.JSONRespFormat, *AErr.APIErr) {
	hello := c.Query("hello")
	println("hello=", hello)
	request := &pb_gen.HelloWorldRequest{
		HelloText: hello,
	}
	resp, _ := rpc_cli.RemoteCodeCli.HelloWorld(context.Background(), request)
	return api_format.SimpleOKResp(resp.ThanksText), nil
}

func (h *Handler) DownloadPypi(c *gin.Context) (*api_format.JSONRespFormat, *AErr.APIErr) {
	noDeps := c.Query("no_deps")
	noDepsBool, _ := strconv.ParseBool(noDeps)
	userId := c.Query("userId")
	if len(userId) == 0 {
		return api_format.NewJSONResp(http.StatusOK, "用户id不能为空", nil), nil
	}
	parseUserId, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		return api_format.NewJSONResp(http.StatusOK, "invalid userId", nil), nil
	}
	projectId := c.Query("projectId")
	if len(projectId) == 0 {
		return api_format.NewJSONResp(http.StatusOK, "项目id不能为空", nil), nil
	}
	parseProjectId, err := strconv.ParseInt(projectId, 10, 64)
	if err != nil {
		return api_format.NewJSONResp(http.StatusOK, "invalid projectId", nil), nil
	}
	packageName := c.Query("package")
	if len(packageName) == 0 {
		return api_format.NewJSONResp(http.StatusOK, "请输入包名", nil), nil
	}
	request := &pb_gen.DownloadRemoteCodeRequest{
		Platform:   c.Query("platform"),
		NoDeps:     noDepsBool,
		OnlyBinary: c.Query("only_binary"),
		//PythonVersion: "",
		Package: packageName,
		Version: c.Query("version"),
		Metadata: &pb_gen.UploadMetadata2{
			ProjectId: uint64(parseProjectId),
			UserId:    uint64(parseUserId),
			FileInfo: &pb_gen.FileInfo2{
				FileName: "",
				FileType: 0,
			},
		},
	}
	resp, _ := rpc_cli.RemoteCodeCli.DownloadRemoteCode(context.Background(), request)
	return api_format.SimpleOKResp(resp), nil
}

func (h *Handler) DownloadApt(c *gin.Context) (*api_format.JSONRespFormat, *AErr.APIErr) {

	userId := c.Query("userId")
	if len(userId) == 0 {
		return api_format.NewJSONResp(http.StatusOK, "用户id不能为空", nil), nil
	}
	parseUserId, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		return api_format.NewJSONResp(http.StatusOK, "invalid userId", nil), nil
	}
	projectId := c.Query("projectId")
	if len(projectId) == 0 {
		return api_format.NewJSONResp(http.StatusOK, "项目id不能为空", nil), nil
	}
	parseProjectId, err := strconv.ParseInt(projectId, 10, 64)
	if err != nil {
		return api_format.NewJSONResp(http.StatusOK, "invalid projectId", nil), nil
	}
	packageName := c.Query("package")
	if len(packageName) == 0 {
		return api_format.NewJSONResp(http.StatusOK, "请输入包名", nil), nil
	}
	codeType := c.Query("type")
	if len(codeType) == 0 {
		return api_format.NewJSONResp(http.StatusOK, "请输入文件类别", nil), nil
	}

	request := &pb_gen.DownloadAptDebRequest{
		Package: packageName,
		Version: c.Query("version"),
		Type:    codeType,
		Metadata: &pb_gen.UploadMetadata2{
			ProjectId: uint64(parseProjectId),
			UserId:    uint64(parseUserId),
			FileInfo: &pb_gen.FileInfo2{
				FileName: "",
				FileType: 0,
			},
		},
	}
	resp, _ := rpc_cli.RemoteCodeCli.DownloadAptDeb(context.Background(), request)
	return api_format.SimpleOKResp(resp), nil
}
