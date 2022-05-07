package did

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"go-chain-center/common/errno"
	"go-chain-center/common/log"
)

func CreateKey() (string, string, *errno.Errno) {
	key, err := crypto.GenerateKey()
	if err != nil {
		log.Errorf(err, "GenerateKey error")
		return "", "", errno.HandleError
	}
	privateKeyBytes := crypto.FromECDSA(key)
	privateKey := hexutil.Encode(privateKeyBytes)[2:]
	publicKey := key.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
	}
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	return privateKey, address, nil
}
