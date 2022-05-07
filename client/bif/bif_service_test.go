package bif

import (
	"go-chain-center/pkg"
	"testing"
)

func TestBifService_GetAccessToken(t *testing.T) {
	initConf()
	bifService := NewBifService()
	bifService.GetAccessToken()
}

func initConf() {
	pkg.Init("../../conf/config-local.yaml")
}
