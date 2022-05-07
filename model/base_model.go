package model

type Page struct {
	Start int64 `form:"start,default=0" json:"start"`
	Limit int64 `form:"limit,default=10" json:"limit"`
}

type BaseAuthModelReq struct {
	Hash    string      `form:"hash" binding:"required" json:"hash"`
	Sign    string      `form:"sign" binding:"required" json:"sign"`
	DID     string      `form:"did" json:"did"`
	ReqData interface{} `form:"reqData" json:"reqData"`
}
