package service

import (
	"fmt"
	"github.com/jinzhu/copier"
	"lj-chain-center/common/log"
	"lj-chain-center/common/util"
	"lj-chain-center/model"
	"lj-chain-center/pkg"
	"testing"
)

func initConf() {
	pkg.Init("../conf/config-local.yaml")
}

type User struct {
	Name string
	Age  int
}

type Employee struct {
	Name string
	Age  int
	Role string
}

func TestAAA(t *testing.T) {
	user := User{Name: "dj", Age: 18}
	employee := Employee{}

	copier.Copy(&employee, &user)
	fmt.Printf("%#v\n", employee)
}

func TestLubanService_QueryTx(t *testing.T) {
	initConf()
	service := NewLubanService()
	txHash := "8bda4664c9deacbb2be618772a93451758ce4c1ec6673519cd963a5d1c472af5"
	data, err := service.QueryTx(txHash)
	if err != nil {
		log.Errorf(err, "service.QueryTx error")
		return
	}
	log.Infof("data:%v", util.ToJSONStr(data))
}

func TestLubanService_DepositData(t *testing.T) {
	initConf()
	service := NewLubanService()
	req := &model.LubanDepositModelReq{}
	req.DID = "7d999c56cdeaafd6e208cd14beeca831976a0ddc7046248f9460a04b8b190ad4"
	req.Data = "测试数据"
	data, err := service.DepositData(req)
	if err != nil {
		log.Errorf(err, "service.DepositData error")
		return
	}
	log.Infof("data:%v", data)
}
