package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"lj-chain-center/common/errno"
	"lj-chain-center/common/log"
	"lj-chain-center/model"
	"lj-chain-center/service"
)

//数据上链
func Luban_DepositData(c *gin.Context) {
	req := &model.LubanDepositModelReq{}
	if err := c.ShouldBindBodyWith(req, binding.JSON); err != nil {
		log.Errorf(err, "Luban_DepositData c.ShouldBind error")
		SendResponse(c, errno.ParamError, nil)
		return
	}

	service := service.NewLubanService()
	resp, err := service.DepositData(req)
	if err != nil {
		SendResponse(c, errno.HandleError, nil)
		return
	}

	SendResponse(c, nil, resp)
}

//查询交易
func Luban_QueryTx(c *gin.Context) {
	txHash := c.Query("txHash")
	if txHash == "" {
		log.Errorf(nil, "Luban_QueryTx txHash is empty")
		SendResponse(c, errno.ParamError, nil)
		return
	}

	service := service.NewLubanService()
	resp, err := service.QueryTx(txHash)
	if err != nil {
		SendResponse(c, errno.HandleError, nil)
		return
	}

	SendResponse(c, nil, resp)
}
