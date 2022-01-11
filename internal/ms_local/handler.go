package ms_local

import (
	"GinAPI/api/api_format"
	AErr "GinAPI/err"
	"GinAPI/internal/ms_local/model"
	"GinAPI/internal/ms_local/ms_local_downloader"
	"GinAPI/internal/ms_local/ms_local_uploader"
	"GinAPI/pb_gen"
	"GinAPI/rpc_cli"
	"context"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"sort"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

//type uploadSourceCodeProcessor interface {
//	Process(ctx context.Context, projectName, tag, filepath string, fileType util.ArchiveFileType) *AErr.APIErr
//	GetName() string
//}

type MSLocalUploadProcessor interface {
	// Process
	Process(ctx context.Context, fpath string, metadata *pb_gen.UploadMetadata) *AErr.APIErr
	GetName() string
}

type MSLocalDownloadProcessor interface {
	Process(ctx context.Context, req *pb_gen.DownloadRequest) (string, *pb_gen.DownloadMetadate, *AErr.APIErr)
	GetName() string
}

type Handler struct {
	MSLocalUploadProcessors []MSLocalUploadProcessor
	downloadProcessor       MSLocalDownloadProcessor
}

var handler = &Handler{

	MSLocalUploadProcessors: []MSLocalUploadProcessor{
		ms_local_uploader.Uploader{},
	},
	downloadProcessor: ms_local_downloader.Downloader{},
}

func GetHnadler() *Handler {
	return handler
}

func (h *Handler) Ping(c *gin.Context) (*api_format.JSONRespFormat, *AErr.APIErr) {
	MyHelloText := "Hello From Web Api"
	res, err := rpc_cli.MSLocalCli.SayHello(c, &pb_gen.HelloRequest{Name: MyHelloText})
	if err != nil {
		return nil, AErr.InternalErr.CustomMessage(fmt.Sprintf("Ping Code Sim Failed, err=[%s]", err))
	}
	return api_format.SimpleOKResp(&model.PingMSLocalResponse{
		FullText: fmt.Sprintf("MyHelloText is %s, ThanksText is %s", MyHelloText, res.Message),
	}), nil
}

//post
func (h *Handler) CreateProject(c *gin.Context) (*api_format.JSONRespFormat, *AErr.APIErr) {

	//var req model.CreateProjectRequest
	//json提交需要bind json
	//err := c.BindJSON(&req)
	//reqStr, _ := json.Marshal(req)
	//log.Printf("create project, reqStr=[%s]", reqStr)
	//form

	projectName := c.PostForm("projectName")
	if projectName == "" {
		return nil, AErr.BadRequestErr.CustomMessage("project name can't be empty!")
	}
	userId := c.PostForm("userId")
	tags := c.DefaultPostForm("tags", "")
	license := c.DefaultPostForm("license", "")
	projectDescription := c.DefaultPostForm("projectDescription", "")
	operatingSystem := c.DefaultPostForm("operatingSystem", pb_gen.OperatingSystem_name[0])
	programmingLanguage := c.DefaultPostForm("programmingLanguage", pb_gen.ProgrammingLanguage_name[0])
	naturalLanguage := c.DefaultPostForm("naturalLanguage", pb_gen.NaturalLanguage_name[0])
	topic := c.DefaultPostForm("topic", pb_gen.Topic_name[0])
	uid, _ := strconv.ParseUint(userId, 10, 64)
	osValue := uint8(pb_gen.OperatingSystem_value[operatingSystem])
	plValue := uint8(pb_gen.ProgrammingLanguage_value[programmingLanguage])
	nlValue := uint8(pb_gen.NaturalLanguage_value[naturalLanguage])
	toValue := uint8(pb_gen.Topic_value[topic])
	res, err := rpc_cli.MSLocalCli.CreateProject(context.Background(), &pb_gen.CreateProjectRequest{
		ProjectName:        projectName,
		UserId:             uid,
		Tags:               tags,
		License:            license,
		ProjectDescription: projectDescription,
		Classifiers:        model.GetClassifier(osValue, plValue, nlValue, toValue),
	})
	if err != nil {
		return nil, AErr.WrapRPCError(err)
	}
	log.Printf("project info is %v", res.ProjectInfo)
	pro := &model.ProjectDetail{
		Id:                  (*res.ProjectInfo).Id,
		ProjectName:         (*res.ProjectInfo).ProjectName,
		UserId:              (*res.ProjectInfo).UserId,
		Tags:                (*res.ProjectInfo).Tags,
		License:             (*res.ProjectInfo).License,
		Updatetime:          (*res.ProjectInfo).Updatetime,
		ProjectDescription:  (*res.ProjectInfo).ProjectDescription,
		CodeAddr:            (*res.ProjectInfo).CodeAddr,
		BinaryAddr:          (*res.ProjectInfo).BinaryAddr,
		OperatingSystem:     model.GetOSName((*res.ProjectInfo).Classifiers),
		ProgrammingLanguage: model.GetPLName((*res.ProjectInfo).Classifiers),
		NaturalLanguage:     model.GetNLName((*res.ProjectInfo).Classifiers),
		Topic:               model.GetToName((*res.ProjectInfo).Classifiers),
	}
	return api_format.SimpleOKResp(pro), nil
}

func (h *Handler)GetAllProjects(c *gin.Context)(*api_format.JSONRespFormat, *AErr.APIErr){
	page, _ := strconv.ParseUint(c.DefaultQuery("page", "1"), 10, 32)
	limit, _ := strconv.ParseUint(c.DefaultQuery("limit", "10"), 10, 32)

	stream, err := rpc_cli.MSLocalCli.GetUserProjects(c,&pb_gen.GetUserProjectsRequest{
		Uid:   1,
		Page:  uint32(page),
		Limit: uint32(limit),
	})
	if err != nil {
		return nil, AErr.WrapRPCError(err)
	}
	var pros []model.ProjectDetail
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("receive failed, err=[%v]", err)
			return nil, AErr.WrapRPCError(err)
		}
		log.Printf("Search info=[%v]", res.ProjectInfo)
		pro := &model.ProjectDetail{
			Id:                  (*res.ProjectInfo).Id,
			ProjectName:         (*res.ProjectInfo).ProjectName,
			UserId:              (*res.ProjectInfo).UserId,
			Tags:                (*res.ProjectInfo).Tags,
			License:             (*res.ProjectInfo).License,
			Updatetime:          (*res.ProjectInfo).Updatetime,
			ProjectDescription:  (*res.ProjectInfo).ProjectDescription,
			CodeAddr:            (*res.ProjectInfo).CodeAddr,
			BinaryAddr:          (*res.ProjectInfo).BinaryAddr,
			OperatingSystem:     model.GetOSName((*res.ProjectInfo).Classifiers),
			ProgrammingLanguage: model.GetPLName((*res.ProjectInfo).Classifiers),
			NaturalLanguage:     model.GetNLName((*res.ProjectInfo).Classifiers),
			Topic:               model.GetToName((*res.ProjectInfo).Classifiers),
		}

		pros = append(pros,*pro)
	}
	log.Printf("get user's files success")
	return api_format.SimpleOKResp(pros), nil

}
func (h *Handler) GetProject(c *gin.Context) (*api_format.JSONRespFormat, *AErr.APIErr) {
	pid, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return nil, AErr.BadRequestErr
	}

	res, err := rpc_cli.MSLocalCli.GetProject(c, &pb_gen.GetProjectRequest{
		Pid: pid,
		Uid: 0,
	})
	if err != nil {
		return nil, AErr.WrapRPCError(err)
	}
	log.Printf("project info is:%v", res)
	pro := &model.ProjectDetail{
		Id:                  (*res.ProjectInfo).Id,
		ProjectName:         (*res.ProjectInfo).ProjectName,
		UserId:              (*res.ProjectInfo).UserId,
		Tags:                (*res.ProjectInfo).Tags,
		License:             (*res.ProjectInfo).License,
		Updatetime:          (*res.ProjectInfo).Updatetime,
		ProjectDescription:  (*res.ProjectInfo).ProjectDescription,
		CodeAddr:            (*res.ProjectInfo).CodeAddr,
		BinaryAddr:          (*res.ProjectInfo).BinaryAddr,
		OperatingSystem:     model.GetOSName((*res.ProjectInfo).Classifiers),
		ProgrammingLanguage: model.GetPLName((*res.ProjectInfo).Classifiers),
		NaturalLanguage:     model.GetNLName((*res.ProjectInfo).Classifiers),
		Topic:               model.GetToName((*res.ProjectInfo).Classifiers),
	}
	return api_format.SimpleOKResp(pro), nil
}


func (h *Handler) SearchProjects(c *gin.Context) (*api_format.JSONRespFormat, *AErr.APIErr) {
	page, _ := strconv.ParseUint(c.DefaultQuery("page", "1"), 10, 32)
	limit, _ := strconv.ParseUint(c.DefaultQuery("limit", "10"), 10, 32)
	keyword := c.DefaultQuery("keyword", "")
	operatingSystem := c.DefaultQuery("operatingSystem", pb_gen.OperatingSystem_name[0])
	programmingLanguage := c.DefaultQuery("programmingLanguage", pb_gen.ProgrammingLanguage_name[0])
	naturalLanguage := c.DefaultQuery("naturalLanguage", pb_gen.NaturalLanguage_name[0])
	topic := c.DefaultQuery("topic", pb_gen.Topic_name[0])
	osValue := uint8(pb_gen.OperatingSystem_value[operatingSystem])
	plValue := uint8(pb_gen.ProgrammingLanguage_value[programmingLanguage])
	nlValue := uint8(pb_gen.NaturalLanguage_value[naturalLanguage])
	toValue := uint8(pb_gen.Topic_value[topic])
	req := &pb_gen.SearchProjectRequest{
		Page:        uint32(page),
		Limit:       uint32(limit),
		KeyWord:     keyword,
		Classifiers: model.GetClassifier(osValue, plValue, nlValue, toValue),
	}
	stream, err := rpc_cli.MSLocalCli.SearchProject(c, req)
	if err != nil {
		log.Println("cann't start search, err=[%v]", err)
		return nil, AErr.WrapRPCError(err)
	}
	var pros []model.ProjectDetail
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("receive failed, err=[%v]", err)
			return nil, AErr.WrapRPCError(err)
		}
		pro := &model.ProjectDetail{
			Id:                  (*res.ProjectInfo).Id,
			ProjectName:         (*res.ProjectInfo).ProjectName,
			UserId:              (*res.ProjectInfo).UserId,
			Tags:                (*res.ProjectInfo).Tags,
			License:             (*res.ProjectInfo).License,
			Updatetime:          (*res.ProjectInfo).Updatetime,
			ProjectDescription:  (*res.ProjectInfo).ProjectDescription,
			CodeAddr:            (*res.ProjectInfo).CodeAddr,
			BinaryAddr:          (*res.ProjectInfo).BinaryAddr,
			OperatingSystem:     model.GetOSName((*res.ProjectInfo).Classifiers),
			ProgrammingLanguage: model.GetPLName((*res.ProjectInfo).Classifiers),
			NaturalLanguage:     model.GetNLName((*res.ProjectInfo).Classifiers),
			Topic:               model.GetToName((*res.ProjectInfo).Classifiers),
		}
		log.Printf("Search info=[%v]", res.ProjectInfo)
		pros = append(pros, *pro)
	}
	return api_format.SimpleOKResp(pros), nil
}

//list code files
func (h *Handler) GetCodes(c *gin.Context) (*api_format.JSONRespFormat, *AErr.APIErr) {
	pid, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return nil, AErr.BadRequestErr.CustomMessageF("can't get pid, err=[%v]", err)
	}
	dtype, err := strconv.ParseInt(c.Query("dtype"), 10, 32) //dtype 类型
	fid := c.Query("fileId")
	if fid == "" {
		return nil, AErr.BadRequestErr.CustomMessage("file id can't be empty")
	}
	if pb_gen.FileType(dtype) == pb_gen.FileType_code_dir {
		page, _ := strconv.ParseUint(c.DefaultQuery("page", "1"), 10, 32)
		limit, _ := strconv.ParseUint(c.DefaultQuery("limit", "10"), 10, 32)

		req := &pb_gen.GetCodesRequest{
			Pid:   pid,
			Uid:   1,
			Fid:   fid,
			Page:  uint32(page),
			Limit: uint32(limit),
		}
		return h.receiveCodesDir(c, req)
	} else {

		return nil, AErr.BadRequestErr.CustomMessage("GetCodes: file type need to be code dir")
	}
	return nil, nil
}

func (h *Handler) Delete(c *gin.Context) (*api_format.JSONRespFormat, *AErr.APIErr) {
	pid, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		log.Printf("can't get pid from parameters")
		return nil, AErr.BadRequestErr
	}
	req := &pb_gen.DeleteProjectRequest{
		Pid:      pid,
		Uid:      0,
		FileType: pb_gen.FileType_project,
	}
	dtype, err := strconv.ParseInt(c.Query("dtype"), 10, 32) //dtype 类型
	if pb_gen.FileType(dtype) == pb_gen.FileType_codes {
		req.FileType = pb_gen.FileType_codes
	} else if pb_gen.FileType(dtype) == pb_gen.FileType_binary {
		req.FileType = pb_gen.FileType_binary
	} else if pb_gen.FileType(dtype) == pb_gen.FileType_project {
		req.FileType = pb_gen.FileType_project
	} else {
		return nil, AErr.BadRequestErr.CustomMessage("delete file type is wrong")
	}
	log.Printf("delete fid:%d,filetype:%v ", pid, req.FileType)
	res, err := rpc_cli.MSLocalCli.DeleteProject(c, req)
	if err != nil {
		log.Printf("gprc delete error, err=[%v]", err)
		return nil, AErr.WrapRPCError(err)
	}
	return api_format.SimpleOKResp(res.Message), nil
}

func (h *Handler) Download(c *gin.Context) {
	pid, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"msg":     fmt.Sprintf("DownloadBinary: parse pid filed, err=[%v]", err),
		})
		return
	}
	fid := c.Query("fileId")
	log.Printf("Download fid is %s", fid)
	if fid == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"msg":     "file id can't be empty",
		})
		return
	}
	dtype, err := strconv.ParseInt(c.Query("dtype"), 10, 32) //dtype 类型
	ftype := pb_gen.FileType(dtype)
	if ftype == pb_gen.FileType_code_dir {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"msg":     "can only download project/binary/codes",
		})
		return
	}

	//get file by grpc
	req := &pb_gen.DownloadRequest{
		FileId:    fid,
		UserId:    1,
		FileType:  ftype,
		ProjectId: pid,
	}
	fpath, metadata, aerr := h.downloadProcessor.Process(c, req)
	defer func() {
		err = os.Remove(fpath)
		if err != nil {
			log.Printf("delete file error, err=[%v]", err)
		}
	}()
	if aerr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"msg":     fmt.Sprintf("Download: download processor failed, err=[%v]", aerr),
		})
		return
	}

	log.Println("download metadata:%v", metadata.FileInfo)
	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Transfer-Encoding", "binary")
	c.Header("Content-Disposition", "attachment; filename="+metadata.FileInfo.FileName)
	c.Header("Content-Type", "application/octet-stream")
	c.File(fpath)
}

func (h *Handler) Upload(c *gin.Context) (*api_format.JSONRespFormat, *AErr.APIErr) {
	pid, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return nil, AErr.BadRequestErr.CustomMessageF("can't get project id, err=[%v]", err)
	}
	fileName := c.PostForm("fileName")
	if fileName == "" {
		return nil, AErr.BadRequestErr.CustomMessage("file name can't be empty")
	}
	log.Printf("Upload: file name is %s", fileName)
	dtype, err := strconv.ParseInt(c.PostForm("dtype"), 10, 32) //dtype 类型
	ftype := pb_gen.FileType(dtype)
	if ftype != pb_gen.FileType_binary && ftype != pb_gen.FileType_codes {
		return nil, AErr.BadRequestErr.CustomMessage("can only upload binary/codes")
	}
	log.Printf("Upload File name is %s, type is %v", fileName, ftype)

	metadata := &pb_gen.UploadMetadata{
		ProjectId: pid,
		UserId:    1,
		FileInfo: &pb_gen.FileInfo{
			FileName: fileName,
			FileType: ftype,
		},
	}
	file, err := c.FormFile("file")
	if err != nil {
		e := AErr.BadRequestErr.CustomMessageF("err=[%v]", err)
		return nil, e
	}
	fp, aErr := h.receiveFile(c, file)
	defer func() {
		err = os.Remove(fp)
		if err != nil {
			log.Printf("delete file error, err=[%v]", err)
		}
	}()
	if aErr != nil {
		return nil, aErr
	}

	for _, processor := range h.MSLocalUploadProcessors {
		log.Printf("update file, ready to use %s processor", processor.GetName())
		e := processor.Process(c, fp, metadata)
		if e != nil {
			log.Printf("processor %s upload file failed", processor.GetName())
			api_format.UnwrapErr(c, e)
			return nil, e
		}
		log.Printf("update file, %s processor finished execution", processor.GetName())
	}

	return api_format.SimpleOKResp(&model.MSLocalUploadRequest{}), nil
}

func (h *Handler) receiveFile(c *gin.Context, file *multipart.FileHeader) (string, *AErr.APIErr) {
	fo, err := os.CreateTemp("", fmt.Sprintf("api_upload_temp_%d", time.Now().UnixNano()))

	if err != nil {
		log.Printf("create temp file failed, err=[%v]", err)
		e := AErr.InternalErr.CustomMessageF("create temp file failed")
		return "", e
	}
	log.Printf("create temp file in %s", fo.Name())
	if err := c.SaveUploadedFile(file, fo.Name()); err != nil {
		e := AErr.BadRequestErr.CustomMessageF("upload file failed")
		return "", e
	}
	log.Printf("PatchSourceCode, save to temp file path=[%s]", fo.Name())
	return fo.Name(), nil
}

func (h *Handler) receiveCodesDir(c *gin.Context, req *pb_gen.GetCodesRequest) (*api_format.JSONRespFormat, *AErr.APIErr) {
	stream, err := rpc_cli.MSLocalCli.GetCodes(c, req)
	if err != nil {
		log.Printf("can't get the content of codes, err=[%v]", err)
		return nil, AErr.WrapRPCError(err)
	}
	var codes Files
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("receive failed, err=[%v]", err)
			return nil, AErr.WrapRPCError(err)
		}
		codes = append(codes, *res)
	}
	sort.Sort(codes)
	//sort to do
	return api_format.SimpleOKResp(codes), nil
}

type Files []pb_gen.GetCodesResponse

func (files Files) Len() int {
	return len(files)
}

func (files Files) Swap(i, j int) {
	files[i], files[j] = files[j], files[i]
}

func (files Files) Less(i, j int) bool {
	if files[i].FileInfo.FileType < files[j].FileInfo.FileType {
		return true
	} else if files[i].FileInfo.FileType > files[j].FileInfo.FileType {
		return false
	} else {
		return files[i].FileInfo.FileName < files[j].FileInfo.FileName
	}
}
