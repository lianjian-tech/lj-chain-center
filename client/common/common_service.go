package client_common

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/shopspring/decimal"
	"lj-chain-center/common/errno"
	"lj-chain-center/common/log"
	"math/big"
	"reflect"
	"regexp"
	"strconv"
)

const (
	OwnerAddress = "0x833dE082e21E1250fc112F2654E2441052ca28fB"
	HexKey       = "0ef60a715daaec6125c8ce0f5b6f4f4fa4a2fc5ea01b038c8710f5a9382b046a"
	KovanUrl     = "https://kovan.infura.io/v3/d5e0a32f3ffe4945907d5cb990190352"
	KovanWss     = "wss://kovan.infura.io/ws/v3/d5e0a32f3ffe4945907d5cb990190352"
	KovanChainID = int64(42)
	IpfsLocalUrl = "localhost:5001"
)

//生成账号(私钥和公钥)
func NewAccount() *errno.Errno {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Errorf(err, "crypto generateKey error")
		return errno.HandleError
	}

	privateKeyBytes := crypto.FromECDSA(privateKey)
	log.Infof("privateKey: %v", hexutil.Encode(privateKeyBytes))

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Errorf(nil, "to publicKey error")
		return errno.HandleError
	}
	//publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	//log.Infof("publicKey: %v", hexutil.Encode((publicKeyBytes)[4:]))

	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	log.Infof("address: %v", address)

	return nil
}

//验证是否为地址
func ValidateAddress(address string) bool {
	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
	flag := re.MatchString(address)
	log.Infof("is valid: %v", flag)
	return flag
}

//实例化客户端
func NewClientDial(url string) (*ethclient.Client, *errno.Errno) {
	client, err := ethclient.Dial(url)
	if err != nil {
		log.Errorf(err, "eth client connection error")
		return nil, errno.HandleError
	}
	log.Info("eth client connection success")
	return client, nil
}

//私钥签名
func Signature(hexKey string, value string) (string, string, *errno.Errno) {
	privateKey, err := crypto.HexToECDSA(hexKey)
	if err != nil {
		log.Errorf(err, "crypto hexToECDSA error")
		return "", "", errno.HandleError
	}

	data := []byte(value)
	hash := crypto.Keccak256Hash(data)
	log.Infof("hash: %v", hash.Hex())
	signature, err := crypto.Sign(hash.Bytes(), privateKey)
	if err != nil {
		log.Errorf(err, "crypto sign error")
		return "", "", errno.HandleError
	}

	hexSignature := hexutil.Encode(signature)
	log.Infof("hexSignature: %v", hexSignature)

	return hash.Hex(), hexSignature, nil
}

//验证签名
func VerifySignature(hexData string, hexSign string, hexOwnerAddress string) (bool, *errno.Errno) {
	signByte, err := hexutil.Decode(hexSign)
	if err != nil {
		log.Errorf(err, "hexSign decode error")
		return false, errno.HandleError
	}
	dataByte, err := hexutil.Decode(hexData)
	if err != nil {
		log.Errorf(err, "hexData decode error")
		return false, errno.HandleError
	}
	signPublicKey, err := crypto.Ecrecover(dataByte, signByte)
	if err != nil {
		log.Errorf(err, "crypto ecrecover error")
		return false, errno.HandleError
	}
	hexSignPublicKey := common.BytesToAddress(crypto.Keccak256(signPublicKey[1:])[12:])
	log.Infof("hexSignPublicKey: %v", hexSignPublicKey.Hex())

	matched := hexSignPublicKey.Hex() == hexOwnerAddress
	log.Infof("matched:%v", matched)
	return matched, nil
}

// IsZeroAddress validate if it's a 0 address
func IsZeroAddress(validateAddress interface{}) bool {
	var address common.Address
	switch v := validateAddress.(type) {
	case string:
		address = common.HexToAddress(v)
	case common.Address:
		address = v
	default:
		return false
	}

	zeroAddressBytes := common.FromHex("0x0000000000000000000000000000000000000000")
	addressBytes := address.Bytes()
	return reflect.DeepEqual(addressBytes, zeroAddressBytes)
}

// ToWei decimals to wei
func ToWei(targetAmount interface{}, decimals int) *big.Int {
	amount := decimal.NewFromFloat(0)
	switch v := targetAmount.(type) {
	case string:
		amount, _ = decimal.NewFromString(v)
	case float64:
		amount = decimal.NewFromFloat(v)
	case int64:
		amount = decimal.NewFromFloat(float64(v))
	case decimal.Decimal:
		amount = v
	case *decimal.Decimal:
		amount = *v
	}

	mul := decimal.NewFromFloat(float64(10)).Pow(decimal.NewFromFloat(float64(decimals)))
	result := amount.Mul(mul)

	wei := new(big.Int)
	wei.SetString(result.String(), 10)

	return wei
}

// ToDecimal wei to decimals
func ToDecimal(targetValue interface{}, decimals int) decimal.Decimal {
	value := new(big.Int)
	switch v := targetValue.(type) {
	case string:
		value.SetString(v, 10)
	case *big.Int:
		value = v
	}

	mul := decimal.NewFromFloat(float64(10)).Pow(decimal.NewFromFloat(float64(decimals)))
	num, _ := decimal.NewFromString(value.String())
	result := num.Div(mul)

	return result
}

// CalcGasCost calculate gas cost given gas limit (units) and gas price (wei)
func CalcGasCost(gasLimit uint64, gasPrice *big.Int) *big.Int {
	gasLimitBig := big.NewInt(int64(gasLimit))
	return gasLimitBig.Mul(gasLimitBig, gasPrice)
}

// SigRSV signatures R S V returned as arrays
// r,s,v是交易签名后的值，它们可以被用来生成签名者的公钥；R，S是ECDSA椭圆加密算法的输出值，V是用于恢复结果的ID
func SigRSV(targetSig interface{}) ([32]byte, [32]byte, uint8) {
	var sig []byte
	switch v := targetSig.(type) {
	case []byte:
		sig = v
	case string:
		sig, _ = hexutil.Decode(v)
	}

	sigStr := common.Bytes2Hex(sig)
	rS := sigStr[0:64]
	sS := sigStr[64:128]
	R := [32]byte{}
	S := [32]byte{}
	copy(R[:], common.FromHex(rS))
	copy(S[:], common.FromHex(sS))
	vStr := sigStr[128:130]
	vI, _ := strconv.Atoi(vStr)
	V := uint8(vI + 27)

	return R, S, V
}
