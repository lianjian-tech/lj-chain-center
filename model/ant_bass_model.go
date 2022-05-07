package model

type AntBassBaseCallResp struct {
	Code    int64       `json:"code"`
	Msg     string      `json:"msg"`
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
}

type AntBassCreateDIDCallReq struct {
	Type     int    `json:"type"`
	OwnerUID string `json:"ownerUid"`
}

type AntBassDepositModelReq struct {
	DID     string      `form:"did" binding:"required" json:"did"`
	ReqData interface{} `form:"reqData" binding:"required" json:"reqData"`
}
