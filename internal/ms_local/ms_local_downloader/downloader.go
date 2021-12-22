package ms_local_downloader

import (
	AErr "GinAPI/err"
	"GinAPI/pb_gen"
	"GinAPI/rpc_cli"
	"context"
	"io"
	"log"
	"os"
)

type Downloader struct {
}

func (dl Downloader) GetName() string {
	return "MSLocalDownloader"
}

func (dl Downloader) Process(ctx context.Context, req *pb_gen.DownloadRequest) (string, *pb_gen.DownloadMetadate, *AErr.APIErr) {

	stream, err := rpc_cli.MSLocalCli.Download(ctx, req)
	if err != nil {
		log.Printf("failed to create download stream, err=[%v]", err)
		return "", nil, AErr.InternalErr.CustomMessageF("failed to create download stream, err=[%v]", err)
	}
	rec, err := stream.Recv() //receive metadata
	if err != nil {
		log.Printf("receive metadata failed, err=[%v]", err)
		return "", nil, AErr.InternalErr.CustomMessageF("receive metadata failed, err=[%v]", err)
	}
	metadata := rec.GetMetadata()
	fpath, err := dl.receiveStream(stream, metadata)
	if err != nil {
		return "", nil, AErr.InternalErr.CustomMessageF("receiver stream err, err=[%v]", err)
	}
	log.Printf("metadata = [%v]", *metadata)
	log.Printf("receive file path is %s", fpath)
	return fpath, metadata, nil
}

func (dl Downloader) receiveStream(stream pb_gen.MSLocal_DownloadClient, metadata *pb_gen.DownloadMetadate) (string, error) {
	fo, err := os.CreateTemp("", "temp_download_")
	if err != nil {
		log.Printf("create temp file fail, err=[%v]", err)
		return "", err
	}
	defer func() {
		if err := fo.Close(); err != nil {
			log.Printf("Upload close fo failed, err=[%+v]", err)
		}
	}()
	dataSize := 0
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			log.Printf("reseive done, file size is %dkb", dataSize/1024)
			break
		}
		if err != nil {
			log.Printf("receive chunk failed, error=[%v]", err)
			return "", err
		}
		chunk := req.GetContent()
		dataSize += len(chunk)
		_, err = fo.Write(chunk)
		if err != nil {
			log.Printf("cannot write chunk data: err=[%v]", err)
			return "", err
		}
	}
	return fo.Name(), nil //name is path
}
