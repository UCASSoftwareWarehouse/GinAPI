package code_sim_uploader

import (
	"GinAPI/config"
	"GinAPI/rpc_cli"
	"GinAPI/util"
	"context"
	"testing"
)

func TestUploadSourceCode(t *testing.T) {
	config.InitConfigDefault()
	rpc_cli.InitRPCClient()
	err := UploadSourceCode(context.Background(), "yzchnb/GinAPI", "v1.0", "/Users/purchaser/Desktop/GinAPI.zip", util.ArchiveFileTypeZip)
	t.Logf("err=[%+v]", err)
}
