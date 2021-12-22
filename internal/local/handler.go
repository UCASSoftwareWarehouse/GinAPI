package local

import (
	"GinAPI/api/api_format"
	AErr "GinAPI/err"
	"GinAPI/internal/local/code_sim_uploader"
	"GinAPI/internal/local/model"
	"GinAPI/pb_gen"
	"GinAPI/rpc_cli"
	"GinAPI/util"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"mime/multipart"
	"os"
	"time"
)

type uploadSourceCodeProcessor interface {
	Process(ctx context.Context, projectName, tag, filepath string, fileType util.ArchiveFileType) *AErr.APIErr
	GetName() string
}

type uploadBinaryProcessor interface {
	// Process TODO 假想实现
	Process(ctx context.Context, projectName, tag, filepath string) *AErr.APIErr
	GetName() string
}

type Handler struct {
	uploadSourceCodeProcessors []uploadSourceCodeProcessor
	uploadBinaryProcessors     []uploadBinaryProcessor
}

var handler = &Handler{
	uploadSourceCodeProcessors: []uploadSourceCodeProcessor{
		// TODO 具有优先顺序。能够继续添加其他的Processor。报错时终止调用链。
		code_sim_uploader.Uploader{},
	},
	uploadBinaryProcessors: []uploadBinaryProcessor{
		// TODO
	},
}

func GetHandler() *Handler {
	return handler
}

// PatchSourceCode
// @Summary local上传源代码
// @Description
// @Tags source_code
// @Accept multipart/form-data
// @Param file formData file true "file"
// @Param projectName formData string true "projectName"
// @Param tag formData string true "tag"
// @Param fileType formData string true "fileType"
// @Produce json
// @Router /api/v1/source_code/local [patch]
// @Success 200 {object} model.LocalPatchSourceCodeResponse
func (h *Handler) PatchSourceCode(c *gin.Context) (*api_format.JSONRespFormat, *AErr.APIErr) {
	projectName := c.PostForm("projectName")
	tag := c.PostForm("tag")
	fileType := util.ArchiveFileType(c.PostForm("fileType"))
	file, err := c.FormFile("file")
	if projectName == "" || tag == "" || fileType == "" || err != nil {
		e := AErr.BadRequestErr.CustomMessageF("projectName or tag or fileType must not be nil, and file cannot be nil, err=[%v]", err)
		return nil, e
	}
	if !util.ArchiveUtilInstance.Support(fileType) {
		e := AErr.BadRequestErr.CustomMessageF("fileType is not supported, fileType is %s", fileType)
		return nil, e
	}

	fp, aErr := h.receiveFile(c, file)
	if aErr != nil {
		return nil, aErr
	}
	defer func() {
		_ = os.Remove(fp)
	}()
	for _, processor := range h.uploadSourceCodeProcessors {
		log.Printf("PatchSourceCode, ready to use %s processor", processor.GetName())
		e := processor.Process(c, projectName, tag, fp, fileType)
		if e != nil {
			log.Printf("PatchSourceCode, processor %s execution failed", processor.GetName())
			api_format.UnwrapErr(c, e)
			return nil, e
		}
		log.Printf("PatchSourceCode, %s processor finished execution", processor.GetName())
	}

	return api_format.SimpleOKResp(&model.LocalPatchSourceCodeResponse{}), nil
}

// PatchBinary
// @Summary 本地上传二进制包
// @Description
// @Tags binary
// @Accept multipart/form-data
// @Param file formData file true "file"
// @Param projectName formData string true "projectName"
// @Param tag formData string true "tag"
// @Produce json
// @Router /api/v1/binary/local [patch]
// @Success 200 {object} model.LocalPatchBinaryResponse
func (h *Handler) PatchBinary(c *gin.Context) (*api_format.JSONRespFormat, *AErr.APIErr) {
	projectName := c.PostForm("projectName")
	tag := c.PostForm("tag")
	file, err := c.FormFile("file")
	if projectName == "" || tag == "" || err != nil {
		e := AErr.BadRequestErr.CustomMessageF("projectName or tag must not be nil, and file cannot be nil, err=[%v]", err)
		return nil, e
	}

	fp, aErr := h.receiveFile(c, file)
	if aErr != nil {
		return nil, aErr
	}
	defer func() {
		_ = os.Remove(fp)
	}()
	for _, processor := range h.uploadBinaryProcessors {
		log.Printf("PatchBinary, ready to use %s processor", processor.GetName())
		e := processor.Process(c, projectName, tag, fp)
		if e != nil {
			log.Printf("PatchBinary, processor %s execution failed", processor.GetName())
			api_format.UnwrapErr(c, e)
			return nil, e
		}
		log.Printf("PatchBinary, %s processor finished execution", processor.GetName())
	}

	return api_format.SimpleOKResp(&model.LocalPatchSourceCodeResponse{}), nil
}

func (h *Handler) receiveFile(c *gin.Context, file *multipart.FileHeader) (string, *AErr.APIErr) {
	fo, err := os.CreateTemp("", fmt.Sprintf("api_upload_temp_%d", time.Now().UnixNano()))
	if err != nil {
		e := AErr.InternalErr.CustomMessageF("create temp file failed")
		return "", e
	}
	if err := c.SaveUploadedFile(file, fo.Name()); err != nil {
		e := AErr.BadRequestErr.CustomMessageF("upload file failed")
		return "", e
	}
	log.Printf("PatchSourceCode, save to temp file path=[%s]", fo.Name())
	return fo.Name(), nil
}

/*
RemoteCode here
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
