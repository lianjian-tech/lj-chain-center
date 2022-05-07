package service

import (
	"lj-chain-center/common/log"
	"lj-chain-center/common/util"
	"testing"
)

func TestAntBassService_QueryTx(t *testing.T) {
	initConf()
	service := NewAntBassService()
	txHash := "10658f7ffcb61c857a14a4719c60fb9c497d892e58bead94a5b7af774d36d1ee"
	data, err := service.QueryTx(txHash)
	if err != nil {
		return
	}
	log.Infof("data:%v", util.ToJSONStr(data))
}
