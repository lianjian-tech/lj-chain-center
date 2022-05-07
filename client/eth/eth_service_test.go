package eth

import (
	"fmt"
	my_common "lj-chain-center/client/common"
	"lj-chain-center/common/errno"
	"lj-chain-center/common/log"
	"lj-chain-center/common/util"
	"math"
	"testing"
)

//验证地址是否为合约地址
func TestEthService_ValidateContractAddress(t *testing.T) {
	service, err := newEthService4Test(my_common.KovanUrl)
	if err != nil {
		return
	}
	//account := "0xFF33C20B116F77A8F5AccDb372ACDC00a75ceb4a"
	account := "0x1f9840a85d5aF5bf1D1762F925BDADdC4201F984"
	result, err := service.ValidateContractAddress(account)
	if err != nil {
		log.Errorf(err, "ValidateContractAddress error")
		return
	}
	log.Infof("ValidateContractAddress result: %v", result)
}

//查询eth余额
func TestEthService_QueryEthBalance(t *testing.T) {
	service, err := newEthService4Test(my_common.KovanUrl)
	if err != nil {
		return
	}
	account := "0xFF33C20B116F77A8F5AccDb372ACDC00a75ceb4a"
	result, err := service.QueryEthBalance(account)
	if err != nil {
		log.Errorf(err, "QueryEthBalance error")
		return
	}
	balanceStr := fmt.Sprintf("%v", result)
	balance := util.ConvertStringToFloat(balanceStr, 0)
	balance = balance / math.Pow10(18)
	log.Infof("QueryEthBalance success, balance: %v", balance)
}

//查询区块
func TestEthService_QueryBlock(t *testing.T) {
	service, err := newEthService4Test(my_common.KovanUrl)
	if err != nil {
		return
	}
	service.QueryBlock()
}

//查询交易
func TestEthService_QueryTransaction(t *testing.T) {
	service, err := newEthService4Test(my_common.KovanUrl)
	if err != nil {
		return
	}
	service.QueryTransaction()
}

//查询交易(根据hash)
func TestEthService_QueryTransactionReceipt(t *testing.T) {
	service, err := newEthService4Test(my_common.KovanUrl)
	if err != nil {
		return
	}
	//hexHash := "0x7576c59c7b2ea03fa310a199372d93f73c89047953cf6ce78453b24760a6abed"
	hexHash := "0x18ed9f76cd825ce7d7b72343296331cf85c09d17cbbeba313e23722572d5e285"
	service.QueryTransactionReceipt(hexHash)
}

//转账eth
func TestEthService_TransferETH(t *testing.T) {
	service, err := newEthService4Test(my_common.KovanUrl)
	if err != nil {
		return
	}
	hexToAddress := "0x1F9184646884753A0947cace94F8eE33dbD7B060"
	service.TransferETH(my_common.HexKey, hexToAddress)
}

//0x1f9840a85d5af5bf1d1762f925bdaddc4201f984

//转账token(erc-20)
func TestEthService_TransferToken(t *testing.T) {
	service, err := newEthService4Test(my_common.KovanUrl)
	if err != nil {
		return
	}
	hexTokenAddress := "0x1f9840a85d5af5bf1d1762f925bdaddc4201f984"
	hexToAddress := "0x1F9184646884753A0947cace94F8eE33dbD7B060"
	service.TransferToken(my_common.HexKey, hexTokenAddress, hexToAddress)
}

//0x1fc0b5814170ecd1ea6b02c3b02f9ded9b1b2dd3703bbe8c1e5ed80036ab83ea

//监听订阅新区块 不好使?
func TestEthService_BlockSubscribe(t *testing.T) {
	service, err := newEthService4Test(my_common.KovanWss)
	if err != nil {
		return
	}
	service.BlockSubscribe()
}

//创建裸交易
func TestEthService_CreateRawTransaction(t *testing.T) {
	service, err := newEthService4Test(my_common.KovanUrl)
	if err != nil {
		return
	}
	hexToAddress := "0x1F9184646884753A0947cace94F8eE33dbD7B060"
	service.CreateRawTransaction(my_common.HexKey, hexToAddress)
}

//发送裸交易
func TestEthService_SendRawTransaction(t *testing.T) {
	service, err := newEthService4Test(my_common.KovanUrl)
	if err != nil {
		return
	}
	hexToAddress := "0x1F9184646884753A0947cace94F8eE33dbD7B060"
	hexRawTx, err := service.CreateRawTransaction(my_common.HexKey, hexToAddress)
	if err != nil {
		return
	}
	service.SendRawTransaction(hexRawTx)
}

//0xd6db5af897873e69d3212a180a836594f1e6dddcdc8f4f2bfc824bc9a4ea2d51

//部署demo智能合约
func TestEthService_DeployDemoContract(t *testing.T) {
	service, err := newEthService4Test(my_common.KovanUrl)
	if err != nil {
		return
	}
	service.DeployDemoContract(my_common.HexKey)
}

//contract address: 0x046c812B4205b04CB3f345b0Ee98254F7bA9da7B
//tx hash: 0x30c156e966a0511ea835136631f91681e1eab3424b73a5e40f4b22d8276d4f28

//加载demo智能合约
func TestEthService_LodeDemoContract(t *testing.T) {
	service, err := newEthService4Test(my_common.KovanUrl)
	if err != nil {
		return
	}
	hexContractAddress := "0x046c812B4205b04CB3f345b0Ee98254F7bA9da7B"
	service.LodeDemoContract(hexContractAddress)
}

//demo智能合约-getValue-读操作
func TestEthService_DemoContractGetValue(t *testing.T) {
	service, err := newEthService4Test(my_common.KovanUrl)
	if err != nil {
		return
	}
	hexContractAddress := "0x046c812B4205b04CB3f345b0Ee98254F7bA9da7B"
	service.DemoContractGetValue(hexContractAddress)
}

//demo智能合约-setValue-写操作
func TestEthService_DemoContractSetValue(t *testing.T) {
	service, err := newEthService4Test(my_common.KovanUrl)
	if err != nil {
		return
	}
	hexContractAddress := "0x046c812B4205b04CB3f345b0Ee98254F7bA9da7B"
	data := "大话西游，爱你一万年"
	service.DemoContractSetValue(hexContractAddress, my_common.HexKey, data)
}

//0x510569a819bcb601304eac83fed5568fd12696a5c12ea5db8c287dc432fb50a2

//查询智能合约的字节码
func TestEthService_QueryContractByteCode(t *testing.T) {
	service, err := newEthService4Test(my_common.KovanUrl)
	if err != nil {
		return
	}
	hexContractAddress := "0x046c812B4205b04CB3f345b0Ee98254F7bA9da7B"
	service.QueryContractByteCode(hexContractAddress)
}

//ERC20代币智能合约-读操作
func TestEthService_ERC20ContractRead(t *testing.T) {
	service, err := newEthService4Test(my_common.KovanUrl)
	if err != nil {
		return
	}
	hexContractAddress := "0x1f9840a85d5af5bf1d1762f925bdaddc4201f984"
	hexOwnerAddress := my_common.OwnerAddress
	service.ERC20ContractRead(hexContractAddress, hexOwnerAddress)
}

//ERC20代币智能合约-订阅日志事件
func TestEthService_ERC20ContractSubscribe(t *testing.T) {
	service, err := newEthService4Test(my_common.KovanWss)
	if err != nil {
		return
	}
	hexContractAddress := "0x1f9840a85d5aF5bf1D1762F925BDADdC4201F984"
	service.ERC20ContractSubscribe(hexContractAddress)
}

func newEthService4Test(url string) (*EthService, *errno.Errno) {
	client, err := my_common.NewClientDial(url)
	if err != nil {
		log.Errorf(err, "NewClientDial error")
		return nil, errno.HandleError
	}
	//defer client.Close()
	service := NewEthService(client)
	return service, nil
}
