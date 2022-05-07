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

type LubanService struct {
}

func NewLubanService() *LubanService {
	return &LubanService{}
}

//创建luban did
func (service *LubanService) CreateLubanDID(req *model.CreateDIDModelReq) (string, error) {
	LubanCreateDIDReq := &model.LubanCreateDIDCallReq{}
	LubanCreateDIDReq.OwnerUID = req.OwnerUID
	url := fmt.Sprintf(`%v/api/biz/createDID`, pkg.Cfg.Luban.Url)
	log.Infof("url:%v", url)
	log.Infof("CreateLubanDID LubanCreateDIDCallReq:%v", util.ToJSONStr(LubanCreateDIDReq))
	result, err := util.HttpClient(url, util.POST, []byte(util.ToJSONStr(LubanCreateDIDReq)))
	if err != nil {
		log.Errorf(err, "CreateLubanDID util.HttpClient error")
		return "", errors.New(pkg.HANDLE_ERROR)
	}
	log.Infof("CreateAntBassDID result:%v", string(result))

	callResp := &model.LubanBaseCallResp{}
	if err := json.Unmarshal(result, callResp); err != nil {
		log.Errorf(err, "CreateLubanDID json.Unmarshal error")
		return "", errors.New(pkg.HANDLE_ERROR)
	}
	if !callResp.Success {
		log.Errorf(nil, "CreateLubanDID callResp error:%v", util.ToJSONStr(callResp))
		return "", errors.New(pkg.HANDLE_ERROR)
	}
	log.Infof("CreateLubanDID callResp error")

	dataMap := callResp.Data.(map[string]interface{})
	did := dataMap["did"].(string)

	log.Infof("CreateLubanDID did:%v", did)

	return did, nil
}

//查询交易(可信工程平台)
func (service *LubanService) QueryTx(txHash string) (*model.LubanQueryTxData, error) {
	url := fmt.Sprintf(`%v/api/biz/txQuery`, pkg.Cfg.Luban.Url)
	params := fmt.Sprintf(`txHash=%v`, txHash)
	url = fmt.Sprintf(`%v?%v`, url, params)
	log.Infof("QueryTx url:%v", url)

	result, err := util.HttpClient(url, util.GET, nil)
	if err != nil {
		log.Errorf(err, "QueryTx util.HttpClient error")
		return nil, errors.New(pkg.HANDLE_ERROR)
	}
	log.Infof("QueryTx result:%v", string(result))

	callResp := &model.LubanBaseCallResp{}
	if err := json.Unmarshal(result, callResp); err != nil {
		log.Errorf(err, "QueryTx json.Unmarshal error")
		return nil, errors.New(pkg.HANDLE_ERROR)
	}
	if !callResp.Success {
		log.Errorf(nil, "QueryTx callResp error")
		return nil, errors.New(pkg.HANDLE_ERROR)
	}

	dataMap := callResp.Data.(map[string]interface{})
	data := &model.LubanQueryTxData{}
	if err := json.Unmarshal([]byte(util.ToJSONStr(dataMap)), data); err != nil {
		log.Errorf(err, "QueryTx json.Unmarshal error")
		return nil, errors.New(pkg.HANDLE_ERROR)
	}

	return data, nil
}

//存证上链(可信工程平台)
func (service *LubanService) DepositData(req *model.LubanDepositModelReq) (string, error) {
	url := fmt.Sprintf(`%v/api/biz/depositData`, pkg.Cfg.Luban.Url)
	log.Infof("DepositData req:%v", util.ToJSONStr(req))

	result, err := util.HttpClient(url, util.POST, []byte(util.ToJSONStr(req)))
	if err != nil {
		log.Errorf(err, "DepositData util.HttpClient error")
		return "", errors.New(pkg.HANDLE_ERROR)
	}
	log.Infof("DepositData result:%v", string(result))

	callResp := &model.LubanBaseCallResp{}
	if err := json.Unmarshal(result, callResp); err != nil {
		log.Errorf(err, "DepositData json.Unmarshal error")
		return "", errors.New(pkg.HANDLE_ERROR)
	}
	if !callResp.Success {
		log.Errorf(nil, "DepositData callResp error:%v")
		return "", errors.New(pkg.HANDLE_ERROR)
	}
	data := callResp.Data.(string)
	log.Infof("DepositData data:%v", data)
	return data, nil
}
