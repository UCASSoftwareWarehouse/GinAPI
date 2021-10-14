package code_sim_uploader

import (
	AErr "GinAPI/err"
	"GinAPI/pb_gen"
	"GinAPI/rpc_cli"
	"GinAPI/util"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

var (
	sentValue int64 = 1000000 //limit
)

var fileTypeMapper = map[util.ArchiveFileType]pb_gen.CodeSimUploadFileType{
	util.ArchiveFileTypeZip: pb_gen.CodeSimUploadFileType_code_sim_upload_file_type_zip,
}

type Uploader struct {
}

func (Uploader) GetName() string {
	return "CodeSimUploaderProcessor"
}

func (Uploader) Process(ctx context.Context, projectName, tag string, filepath string, fileType util.ArchiveFileType) *AErr.APIErr {
	proj := &pb_gen.CodeSimProject{
		ProjectName: projectName,
		Tag:         tag,
	}
	var (
		buf      []byte
		response *pb_gen.CodeSimUploadResponse
	)
	// open input file
	fi, err := os.Open(filepath)
	if err != nil {
		log.Println("Not able to open")
		return AErr.InternalErr
	}

	stat, err := fi.Stat()
	if err != nil {
		log.Printf("fi Stat failed, err=[%v]", err)
		return AErr.InternalErr
	}
	// close fi on exit and check for its returned error
	defer func() {
		if err := fi.Close(); err != nil {
			fmt.Println("Not able to open")
			return
		}
	}()

	stream, err := rpc_cli.CodeSimCli.Upload(ctx)
	if err != nil {
		return AErr.InternalErr.CustomMessageF("failed to create upload stream for file %+v", fi)
	}
	defer stream.CloseSend()

	buf = make([]byte, stat.Size())
	for {
		// read a chunk
		n, err := fi.Read(buf)
		if err != nil && err != io.EOF {
			return AErr.InternalErr.CustomMessage("failed to send chunk via stream")
		}
		if n == 0 {
			break
		}
		var i int64
		for i = 0; i < ((stat.Size() / sentValue) * sentValue); i += sentValue {
			err = stream.Send(&pb_gen.CodeSimUploadRequest{
				Project:      proj,
				FileType:     fileTypeMapper[fileType],
				ContentChunk: buf[i : i+sentValue],
				TotalSize:    strconv.FormatInt(stat.Size(), 10),
				Received:     strconv.FormatInt(i+sentValue, 10),
			})
		}
		if stat.Size()%sentValue > 0 {
			err = stream.Send(&pb_gen.CodeSimUploadRequest{
				Project:      proj,
				FileType:     fileTypeMapper[fileType],
				ContentChunk: buf[((stat.Size() / sentValue) * sentValue):((stat.Size() / sentValue * sentValue) + (stat.Size() % sentValue))],
				TotalSize:    strconv.FormatInt(stat.Size(), 10),
				Received:     strconv.FormatInt(stat.Size()%sentValue, 10),
			})
		}

		if err != nil {
			return AErr.InternalErr.CustomMessageF("failed to send chunk via stream, err=[%v]", err)
		}
	}

	response, err = stream.CloseAndRecv()
	if err != nil {
		return AErr.InternalErr.CustomMessageF("failed to receive upstream status response, err=[%v]", err)
	}

	if response.GetStatus() != pb_gen.CodeSimUploadStatus_code_sim_upload_status_OK {
		return AErr.InternalErr.CustomMessageF("upload failed, response=[%s], msg=[%s]", response, response.GetMessage())
	}

	resStr, _ := json.Marshal(response)
	log.Printf("PatchSourceCode success, response str is [%s]", resStr)

	return nil
}
