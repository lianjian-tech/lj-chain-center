package service

import (
	"lj-chain-center/common/log"
	"lj-chain-center/common/util"
	"lj-chain-center/model"
	"testing"
)

func TestCoreService_CreateCompanyDID(t *testing.T) {
	initConf()
	service := NewCoreService()
	req := &model.CreateDIDModelReq{}
	req.BizID = "122"
	req.Identify = "450923122"
	resp, err := service.CreateCompanyDID(req)
	if err != nil {
		log.Errorf(err, "service.CreateCompanyDID error")
		return
	}
	log.Infof("resp:%v", util.ToJSONStr(resp))
}

func TestCoreService_CreatePersonDID(t *testing.T) {
	initConf()
	service := NewCoreService()
	req := &model.CreateDIDModelReq{}
	req.BizID = "221"
	req.Identify = "550923113"
	resp, err := service.CreatePersonDID(req)
	if err != nil {
		log.Errorf(err, "service.CreatePersonDID error")
		return
	}
	log.Infof("resp:%v", util.ToJSONStr(resp))
}

func TestCoreService_GetCompanyDIDList(t *testing.T) {
	initConf()
	service := NewCoreService()
	identity := "450923122"
	resp, err := service.GetCompanyDIDList(identity)
	if err != nil {
		log.Errorf(err, "service.GetCompanyDIDList error")
		return
	}
	log.Infof("resp:%v", util.ToJSONStr(resp))
	log.Infof("resp len:%v", len(resp))
}

func TestCoreService_GetPersonDIDList(t *testing.T) {
	initConf()
	service := NewCoreService()
	identity := "550923113"
	resp, err := service.GetPersonDIDList(identity)
	if err != nil {
		log.Errorf(err, "service.GetPersonDIDList error")
		return
	}
	log.Infof("resp:%v", util.ToJSONStr(resp))
	log.Infof("resp len:%v", len(resp))
}
