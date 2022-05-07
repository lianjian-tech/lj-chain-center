package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"lj-chain-center/common/log"
	"lj-chain-center/common/util"
	"lj-chain-center/model"
	"lj-chain-center/pkg"
)

type AntBassService struct {
}

func NewAntBassService() *AntBassService {
	return &AntBassService{}
}

// 创建DID (公司和个人)
func (service *AntBassService) CreateAntBassDID(bizType int, req *model.CreateDIDModelReq) (string, error) {
	antBassCreateDIDReq := &model.AntBassCreateDIDCallReq{}
	antBassCreateDIDReq.Type = bizType
	antBassCreateDIDReq.OwnerUID = req.OwnerUID

	url := fmt.Sprintf(`%v/api/biz/createDID`, pkg.Cfg.AntBass.Url)
	log.Infof("CreateAntBassDID url:%v", url)
	log.Infof("CreateAntBassDID req:%v", util.ToJSONStr(antBassCreateDIDReq))

	result, err := util.HttpClient(url, util.POST, []byte(util.ToJSONStr(antBassCreateDIDReq)))
	if err != nil {
		log.Errorf(err, "CreateAntBassDID util.HttpClient error")
		return "", errors.New(pkg.HANDLE_ERROR)
	}
	log.Infof("CreateAntBassDID result:%v", string(result))

	callResp := &model.AntBassBaseCallResp{}
	if err := json.Unmarshal(result, callResp); err != nil {
		log.Errorf(err, "CreateAntBassDID json.Unmarshal error")
		return "", errors.New(pkg.HANDLE_ERROR)
	}
	if !callResp.Success {
		log.Errorf(nil, "CreateAntBassDID callResp error")
		return "", errors.New(pkg.HANDLE_ERROR)
	}
	log.Infof("CreateAntBassDID callResp:%v", util.ToJSONStr(callResp))

	dataMap := callResp.Data.(map[string]interface{})
	did := dataMap["did"].(string)

	log.Infof("CreateAntBassDID did:%v", did)

	return did, nil
}

// 链上交易信息查询
func (service *AntBassService) QueryTx(txHash string) (interface{}, error) {
	url := fmt.Sprintf(`%v/api/biz/txQuery`, pkg.Cfg.AntBass.Url)
	params := fmt.Sprintf(`txHash=%v`, txHash)
	url = fmt.Sprintf(`%v?%v`, url, params)
	log.Infof("QueryTx url:%v", url)
	result, err := util.HttpClient(url, util.GET, nil)
	if err != nil {
		log.Errorf(err, "QueryTx util.HttpClient error")
		return nil, errors.New(pkg.HANDLE_ERROR)
	}
	log.Infof("QueryTx result:%v", string(result))

	callResp := &model.AntBassBaseCallResp{}
	if err := json.Unmarshal(result, callResp); err != nil {
		log.Errorf(err, "QueryTx json.Unmarshal error")
		return nil, errors.New(pkg.HANDLE_ERROR)
	}
	if !callResp.Success {
		log.Errorf(nil, "QueryTx callResp error")
		return "", errors.New(pkg.HANDLE_ERROR)
	}
	return callResp.Data, nil
}
