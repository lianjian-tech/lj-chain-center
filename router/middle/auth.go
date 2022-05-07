package middle

import (
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"lj-chain-center/common/errno"
	"lj-chain-center/common/log"
	"lj-chain-center/dao"
	"lj-chain-center/handler"
	"lj-chain-center/model"
	"lj-chain-center/pkg"
)

func Auth(validateDID bool, bizType int) gin.HandlerFunc {
	return func(c *gin.Context) {
		if errno := verifySign(c, validateDID, bizType); errno != nil {
			log.Errorf(errno, "Auth verifySign error")
			handler.SendResponse(c, errno, nil)
			c.Abort()
			return
		}
		c.Next()
	}
}

func verifySign(c *gin.Context, validateDID bool, bizType int) *errno.Errno {
	req := &model.BaseAuthModelReq{}
	if err := c.ShouldBindBodyWith(req, binding.JSON); err != nil {
		log.Errorf(err, "verifySign c.ShouldBind error")
		return errno.ParamError
	}

	if validateDID {
		if req.DID == "" {
			log.Errorf(nil, "verifySign req.DID empty")
			return errno.AuthError
		}
		match := verifyDID(req.DID, bizType)
		if !match {
			log.Errorf(nil, "verifySign verifyDID false")
			return errno.AuthError
		}
	}

	return nil
}

func sign(hexPrivateKey string, data string) (string, string, error) {
	privateKey, err := crypto.HexToECDSA(hexPrivateKey)
	if err != nil {
		log.Errorf(err, "sign crypto.HexToECDSA error")
		return "", "", errors.New(pkg.HANDLE_ERROR)
	}
	dataByte := []byte(data)
	dataHash := crypto.Keccak256Hash(dataByte)
	hexDataHash := dataHash.Hex()

	dataSign, err := crypto.Sign(dataHash.Bytes(), privateKey)
	if err != nil {
		log.Errorf(err, "sign crypto.Sign error")
		return "", "", errors.New(pkg.HANDLE_ERROR)
	}

	hexDataSign := hexutil.Encode(dataSign)

	return hexDataHash, hexDataSign, nil
}

func verifySignSub(hexPublicKey string, hexData string, hexSign string) bool {
	dataByte, err := hexutil.Decode(hexData)
	if err != nil {
		return false
	}
	signByte, err := hexutil.Decode(hexSign)
	if err != nil {
		return false
	}
	signPublicKey, err := crypto.Ecrecover(dataByte, signByte)
	if err != nil {
		return false
	}
	hexSignPublicKey := common.BytesToAddress(crypto.Keccak256(signPublicKey[1:])[12:])

	matched := hexSignPublicKey.Hex() == hexPublicKey

	return matched
}

func verifyDID(did string, bizType int) bool {
	companyDIDDao := dao.NewCompanyDIDDao()
	item, err := companyDIDDao.GetCompanyDIDByID(did, bizType)
	if err != nil {
		return false
	}
	if item.DID == "" {
		return false
	}

	return true
}
