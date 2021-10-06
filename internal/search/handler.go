package search

import (
	"GinAPI/api/api_format"
	AErr "GinAPI/err"
	"GinAPI/internal/common_model"
	"GinAPI/internal/search/model"
	"GinAPI/pb_gen"
	"GinAPI/rpc_cli"
	"fmt"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

type Handler struct{}

var handler = &Handler{}

func GetHandler() *Handler {
	return handler
}

var validCodeTypes = []pb_gen.CodeSimSearchRequest_CodeType{
	pb_gen.CodeSimSearchRequest_python,
}

func isValidCodeType(codeType pb_gen.CodeSimSearchRequest_CodeType) bool {
	for _, ct := range validCodeTypes {
		if ct == codeType {
			return true
		}
	}
	return false
}

// Ping
// @Summary Print
// @Accept json
// @Tags Search
// @Produce json
// @Router /search/ping [get]
// @Success 200 {object} PingCodeSimResponse
func (h Handler) Ping(c *gin.Context) (*api_format.JSONRespFormat, *AErr.APIErr) {
	myHelloText := "HelloFromWebAPI!"
	resp, mErr := rpc_cli.CodeSimCli.HelloWorld(c, &pb_gen.CodeSimHelloWorldRequest{
		HelloText: myHelloText,
	})
	if mErr != nil {
		return nil, AErr.InternalErr.CustomMessage(fmt.Sprintf("Ping Code Sim Failed, err=[%s]", mErr))
	}
	return api_format.SimpleOKResp(&model.PingCodeSimResponse{
		FullText: fmt.Sprintf("MyHelloText is %s, ThanksText is %s", myHelloText, resp.GetThanksText()),
	}), nil
}

// Query
// @Summary 模糊查询代码
// @Accept json
// @Tags Search
// @param queryCodeRequest body QueryCodeRequest true "query code"
// @Produce json
// @Router /search/query [post]
// @Success 200 {object} PingCodeSimResponse
func (h Handler) Query(c *gin.Context) (*api_format.JSONRespFormat, *AErr.APIErr) {
	req := &model.QueryCodeRequest{}
	err := c.ShouldBindJSON(req)
	if err != nil {
		return nil, AErr.BadRequestErr
	}
	if !isValidCodeType(req.CodeType) {
		req.CodeType = pb_gen.CodeSimSearchRequest_unknown
	}
	res, err := rpc_cli.CodeSimCli.Search(c, &pb_gen.CodeSimSearchRequest{
		MatchText:       req.Content,
		MatchTextIsCode: isValidCodeType(req.CodeType),
		CodeType:        req.CodeType,
		Limit:           int32(req.Size),
		Offset:          int32(req.From),
	})
	if err != nil {
		// TODO 使用status透传下游错误，之后需要统一处理rpc错误。
		errStatus, ok := status.FromError(err)
		if !ok {
			return nil, AErr.InternalErr
		}
		log.Println(errStatus.Message())
		log.Println(errStatus.Code())
		if codes.InvalidArgument == errStatus.Code() {
			return nil, AErr.BadRequestErr.CustomMessage(errStatus.Message())
		}
	}
	packedFiles := h.packFiles(res.GetFiles())
	return api_format.SimpleOKResp(&model.QueryCodeResponse{
		ProjectFiles: packedFiles,
	}), nil
}

func (h Handler) packFiles(files []*pb_gen.CodeSimProjectFile) []*common_model.ProjectFile {
	packed := make([]*common_model.ProjectFile, 0, len(files))
	for _, f := range files {
		packed = append(packed, &common_model.ProjectFile{
			Project: &common_model.Project{
				ProjectName: f.GetProjectInfo().GetProjectName(),
				Tag:         f.GetProjectInfo().GetTag(),
			},
			RelPath: f.GetRelativePath(),
		})
	}
	// TODO get contents of each project file
	return packed
}
