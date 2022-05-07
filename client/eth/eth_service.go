package eth

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rlp"
	"golang.org/x/crypto/sha3"
	my_common "lj-chain-center/client/common"
	eth_contract "lj-chain-center/client/eth/contract"
	"lj-chain-center/common/errno"
	"lj-chain-center/common/log"
	"lj-chain-center/common/util"
	"math/big"
	"strings"
)

type EthService struct {
	client *ethclient.Client
}

func NewEthService(client *ethclient.Client) *EthService {
	return &EthService{client: client}
}

//验证地址是否为合约地址
func (service *EthService) ValidateContractAddress(address string) (bool, *errno.Errno) {

	byteCode, err := service.client.CodeAt(context.Background(), common.HexToAddress(address), nil)
	if err != nil {
		log.Errorf(err, "client codeAt error")
		return false, errno.HandleError
	}
	isContractAddress := len(byteCode) > 0
	log.Infof("is contract address: %v", isContractAddress)
	return isContractAddress, nil
}

//查询eth余额
func (service *EthService) QueryEthBalance(account string) (*big.Int, *errno.Errno) {
	balance, err := service.client.BalanceAt(context.Background(), common.HexToAddress(account), nil)
	//待处理的余额
	//balance, err := service.client.PendingBalanceAt(context.Background(), common.HexToAddress(account))
	if err != nil {
		log.Errorf(err, "query eth balance error")
		return nil, errno.HandleError
	}
	return balance, nil
}

//查询区块
func (service *EthService) QueryBlock() {
	header, err := service.client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Errorf(err, "query header by number error")
		return
	}
	log.Infof("header number: %", header.Number.String())

	block, err := service.client.BlockByNumber(context.Background(), header.Number)
	if err != nil {
		log.Errorf(err, "query block by number error")
		return
	}

	log.Infof("block number: %v", block.Number().Uint64())
	log.Infof("block time: %v", block.Time())
	log.Infof("block difficulty: %v", block.Difficulty().Uint64())
	log.Infof("block hash: %v", block.Hash().Hex())
	log.Infof("block transactions: %v", len(block.Transactions()))
}

//查询交易
func (service *EthService) QueryTransaction() *errno.Errno {
	blockNumber := big.NewInt(27136004)
	block, err := service.client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Errorf(err, "queryTransaction error")
		return errno.HandleError
	}
	for index, tx := range block.Transactions() {
		log.Infof("index:%v, hash:%v", index, tx.Hash().Hex())
		log.Infof("index:%v, gas:%v, gasPrice:%v, nonce:%v", index, tx.Gas(), tx.GasPrice().Uint64(), tx.Nonce())
		log.Infof("index:%v, data:%v, toAddress:%v", index, tx.Data(), tx.To().Hex())

		/*chainID, err := service.client.NetworkID(context.Background())
		if err != nil {
			log.Errorf(err, "networkID error")
			continue
		}*/
		/*chainID := big.NewInt(42)
		log.Infof("index:%v, chainID:%v", index, chainID)
		msg, err := tx.AsMessage(types.NewEIP155Signer(chainID), nil)
		if err != nil {
			log.Errorf(err, "asMessage error")
		}
		log.Infof("index:%v, msg:%v", index, msg.From().Hex())*/

		receipt, err := service.client.TransactionReceipt(context.Background(), tx.Hash())
		if err != nil {
			log.Errorf(err, "transactionReceipt error")
		}
		log.Infof("index:%v, receipt status:%v, receipt logs:%v", index, receipt.Status, receipt.Logs)
		for index, item := range receipt.Logs {
			//如何解析日志?
			log.Infof("log index:%v, log data:%v", index, item.Data)
		}

	}
	return nil
}

//查询交易(根据交易hash)
func (service *EthService) QueryTransactionReceipt(hexHash string) *errno.Errno {
	receipt, err := service.client.TransactionReceipt(context.Background(), common.HexToHash(hexHash))
	if err != nil {
		log.Errorf(err, "transactionReceipt error")
		return errno.HandleError
	}
	log.Infof("receipt status:%v, receipt logs: %v", receipt.Status, receipt.Logs)
	for index, item := range receipt.Logs[0].Topics {
		log.Infof("index: %v, item:%v", index, common.HexToAddress(item.Hex()))
	}
	log.Infof("data value: %v", util.BytesToInt(receipt.Logs[0].Data))
	return nil
}

//转账eth
func (service *EthService) TransferETH(hexKey string, hexToAddress string) *errno.Errno {
	toAddress := common.HexToAddress(hexToAddress)
	privateKey, err := crypto.HexToECDSA(hexKey)
	if err != nil {
		log.Errorf(err, "crypto hexToECDSA error")
		return errno.HandleError
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Errorf(nil, "publicKey to ECDSA error")
		return errno.HandleError
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	log.Infof("fromAddress: %v", fromAddress)
	nonce, err := service.client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Errorf(err, "pendingNonceAt error")
		return errno.HandleError
	}
	log.Infof("nonce: %v", nonce)

	value := big.NewInt(1 * 1e16)
	gasPrice, err := service.client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Errorf(err, "suggestGasPrice error")
		return errno.HandleError
	}
	log.Infof("gasPrice: %v", gasPrice.String())
	gasLimit := uint64(1000000)

	data := []byte("我是谁 我从哪里来 我该往何处去")
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(big.NewInt(my_common.KovanChainID)), privateKey)
	if err != nil {
		log.Errorf(err, "signTx error")
		return errno.HandleError
	}
	err = service.client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Errorf(err, "sendTransaction error")
		return errno.HandleError
	}

	log.Infof("sendETH tx:%v", signedTx.Hash().Hex())
	return nil
}

//转账token(erc-20)
func (service *EthService) TransferToken(hexKey string, hexTokenAddress string, hexToAddress string) *errno.Errno {
	tokenAddress := common.HexToAddress(hexTokenAddress)
	toAddress := common.HexToAddress(hexToAddress)

	privateKey, err := crypto.HexToECDSA(hexKey)
	if err != nil {
		log.Errorf(err, "crypto hexToECDSA error")
		return errno.HandleError
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Errorf(nil, "publicKey to ECDSA error")
		return errno.HandleError
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	log.Infof("fromAddress: %v", fromAddress)

	nonce, err := service.client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Errorf(err, "pendingNonceAt error")
		return errno.HandleError
	}
	log.Infof("nonce: %v", nonce)

	value := big.NewInt(0)

	gasPrice, err := service.client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Errorf(err, "suggestGasPrice error")
		return errno.HandleError
	}
	log.Infof("gasPrice: %v", gasPrice.String())

	transferSignature := []byte("transfer(address,uint256)")
	hash := sha3.NewLegacyKeccak256()
	hash.Write(transferSignature)
	methodID := hash.Sum(nil)[:4]
	log.Infof("methodID: %v", hexutil.Encode(methodID))

	paddedToAddress := common.LeftPadBytes(toAddress.Bytes(), 32)
	log.Infof("paddedToAddress:%v", hexutil.Encode(paddedToAddress))

	amount := new(big.Int)
	amount.SetString("10000000000000000", 10) //0.1 token
	paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)
	log.Infof("paddedAmount: %v", hexutil.Encode(paddedAmount))

	var data []byte
	data = append(data, methodID...)
	data = append(data, paddedToAddress...)
	data = append(data, paddedAmount...)

	gasLimit, err := service.client.EstimateGas(context.Background(), ethereum.CallMsg{
		To:   &toAddress,
		Data: data,
	})
	if err != nil {
		log.Errorf(err, "estimateGas error")
		return errno.HandleError
	}
	log.Infof("gasLimit01: %v", gasLimit)

	gasLimit += 1000000
	log.Infof("gasLimit02: %v", gasLimit)

	tx := types.NewTransaction(nonce, tokenAddress, value, gasLimit, gasPrice, data)

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(big.NewInt(my_common.KovanChainID)), privateKey)
	if err != nil {
		log.Errorf(err, "signTx error")
		return errno.HandleError
	}
	err = service.client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Errorf(err, "sendTransaction error")
		return errno.HandleError
	}

	log.Infof("sendToken tx:%v", signedTx.Hash().Hex())
	return nil
}

//监听订阅新区块 不好使?
func (service *EthService) BlockSubscribe() *errno.Errno {
	headers := make(chan *types.Header)
	sub, err := service.client.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		log.Errorf(err, "subscribeNewHead error")
		return errno.HandleError
	}
	for {
		select {
		case err := <-sub.Err():
			log.Errorf(err, "sub error")
		case header := <-headers:
			log.Infof("header hash: %v", header.Hash().Hex())
			block, err := service.client.BlockByHash(context.Background(), header.Hash())
			if err != nil {
				log.Errorf(err, "query blockByHash error")
			} else {
				log.Infof("block hash: %v", block.Hash().Hex())
				log.Infof("block number: %v", block.Number().Uint64())
				log.Infof("block time: %v", block.Time())
				log.Infof("block nonce: %v", block.Nonce())
				log.Infof("block transactions: %v", block.Transactions())
			}
		}
	}
}

//创建裸交易
func (service *EthService) CreateRawTransaction(hexKey string, hexToAddress string) (string, *errno.Errno) {
	toAddress := common.HexToAddress(hexToAddress)
	privateKey, err := crypto.HexToECDSA(hexKey)
	if err != nil {
		log.Errorf(err, "crypto hexToECDSA error")
		return "", errno.HandleError
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Errorf(nil, "publicKey to ECDSA error")
		return "", errno.HandleError
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	log.Infof("fromAddress: %v", fromAddress)
	nonce, err := service.client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Errorf(err, "pendingNonceAt error")
		return "", errno.HandleError
	}
	log.Infof("nonce: %v", nonce)

	value := big.NewInt(1 * 1e16)
	gasPrice, err := service.client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Errorf(err, "suggestGasPrice error")
		return "", errno.HandleError
	}
	log.Infof("gasPrice: %v", gasPrice.String())
	gasLimit := uint64(1000000)

	data := []byte("我是谁？凯爹！")
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(big.NewInt(my_common.KovanChainID)), privateKey)
	if err != nil {
		log.Errorf(err, "signTx error")
		return "", errno.HandleError
	}

	//ts := types.Transactions{signedTx}
	rawTxBytes, err := signedTx.MarshalBinary()
	if err != nil {
		log.Errorf(err, " signedTx marshalBinary error")
		return "", errno.HandleError
	}
	rawTxHex := hex.EncodeToString(rawTxBytes)
	log.Infof("rawTxHex: %v", rawTxHex)
	return rawTxHex, nil
}

//发送裸交易
func (service *EthService) SendRawTransaction(hexRawTx string) *errno.Errno {
	rawTxBytes, err := hex.DecodeString(hexRawTx)
	if err != nil {
		log.Errorf(err, "hex decodeString error")
		return errno.HandleError
	}
	tx := new(types.Transaction)
	rlp.DecodeBytes(rawTxBytes, &tx)

	err = service.client.SendTransaction(context.Background(), tx)
	if err != nil {
		log.Errorf(err, "sendTransaction error")
		return errno.HandleError
	}

	log.Infof("send tx: %v", tx.Hash().Hex())
	return nil
}

//部署demo智能合约
func (service *EthService) DeployDemoContract(hexKey string) {
	privateKey, err := crypto.HexToECDSA(hexKey)
	if err != nil {
		log.Errorf(err, "crypto hexToECDSA error")
		return
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Errorf(nil, "publicKey to ECDSA error")
		return
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	log.Infof("fromAddress: %v", fromAddress)

	nonce, err := service.client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Errorf(err, "pendingNonceAt error")
		return
	}
	log.Infof("nonce: %v", nonce)

	gasPrice, err := service.client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Errorf(err, "suggestGasPrice error")
		return
	}
	log.Infof("gasPrice: %v", gasPrice.String())

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(my_common.KovanChainID))
	if err != nil {
		log.Errorf(err, "newKeyedTransactorWithChainID error")
		return
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(1000000) // in units
	auth.GasPrice = gasPrice

	address, tx, instance, err := eth_contract.DeployEthDemo(auth, service.client, "我是谁？凯爹")
	if err != nil {
		log.Errorf(err, "deployStore error")
		return
	}
	log.Infof("contract address: %v", address.Hex())
	log.Infof("tx hash: %v", tx.Hash().Hex())
	_ = instance
}

//加载demo智能合约
func (service *EthService) LodeDemoContract(hexContractAddress string) (*eth_contract.EthDemo, *errno.Errno) {
	contractAddress := common.HexToAddress(hexContractAddress)
	instance, err := eth_contract.NewEthDemo(contractAddress, service.client)
	if err != nil {
		log.Errorf(err, "contract new file error")
		return nil, errno.HandleError
	}
	log.Infof("load contract :%v success", hexContractAddress)
	return instance, nil
}

//demo智能合约-getValue-读操作
func (service *EthService) DemoContractGetValue(hexContractAddress string) *errno.Errno {
	instance, error := service.LodeDemoContract(hexContractAddress)
	if error != nil {
		log.Errorf(error, "load contract error")
		return errno.HandleError
	}

	value, err := instance.GetValue(nil)
	if err != nil {
		log.Errorf(err, "getValue error")
		return errno.HandleError
	}
	log.Infof("value: %v", value)
	return nil
}

//demo智能合约-setValue-写操作
func (service *EthService) DemoContractSetValue(hexContractAddress string, hexKey string, data string) *errno.Errno {
	privateKey, err := crypto.HexToECDSA(hexKey)
	if err != nil {
		log.Errorf(err, "crypto hexToECDSA error")
		return errno.HandleError
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Errorf(nil, "publicKey to ECDSA error")
		return errno.HandleError
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	log.Infof("fromAddress: %v", fromAddress)

	nonce, err := service.client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Errorf(err, "pendingNonceAt error")
		return errno.HandleError
	}
	log.Infof("nonce: %v", nonce)

	gasPrice, err := service.client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Errorf(err, "suggestGasPrice error")
		return errno.HandleError
	}
	log.Infof("gasPrice: %v", gasPrice.String())

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(my_common.KovanChainID))
	if err != nil {
		log.Errorf(err, "newKeyedTransactorWithChainID error")
		return errno.HandleError
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(1000000) // in units
	auth.GasPrice = gasPrice

	instance, error := service.LodeDemoContract(hexContractAddress)
	if error != nil {
		log.Errorf(error, "load contract error")
		return errno.HandleError
	}
	tx, err := instance.SetValue(auth, data)
	if err != nil {
		log.Errorf(err, "setValue error")
		return errno.HandleError
	}

	log.Infof("setValue tx: %v", tx.Hash().Hex())

	return nil
}

//查询智能合约的字节码
func (service *EthService) QueryContractByteCode(hexContractAddress string) (string, *errno.Errno) {
	contactAddress := common.HexToAddress(hexContractAddress)
	byteCode, err := service.client.CodeAt(context.Background(), contactAddress, nil)
	if err != nil {
		log.Errorf(err, "codeAt error")
		return "", errno.HandleError
	}
	hexByteCode := hex.EncodeToString(byteCode)
	log.Infof("hexByteCode: %v", hexByteCode)
	return hexByteCode, nil
}

//ERC20代币智能合约-读操作
func (service *EthService) ERC20ContractRead(hexTokenAddress string, hexOwnerAddress string) *errno.Errno {
	tokenAddress := common.HexToAddress(hexTokenAddress)
	instance, err := eth_contract.NewEthErc20(tokenAddress, service.client)
	if err != nil {
		log.Errorf(err, "new erc20 error")
		return errno.HandleError
	}

	name, err := instance.Name(nil)
	if err != nil {
		log.Errorf(err, "name error")
		return errno.HandleError
	}
	log.Infof("name: %v", name)

	symbol, err := instance.Symbol(nil)
	if err != nil {
		log.Errorf(err, "symbol error")
		return errno.HandleError
	}
	log.Infof("symbol: %v", symbol)

	decimals, err := instance.Decimals(nil)
	if err != nil {
		log.Errorf(err, "decimals error")
		return errno.HandleError
	}
	log.Infof("decimals: %v", decimals)

	ownerAddress := common.HexToAddress(hexOwnerAddress)
	balance, err := instance.BalanceOf(nil, ownerAddress)
	if err != nil {
		log.Errorf(err, "balanceOf error")
		return errno.HandleError
	}
	log.Infof("ownerAddress: %v, balance: %v", hexOwnerAddress, balance.String())

	return nil
}

//ERC20代币智能合约-订阅日志事件
func (service *EthService) ERC20ContractSubscribe(hexContactAddress string) *errno.Errno {
	contractAddress := common.HexToAddress(hexContactAddress)
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}
	logs := make(chan types.Log)
	sub, err := service.client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Errorf(err, "subscribe filter logs error")
		return errno.HandleError
	}

	log.Infof("abi: %v", eth_contract.EthErc20ABI)
	contractAbi, err := abi.JSON(strings.NewReader(eth_contract.EthErc20ABI))
	if err != nil {
		log.Errorf(err, "abi json error")
		return errno.HandleError
	}

	logTransferSig := []byte("Transfer(address,address,uint256)")
	logTransferSigHash := crypto.Keccak256Hash(logTransferSig)
	log.Infof("logTransferSigHash: %v", logTransferSigHash.Hex())

	type TransferEvent struct {
		From   common.Address
		To     common.Address
		Amount *big.Int
	}

	for {
		select {
		case err := <-sub.Err():
			log.Errorf(err, "sub error")
		case vLog := <-logs:
			log.Infof("vlog all: #%v", vLog)
			log.Infof("vlog contract address hex: %v", vLog.Address.Hex())

			log.Infof("block number: %v", vLog.BlockNumber)
			log.Infof("block hash: %v", vLog.BlockHash.Hex())
			log.Infof("tx hash: %v", vLog.TxHash.Hex())
			log.Infof("tx index: %v", vLog.TxIndex)
			log.Infof("index: %v", vLog.Index)
			log.Infof("removed: %v", vLog.Removed)

			log.Infof("topics length: %v", len(vLog.Topics))
			//event Transfer(address indexed from, address indexed to, uint256 amount)
			//前两个参数声明为indexed，被视为主题，最后一个参数没有indexed，被作为数据
			switch vLog.Topics[0].Hex() {
			case logTransferSigHash.Hex():
				for index, item := range vLog.Topics {
					//第一个主题通常为事件名称及其参数类型
					log.Infof("topics index: %v, item: %v", index, item.Hex())
				}

				var transferEvent TransferEvent
				err := contractAbi.UnpackIntoInterface(&transferEvent, "Transfer", vLog.Data)
				if err != nil {
					log.Errorf(err, "contractAbi unpackIntoInterface transfer error")
					return errno.HandleError
				}
				transferEvent.From = common.HexToAddress(vLog.Topics[1].Hex())
				transferEvent.To = common.HexToAddress(vLog.Topics[2].Hex())

				log.Infof("fromAddress: %v", transferEvent.From.Hex())
				log.Infof("toAddress: %v", transferEvent.To.Hex())
				log.Infof("amount: %v", transferEvent.Amount.String())
			}
		}
	}
}
