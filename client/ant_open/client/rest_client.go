package ant_open_client

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/valyala/fasthttp"
	"go-chain-center/client/ant_open/model"
	"go-chain-center/client/ant_open/response"
	"go-chain-center/client/ant_open/util"
	"time"
)

const (
	ShakeHandPath       = "/api/contract/shakeHand"
	ChainCallForBizPath = "/api/contract/chainCallForBiz"
	ContentType         = "application/json;charset=utf-8"
	TimeOut             = 10 * 1000
)

type RestClientProperties struct {
	BizId        string `json:"bizId"`        // 链的ID
	Account      string `json:"account"`      // 账户名称
	TenantId     string `json:"tenantId"`     // 租户ID
	KmsId        string `json:"kmsId"`        // MyKmsKeyId
	RestUrl      string `json:"restUrl"`      // rest的服务地址
	AccessId     string `json:"accessId"`     // 分配给用户用于访问rest的账户名（可通过开放联盟链控制台获得）
	AccessSecret string `json:"accessSecret"` // 分配给用于用户访问rest的密钥的路径（可通过开放联盟链控制台获得

}

type RestClient struct {
	Properties *RestClientProperties
}

func (client *RestClient) GetRestToken() (string, error) {
	req, resp := fasthttp.AcquireRequest(), fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)

	req.SetRequestURI(client.Properties.RestUrl + ShakeHandPath)
	req.Header.SetContentType(ContentType)
	req.Header.SetMethod("POST")

	nowMill := time.Now().UnixNano() / 1e6
	secret, err := util.Sign(fmt.Sprintf("%v%v", client.Properties.AccessId, nowMill), client.Properties.AccessSecret)
	if err != nil {
		fmt.Errorf("sign error\n")
		return "", errors.New("sign error")
	}
	shakeRequest := &model.ShakeRequest{
		AccessId: client.Properties.AccessId,
		Time:     fmt.Sprintf("%v", nowMill),
		Secret:   secret,
	}
	body, err := json.Marshal(shakeRequest)
	fmt.Println(string(body))
	if err != nil {
		fmt.Errorf("json.Marshal error\n")
		return "", errors.New("json.Marshal error")
	}
	req.SetBody(body)
	httpClient := fasthttp.Client{}
	if err := httpClient.DoTimeout(req, resp, TimeOut*time.Millisecond); err != nil {
		fmt.Errorf("client.DoTimeout error\n")
		return "", errors.New("client.DoTimeout error")
	}
	respToken := &model.RestToken{}
	if err := json.Unmarshal(resp.Body(), respToken); err != nil {
		fmt.Errorf("json.Unmarshal error\n")
		return "", errors.New("json.Unmarshal error")
	}
	if respToken.Success && respToken.Code == "200" && respToken.Data != nil {
		return respToken.Data.(string), nil
	} else {
		return "", nil
	}
}

func (client *RestClient) QueryTransaction(hash string) (*response.BaseResp, error) {
	token, err := client.GetRestToken()
	if err != nil {
		fmt.Errorf("GetRestToken error")
		return nil, errors.New("GetRestToken error")
	}
	param := model.CallRestBizParam{
		BaseParam: model.BaseParam{
			AccessId: client.Properties.AccessId,
			BizId:    client.Properties.BizId,
			Hash:     hash,
			Method:   model.QUERYTRANSACTION,
		},
	}
	param.Token = token

	return client.SendRequest(param, client.Properties.RestUrl+ChainCallForBizPath)
}

func (client *RestClient) QueryReceipt(hash string) (*response.BaseResp, error) {
	token, err := client.GetRestToken()
	if err != nil {
		fmt.Errorf("GetRestToken error")
		return nil, errors.New("GetRestToken error")
	}
	param := model.CallRestBizParam{
		BaseParam: model.BaseParam{
			AccessId: client.Properties.AccessId,
			BizId:    client.Properties.BizId,
			Hash:     hash,
			Method:   model.QUERYRECEIPT,
		},
	}
	param.Token = token

	return client.SendRequest(param, client.Properties.RestUrl+ChainCallForBizPath)
}

func (client *RestClient) Deposit(orderId, content string, gas int64) (*response.BaseResp, error) {
	token, err := client.GetRestToken()
	if err != nil {
		fmt.Errorf("GetRestToken error")
		return nil, errors.New("GetRestToken error")
	}
	param := model.CallRestBizParam{
		BaseParam: model.BaseParam{
			AccessId: client.Properties.AccessId,
			BizId:    client.Properties.BizId,
			Method:   model.DEPOSIT,
		},
		OrderId:    orderId,
		Account:    client.Properties.Account,
		Content:    content,
		MykmsKeyId: client.Properties.KmsId,
		TenantId:   client.Properties.TenantId,
		Gas:        gas,
	}
	param.Token = token
	return client.SendRequest(param, client.Properties.RestUrl+ChainCallForBizPath)
}

func (client *RestClient) DeployContract(orderId, contractName, contractCode string, gas int64) (*response.BaseResp, error) {
	token, err := client.GetRestToken()
	if err != nil {
		fmt.Errorf("GetRestToken error")
		return nil, errors.New("GetRestToken error")
	}
	param := model.CallRestBizParam{
		BaseParam: model.BaseParam{
			AccessId: client.Properties.AccessId,
			BizId:    client.Properties.BizId,
			Method:   model.DEPLOYCONTRACTFORBIZ,
		},
		OrderId:      orderId,
		Account:      client.Properties.Account,
		MykmsKeyId:   client.Properties.KmsId,
		TenantId:     client.Properties.TenantId,
		ContractName: contractName,
		ContractCode: contractCode,
		Gas:          gas,
	}
	param.Token = token
	return client.SendRequest(param, client.Properties.RestUrl+ChainCallForBizPath)
}

func (client *RestClient) CallContract(orderId, contractName, methodSignature, inputParamListStr, outTypes string, gas int64) (*response.BaseResp, error) {
	token, err := client.GetRestToken()
	if err != nil {
		fmt.Errorf("GetRestToken error")
		return nil, errors.New("GetRestToken error")
	}
	param := model.CallRestBizParam{
		BaseParam: model.BaseParam{
			AccessId: client.Properties.AccessId,
			BizId:    client.Properties.BizId,
			Method:   model.CALLCONTRACTBIZ,
		},
		OrderId:           orderId,
		Account:           client.Properties.Account,
		TenantId:          client.Properties.TenantId,
		ContractName:      contractName,
		MethodSignature:   methodSignature,
		InputParamListStr: inputParamListStr,
		OutTypes:          outTypes,
		MykmsKeyId:        client.Properties.KmsId,
		Gas:               gas,
	}
	param.Token = token
	return client.SendRequest(param, client.Properties.RestUrl+ChainCallForBizPath)
}

func (client *RestClient) SendRequest(param interface{}, url string) (*response.BaseResp, error) {
	req, resp := fasthttp.AcquireRequest(), fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)

	body, err := json.Marshal(&param)
	if err != nil {
		return nil, err
	}
	fmt.Printf("SendRequest url: \n%v\n", url)
	fmt.Printf("SendRequest param body: \n%v\n", string(body))
	req.SetRequestURI(url)
	req.Header.SetContentType(ContentType)
	req.Header.SetMethod("POST")

	req.SetBody(body)
	httpClient := fasthttp.Client{}
	if err := httpClient.DoTimeout(req, resp, TimeOut*time.Millisecond); err != nil {
		fmt.Errorf("client.DoTimeout\n")
		return nil, errors.New("client.DoTimeout error")
	}
	baseResp := &response.BaseResp{}
	err = json.Unmarshal(resp.Body(), &baseResp)
	if err != nil {
		fmt.Errorf("json.Unmarshal error\n")
		return nil, errors.New("json.Unmarshal error")
	}
	return baseResp, nil
}
