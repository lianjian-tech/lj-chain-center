package usercontract

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/pkg/errors"

	"git.huawei.com/poissonsearch/wienerchain/contract/docker-container/contract-go/contractapi"
)

type Contract struct {
}

// package必须是usercontract，实例化函数必须是NewSmartContract，才能部署成功？
func NewSmartContract() contractapi.Contract {
	return &Contract{}
}

// Init init function implemented by smart contract developer
func (e Contract) Init(stub contractapi.ContractStub) ([]byte, error) {
	fmt.Printf("Enter Contract init function\n")
	args := stub.Parameters()
	const numOfArgs = 2
	if len(args) < numOfArgs {
		return nil, errors.New("init parameter is not correct")
	}
	key := args[0]
	value := args[1]
	err := stub.PutKV(string(key), value)
	if err != nil {
		return nil, errors.New("init put kv failed")
	}

	return nil, nil
}

func (e Contract) Invoke(stub contractapi.ContractStub) ([]byte, error) {
	funcName := stub.FuncName()
	args := stub.Parameters()

	switch funcName {
	case "initMarble":
		return initMarble(stub, args)
	case "initMarbleCommon":
		return initMarbleCommon(stub, args)
	case "getMarbleCommon":
		return getMarbleCommon(stub, args)
	case "getMarble":
		return getMarble(stub, args)
	case "delMarble":
		return deleteMarble(stub, args)
	case "initRange":
		return initRange(stub, args)
	case "getRange":
		return getRange(stub, args)
	case "getMarbleCom":
		return getMarbleCom(stub, args)
	case "initMarbleAndIndex":
		return initMarbleAndIndex(stub, args)
	case "transferMarblesByIndex":
		return transferMarblesByIndex(stub, args)
	case "getPartRangeAndPutKV":
		return getPartRangeAndPutKV(stub, args)
	case "getJSONMarble":
		return getJSONMarble(stub)
	case "getRangeAndPutKV":
		return getRangeAndPutKV(stub, args)
	case "writeMarble":
		return writeMarble(stub, args)
	case "deleteComIndexOneRow":
		return deleteComIndexOneRow(stub, args)
	case "getKeyHistory":
		return getKeyHistory(stub, args)
	}

	return nil, errors.Errorf("func name is not correct, the function name is %s ", funcName)
}

func deleteComIndexOneRow(stubInterface contractapi.ContractStub, args [][]byte) ([]byte, error) {
	const numOfArgs = 3
	if len(args) < numOfArgs {
		return nil, errors.New("the number of args is not correct")
	}

	indexName := string(args[0])
	attr := string(args[1])
	objectName := string(args[2])

	err := stubInterface.DelComIndexOneRow(indexName, []string{attr}, objectName)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func initMarbleCommon(stubInterface contractapi.ContractStub, args [][]byte) ([]byte, error) {
	const argNum = 4
	if len(args) < argNum {
		fmt.Printf("The argNum for init marble is not correct\n")
		return nil, errors.New("the argNum for init marble is not correct")
	}
	marbleName := string(args[0])
	marbleOwner := string(args[1])
	marbleColor := string(args[2])
	marbleSizeStr := string(args[3])
	marbleSize, err := strconv.Atoi(marbleSizeStr)
	if err != nil {
		fmt.Printf("The marble size is not int type\n")
		return nil, errors.New("the marble size is not int type")
	}

	marbleInfo := &marble{Name: marbleName, Color: marbleColor,
		Size: marbleSize, Owner: marbleOwner, ObjectType: "marble"}

	err = stubInterface.PutKVCommon(marbleInfo.Name, marbleInfo)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return nil, nil
}

func getMarbleCommon(stubInterface contractapi.ContractStub, args [][]byte) ([]byte, error) {
	if len(args) < 1 {
		fmt.Printf("enter getMarbleCommon\n")
		return nil, errors.New("not enough parameters")
	}
	key := string(args[0])
	value, err := stubInterface.GetKV(key)
	if err != nil {
		fmt.Printf("getkv error, the err is :%s\n", err.Error())
		return nil, errors.New("get kv failed")
	}
	unmarshal, err := Unmarshal(value)
	if err != nil {
		fmt.Printf("Unmarshal error, the err is :%s\n", err.Error())
		return nil, errors.New("unmarshal value failed")
	}
	m, ok := unmarshal.(*marble)
	if !ok {
		fmt.Printf("the value is not marble\n")
		return nil, errors.New("the value is not marble struct1111")
	}
	fmt.Printf("the marble name is %s\n", m.Name)
	return []byte(m.Name), nil
}

func initMarble(stubInterface contractapi.ContractStub, args [][]byte) ([]byte, error) {
	const numOfArgs = 2
	if len(args) < numOfArgs {
		fmt.Println("The args number is not correct")
		return nil, errors.New("the args is not correct")
	}
	key := args[0]
	keyStr := string(key)

	value, err := stubInterface.GetKV(keyStr)
	if err != nil {
		return nil, errors.Errorf("Get the marble info err:%s", err.Error())
	}
	if value != nil {
		fmt.Printf("The key to be add is already exist")
	}
	value = args[1]

	err = stubInterface.PutKV(keyStr, value)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func writeMarble(stubInterface contractapi.ContractStub, args [][]byte) ([]byte, error) {
	key := args[0]
	keyStr := string(key)
	value := args[1]

	err := stubInterface.PutKV(keyStr, value)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func getMarble(stubInterface contractapi.ContractStub, args [][]byte) ([]byte, error) {
	if len(args) < 1 {
		fmt.Println("The args number is not correct")
		return nil, errors.New("the args is not correct")
	}

	key := args[0]
	value, err := stubInterface.GetKV(string(key))
	if err != nil {
		errInfo := fmt.Sprintf("get the key: %s failed", key)
		return nil, errors.New(errInfo)
	}

	return value, nil
}

func deleteMarble(stubInterface contractapi.ContractStub, args [][]byte) ([]byte, error) {
	if len(args) < 1 {
		return nil, errors.New("the args for delete marble is not correct")
	}
	fmt.Printf("Enter the delete marble\n")

	key := args[0]
	value, err := stubInterface.GetKV(string(key))
	if err != nil {
		errInfo := fmt.Sprintf("get the marble info err:%s", err.Error())
		return nil, errors.New(errInfo)
	}
	if value == nil {
		return nil, errors.New("the key to be delete is not exist")
	}
	stubInterface.Debugf("Before the delete marble\n")
	err = stubInterface.DelKV(string(key))
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func initRange(stubInterface contractapi.ContractStub, args [][]byte) ([]byte, error) {
	stubInterface.Debugf("Enter initRange\n")

	if len(args) < 1 {
		fmt.Printf("The args for initRange is not correct\n")
		return nil, errors.New("The args for initRange is not correct")
	}
	numberStr := string(args[0])
	number, err := strconv.Atoi(numberStr)
	if err != nil {
		fmt.Printf("The arg for initRange is not integer\n")
		return nil, errors.New("the arg for initRange is not integer")
	}

	stubInterface.Debugf("The number is %d\n", number)

	for i := 0; i < number; i++ {
		var key string
		switch {
		case 0 <= i && i <= 9:
			key = fmt.Sprintf("marble00%d", i)
		case 10 <= i && i <= 99:
			key = fmt.Sprintf("marble0%d", i)
		case 100 <= i && i <= 999:
			key = fmt.Sprintf("marble%d", i)
		default:
			return nil, errors.New("invalid number")
		}

		value := "white"
		err := stubInterface.PutKV(key, []byte(value))
		if err != nil {
			return nil, err
		}
	}
	return nil, nil
}

type keyHistory struct {
	Value     string
	BlockNum  uint64
	TxNum     int32
	TxHash    []byte
	IsDeleted bool
	Timestamp uint64
}

func getKeyHistory(stubInterface contractapi.ContractStub, args [][]byte) ([]byte, error) {
	stubInterface.Debugf("Enter getKeyHistory")
	const numOfArgs = 1
	if len(args) != numOfArgs {
		stubInterface.Errorf("the args for getKeyHistory is not correct")
		return nil, errors.New("the args for getKeyHistory is not correct")
	}

	key := string(args[0])

	iterator, err := stubInterface.GetKeyHistoryIterator(key)
	if err != nil {
		return nil, err
	}
	defer iterator.Close()

	var historyArray []keyHistory
	for {
		b := iterator.Next()
		if !b {
			stubInterface.Debugf("the iterator break")
			break
		}

		var history keyHistory
		history.Value = string(iterator.Value())
		history.TxHash = iterator.TxHash()
		history.BlockNum, history.TxNum = iterator.Version()
		history.IsDeleted = iterator.IsDeleted()
		history.Timestamp = iterator.Timestamp()
		historyArray = append(historyArray, history)
	}

	historyMapBytes, err := json.Marshal(historyArray)
	if err != nil {
		return nil, err
	}
	stubInterface.Debugf("historyArray is %v", historyArray)

	return historyMapBytes, nil
}
func getRange(stubInterface contractapi.ContractStub, args [][]byte) ([]byte, error) {
	stubInterface.Debugf("Enter getRange\n")
	begin, end := getBeginAndEnd(args)

	iterator, err := stubInterface.GetIterator(begin, end)
	if err != nil {
		return nil, err
	}
	defer iterator.Close()

	rangeMap := make(map[string]string)
	var count = 0
	for {
		b := iterator.Next()
		if b {
			key := iterator.Key()
			value := string(iterator.Value())
			rangeMap[key] = value
			count++
			stubInterface.Debugf("The iterator read key is %s, value is %s, count is %d\n", key, value, count)
		} else {
			stubInterface.Debugf("The iterator break\n")
			break
		}
	}
	rangeMapBytes, err := json.Marshal(rangeMap)
	if err != nil {
		return nil, err
	}
	stubInterface.Debugf("rangeMap is %v\n", rangeMap)

	return rangeMapBytes, nil
}

func getRangeAndPutKV(stubInterface contractapi.ContractStub, args [][]byte) ([]byte, error) {
	stubInterface.Debugf("Enter getRangeAndPutKV\n")
	beginKey, endKey := getBeginAndEnd(args)

	iterator, err := stubInterface.GetIterator(beginKey, endKey)
	if err != nil {
		return nil, err
	}
	defer iterator.Close()

	rangeMap := make(map[string]string)
	var count = 0
	for {
		b := iterator.Next()
		if b {
			key := iterator.Key()
			value := string(iterator.Value())
			rangeMap[key] = value
			count++
			stubInterface.Debugf("The iterator read key is %s, value is %s, count is %d\n", key, value, count)
		} else {
			stubInterface.Debugf("The iterator break\n")
			break
		}
	}
	rangeBytes, err := json.Marshal(rangeMap)
	if err != nil {
		return nil, err
	}
	stubInterface.Debugf("getRangeAndPutKV rangeMap is %v\n", rangeMap)

	err = stubInterface.PutKV("jsonMarble", rangeBytes)
	if err != nil {
		return nil, errors.New("putstate failed")
	}

	return rangeBytes, nil
}

func getPartRangeAndPutKV(stubInterface contractapi.ContractStub, args [][]byte) ([]byte, error) {
	stubInterface.Debugf("Enter getPartRange\n")
	beginKey, endKey := getBeginAndEnd(args)
	iterator, err := stubInterface.GetIterator(beginKey, endKey)
	if err != nil {
		return nil, err
	}
	defer iterator.Close()

	rangeMap := make(map[string]string)
	var count = 2

	for iterator.Next() {
		key := iterator.Key()
		value := iterator.Value()

		rangeMap[key] = string(value)
		count--
		if count <= 0 {
			break
		}
	}
	stubInterface.Debugf("getPartRange is %v\n", rangeMap)
	rangeMapBytes, err := json.Marshal(rangeMap)
	if err != nil {
		return nil, err
	}
	err = stubInterface.PutKV("jsonMarble", rangeMapBytes)
	if err != nil {
		return nil, errors.New("putstate failed")
	}
	return nil, nil
}

func getJSONMarble(stub contractapi.ContractStub) ([]byte, error) {
	value, err := stub.GetKV("jsonMarble")
	if err != nil {
		stub.Errorf("Get json marble failed\n")
		return nil, errors.New("get json marble failed")
	}
	return value, nil
}

type marble struct {
	ObjectType string
	Name       string
	Color      string
	Size       int
	Owner      string
}

func (m marble) Marshal() ([]byte, error) {
	return json.Marshal(m)
}

// Unmarshal Unmarshal json data.
func Unmarshal(data []byte) (interface{}, error) {
	var value marble
	err := json.Unmarshal(data, &value)
	if err != nil {
		return nil, err
	}
	return &value, nil
}

var storeKey string
var storeValue []byte
var readKey string
var readValue []byte

func getMarbleCom(stubInterface contractapi.ContractStub, args [][]byte) ([]byte, error) {
	if len(args) < 1 {
		fmt.Println("The args number is not correct")
		return nil, errors.New("the args is not correct")
	}

	key := args[0]
	fmt.Printf("getMarbleCom the key is %s\n", string(key))
	value, err := stubInterface.GetKV(string(key))
	if err != nil {
		errInfo := fmt.Sprintf("get the key: %s failed", key)
		return nil, errors.New(errInfo)
	}
	readKey = string(key)
	readValue = value
	compare := bytes.Compare([]byte(storeKey), []byte(readKey))
	if compare == 0 {
		fmt.Printf("key .Compare equal\n")
	}
	i := bytes.Compare(storeValue, readValue)
	if i == 0 {
		fmt.Printf("value .Compare equal\n")
	}

	mb := &marble{}
	err = json.Unmarshal(value, mb)
	if err != nil {
		fmt.Printf("getMarbleCom marble unmarshal failed\n")
		return nil, errors.New("getMarbleCom marble unmarshal failed")
	}

	sprintf := fmt.Sprintf("marble name is %s, marble size %d, marble owner %s, "+
		"                          marble color %s, marble object :%s",
		mb.Name, mb.Size, mb.Owner, mb.Color, mb.ObjectType)

	return []byte(sprintf), nil
}

func initMarbleAndIndex(stubInterface contractapi.ContractStub, args [][]byte) ([]byte, error) {
	const numOfArgs = 4
	if len(args) < numOfArgs {
		fmt.Printf("The args number for init marble is not correct\n")
		return nil, errors.New("the args number for init marble is not correct")
	}
	marbleName := string(args[0])
	marbleOwner := string(args[1])
	marbleColor := string(args[2])
	marbleSizeStr := string(args[3])
	marbleSize, err := strconv.Atoi(marbleSizeStr)
	if err != nil {
		fmt.Printf("The marble size is not int type\n")
		return nil, errors.New("the marble size is not int type")
	}

	marbleInfo := &marble{Name: marbleName, Color: marbleColor, Size: marbleSize,
		Owner: marbleOwner, ObjectType: "marble"}
	marbleBytes, err := json.Marshal(marbleInfo)
	if err != nil {
		fmt.Printf("MarbleInfo marshal error\n")
		return nil, errors.New("marbleInfo marshal error")
	}
	fmt.Printf("The key is marbleInfo.Name:%s\n", marbleInfo.Name)
	err = stubInterface.PutKV(marbleInfo.Name, marbleBytes)
	if err != nil {
		return nil, err
	}

	indexName := "colorindex"
	err = stubInterface.SaveComIndex(indexName, []string{marbleInfo.Color}, marbleInfo.Name)
	if err != nil {
		fmt.Printf("SaveComIndex failed\n")
	}

	return nil, nil
}

func transferMarblesByIndex(stub contractapi.ContractStub, args [][]byte) ([]byte, error) {
	const numOfArgs = 2
	if len(args) < numOfArgs {
		return nil, errors.New("incorrect number of arguments. Expecting 2")
	}

	color := string(args[0])
	newOwner := strings.ToLower(string(args[1]))
	fmt.Println("- start transferMarblesBasedOnColor ", color, newOwner)

	// Query the color~name index by color
	// This will execute a key range query on all keys starting with 'color'
	kvIterator, err := stub.GetKVByComIndex("colorindex", []string{color})
	if err != nil {
		return nil, err
	}

	for {
		hasNext := kvIterator.Next()
		if hasNext {
			key := kvIterator.Key()
			value := kvIterator.Value()
			fmt.Printf("The index has next, final key is %s, value is %s\n", key, string(value))
			marbleToTransfer := marble{}
			err = json.Unmarshal(value, &marbleToTransfer)
			if err != nil {
				fmt.Printf("Unmarshal marble failed\n")
				return nil, err
			}
			marbleToTransfer.Owner = newOwner

			marbleJSONasBytes, err := json.Marshal(marbleToTransfer)
			if err != nil {
				return nil, err
			}
			err = stub.PutKV(key, marbleJSONasBytes)
			if err != nil {
				return nil, err
			}
		} else {
			break
		}
	}

	return nil, nil
}

func getBeginAndEnd(args [][]byte) (string, string) {
	const numOfArgs = 2
	if len(args) < numOfArgs {
		return "", ""
	}
	return string(args[0]), string(args[1])
}
