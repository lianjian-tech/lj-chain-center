package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"lj-chain-center/common/errno"
	"lj-chain-center/common/log"
	"lj-chain-center/model"
	"lj-chain-center/service"
)

//创建公司DID
func Core_CreateCompanyDID(c *gin.Context) {
	req := &model.CreateDIDModelReq{}
	if err := c.ShouldBindBodyWith(req, binding.JSON); err != nil {
		log.Errorf(err, "Core_CreateCompanyDID c.ShouldBind error")
		SendResponse(c, errno.ParamError, nil)
		return
	}
	service := service.NewCoreService()
	resp, err := service.CreateCompanyDID(req)
	if err != nil {
		SendResponse(c, errno.HandleError, nil)
		return
	}

	SendResponse(c, nil, resp)
}

//创建个人DID
func Core_CreatePersonDID(c *gin.Context) {
	req := &model.CreateDIDModelReq{}
	if err := c.ShouldBindBodyWith(req, binding.JSON); err != nil {
		log.Errorf(err, "Core_CreatePersonDID c.ShouldBind error")
		SendResponse(c, errno.ParamError, nil)
		return
	}
	service := service.NewCoreService()
	resp, err := service.CreatePersonDID(req)
	if err != nil {
		SendResponse(c, errno.HandleError, nil)
		return
	}

	SendResponse(c, nil, resp)
}

// 查询公司DID
func Core_QueryCompanyDIDList(c *gin.Context) {
	req := &model.QueryCompanyDIDListReq{}
	if err := c.ShouldBindBodyWith(req, binding.JSON); err != nil {
		log.Errorf(err, "Core_QueryCompanyDIDList c.ShouldBind error")
		SendResponse(c, errno.ParamError, nil)
		return
	}
	service := service.NewCoreService()
	resp, err := service.GetCompanyDIDList(req.Identify)
	if err != nil {
		SendResponse(c, errno.HandleError, nil)
		return
	}

	SendResponse(c, nil, resp)
}

// 查询个人DID
func Core_QueryPersonDIDList(c *gin.Context) {
	req := &model.QueryPersonDIDListReq{}
	if err := c.ShouldBindBodyWith(req, binding.JSON); err != nil {
		log.Errorf(err, "Core_QueryPersonDIDList c.ShouldBind error")
		SendResponse(c, errno.ParamError, nil)
		return
	}
	service := service.NewCoreService()
	resp, err := service.GetPersonDIDList(req.Identify)
	if err != nil {
		SendResponse(c, errno.HandleError, nil)
		return
	}

	SendResponse(c, nil, resp)
}
