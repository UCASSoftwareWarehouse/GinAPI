package ms_local_uploader

import (
	AErr "GinAPI/err"
	"GinAPI/pb_gen"
	"GinAPI/rpc_cli"
	"bufio"

	//"GinAPI/util"
	"context"
	"io"
	"log"
	"os"
)

type Uploader struct {
}

func (Uploader) GetName() string {
	return "MSLocalUploader"
}

func (Uploader) Process(ctx context.Context, fpath string, metadata *pb_gen.UploadMetadata) *AErr.APIErr {
	file, err := os.Open(fpath)
	if err != nil {
		log.Printf("cannot open file: ", err)
		return AErr.InternalErr
	}
	defer file.Close()

	stream, err := rpc_cli.MSLocalCli.Upload(ctx)
	if err != nil {
		log.Printf("cannot upload file: ", err)
		return AErr.InternalErr.CustomMessageF("failed to create upload stream for file %s", fpath)
	}

	req := &pb_gen.UploadRequest{
		Data: &pb_gen.UploadRequest_Metadata{
			Metadata: metadata,
		},
	}
	log.Printf("request is:\n%v", req)
	err = stream.Send(req)
	if err != nil {
		log.Printf("cannot send image info to server: ", err, stream.RecvMsg(nil))
		return AErr.InternalErr.CustomMessageF("failed to send upload metadata via stream, err=[%v]", err)
	}
	reader := bufio.NewReader(file)
	buffer := make([]byte, 1024)

	for {
		n, err := reader.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("cannot read chunk to buffer: ", err)
			return AErr.InternalErr.CustomMessageF("failed to read data, err=[%v]", err)
		}
		req := &pb_gen.UploadRequest{
			Data: &pb_gen.UploadRequest_Content{
				Content: buffer[:n],
			},
		}

		err = stream.Send(req)
		if err != nil {
			log.Printf("cannot send chunk to server: ", err, stream.RecvMsg(nil))
			return AErr.InternalErr.CustomMessageF("failed to send upload chunkdata via stream, err=[%v]", err)
		}
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Printf("cannot receive response: ", err)
		return AErr.InternalErr.CustomMessageF("ailed to receive upstream status response, err=[%v]", err)
	}
	log.Printf("file info:\n%v", res.ProjectInfo)
	defer func() {
		err = os.Remove(fpath)
		if err != nil {
			log.Printf("Uploader Process: delete file error, err=[%v]", err)
		}
	}()
	return nil
}
