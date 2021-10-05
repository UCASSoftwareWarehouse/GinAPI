package search

import (
	//_ "GinAPI/docs"
	"GinAPI/err"
	"GinAPI/format"
	"GinAPI/pb_gen"
	"GinAPI/rpc_cli"
	"fmt"
	"github.com/gin-gonic/gin"
)

type Handler struct {}

var handler = &Handler{}

func GetHandler() *Handler {
	return handler
}

// Ping
// @Summary Print
// @Accept json
// @Tags Search
// @Produce  json
// @Router /search/ping [get]
// @Success 200 {object} search.PingCodeSimResponse
func (h Handler) Ping(c *gin.Context) (*format.JSONRespFormat, *err.APIErr){
	myHelloText := "HelloFromWebAPI!"
	resp, mErr := rpc_cli.CodeSimCli.HelloWorld(c, &pb_gen.CodeSimHelloWorldRequest{
		HelloText: myHelloText,
	})
	if mErr != nil {
		return nil, err.InternalErr.CustomMessage(fmt.Sprintf("Ping Code Sim Failed, err=[%s]", mErr))
	}
	return format.SimpleOKResp(&PingCodeSimResponse{
		FullText: fmt.Sprintf("MyHelloText is %s, ThanksText is %s", myHelloText, resp.GetThanksText()),
	}), nil
}