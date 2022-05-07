package bif

import (
	"errors"
	"go-chain-center/common/log"
	"go-chain-center/common/util"
	"go-chain-center/pkg"
)

type BifService struct {
}

func NewBifService() *BifService {
	return &BifService{}
}

func (service *BifService) GetAccessToken() (string, error) {
	url := pkg.Cfg.Bif.BaseUrl + "/auth/getAccessToken"
	log.Infof("getAccessToken url: %v", url)
	params := map[string]string{"api_key": pkg.Cfg.Bif.ApiKey, "api_secret": pkg.Cfg.Bif.ApiSecret}
	log.Infof("getAccessToken param: %v", params)
	result, err := util.HttpClient(url, util.GET, []byte(util.ToJSONStr(params)))
	if err != nil {
		log.Errorf(err, "GetAccessToken util.HttpClient error")
		return "", errors.New(pkg.HANDLE_ERROR)
	}
	log.Infof("GetAccessToken result:%v", string(result))
	return "", nil
}
