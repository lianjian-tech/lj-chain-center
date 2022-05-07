package handler

import (
	"github.com/gin-gonic/gin"
	"lj-chain-center/common/errno"
	"lj-chain-center/common/log"
	"lj-chain-center/service"
)

//8.链上交易信息查询
func AntBass_QueryTx(c *gin.Context) {
	txHash := c.Query("txHash")
	if txHash == "" {
		log.Errorf(nil, "AntBass_QueryTx txHash is empty")
		SendResponse(c, errno.ParamError, nil)
		return
	}

	service := service.NewAntBassService()
	resp, err := service.QueryTx(txHash)
	if err != nil {
		SendResponse(c, errno.HandleError, nil)
		return
	}

	SendResponse(c, nil, resp)
}
