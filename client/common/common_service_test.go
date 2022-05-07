package client_common

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/shopspring/decimal"
	"go-chain-center/common/log"
	"go-chain-center/pkg"
	"math/big"
	"testing"
)

func TestNewAccount(t *testing.T) {
	initConfig()
	NewAccount()
}

func TestValidateAddress(t *testing.T) {
	initConfig()
	//address := "0xZYXb5d4c32345ced77393b3530b1eed0f346429d"
	address := "0x1F9184646884753A0947cace94F8eE33dbD7B060"
	result := ValidateAddress(address)
	log.Infof("validate address result: %v", result)
}

func TestNewClientDial(t *testing.T) {
	initConfig()
	url := KovanUrl
	_, err := NewClientDial(url)
	if err != nil {
		log.Errorf(err, "NewClientDial error")
		return
	}
	log.Info("NewClientDial success")
}

func TestSignature(t *testing.T) {
	initConfig()
	data := "大话西游，爱你一万年"
	Signature(HexKey, data)
}

func TestVerifySignature(t *testing.T) {
	initConfig()
	data := "大话西游，爱你一万年"
	hexData, hexSign, err := Signature(HexKey, data)
	if err != nil {
		return
	}
	VerifySignature(hexData, hexSign, OwnerAddress)
}

func TestIsZeroAddress(t *testing.T) {
	initConfig()
	validAddress := common.HexToAddress("0x323b5d4c32345ced77393b3530b1eed0f346429d")
	zeroAddress := common.HexToAddress("0x0000000000000000000000000000000000000000")

	log.Infof("111: %v", IsZeroAddress(validAddress))
	log.Infof("222: %v", IsZeroAddress(zeroAddress))
}

func TestToWei(t *testing.T) {
	initConfig()
	amount := decimal.NewFromFloat(0.02)
	result := ToWei(amount, 18)
	log.Infof("result: %v", result.String())
}

func TestToDecimal(t *testing.T) {
	initConfig()
	weiAmount := big.NewInt(0)
	weiAmount.SetString("20000000000000000", 10)
	log.Infof("weiAmount: %v", weiAmount.String())
	ethAmount := ToDecimal(weiAmount, 18)
	f64, _ := ethAmount.Float64()
	log.Infof("result: %v", f64)
}

func TestCalcGasCost(t *testing.T) {
	initConfig()
	gasPrice := big.NewInt(0)
	gasPrice.SetString("2000000000", 10)
	gasLimit := uint64(21000)
	result := CalcGasCost(gasLimit, gasPrice)
	log.Infof("result: %v", result.String())
}

func TestSigRSV(t *testing.T) {
	initConfig()
	sig := "0x789a80053e4927d0a898db8e065e948f5cf086e32f9ccaa54c1908e22ac430c62621578113ddbb62d509bf6049b8fb544ab06d36f916685a2eb8e57ffadde02301"
	r, s, v := SigRSV(sig)
	log.Infof("r: %v", hexutil.Encode(r[:])[2:])
	log.Infof("s: %v", hexutil.Encode(s[:])[2:])
	log.Infof("v: %v", v)
}

func initConfig() {
	pkg.Init("../../conf/config-dev.yaml")
}
