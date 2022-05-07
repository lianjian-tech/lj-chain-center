package model

type LubanBlockReq struct {
	Page
}

type LubanBlockResp struct {
	Total int64         `json:"total"`
	Rows  []*LubanBlock `json:"rows"`
}

type LubanBlock struct {
	BlockNum   int64  `json:"blockNum"`
	BlockHash  string `json:"blockHash"`
	Timestamp  int64  `json:"timestamp"`
	TxCount    int64  `json:"txCount"`
	ParentHash string `json:"parentHash"`
}

type LubanTransactionReq struct {
	Page
	BlockNum int64 `form:"blockNum"`
}

type LubanTransactionResp struct {
	Total int64               `json:"total"`
	Rows  []*LubanTransaction `json:"rows"`
}

type LubanTransaction struct {
	TxHash      string `json:"txHash"`
	TxTypeNum   int64  `json:"txTypeNum"`
	TxTypeName  string `json:"txTypeName"`
	Timestamp   int64  `json:"timestamp"`
	BlockNum    int64  `json:"blockNum"`
	BlockHash   string `json:"blockHash"`
	FromAddress string `json:"fromAddress"`
	ToAddress   string `json:"toAddress"`
	TxData      string `json:"txData"`
}

type LubanRealData struct {
	BlockNum  int64 `json:"blockNum"`
	BizAmount int64 `json:"bizAmount"`
	RunDay    int64 `json:"runDay"`
	//运行状态 0:异常 1:正常
	ChainStatus int64 `json:"chainStatus"`
}

type LubanTxDataResp struct {
	Total int64          `json:"total"`
	Rows  []*LubanTxData `json:"rows"`
}

type LubanTxData struct {
	Datetime int64 `json:"datetime"`
	TxAmount int64 `json:"txAmount"`
}
type LubanDeposit4KxReq struct {
	Data string `form:"data" binding:"required" json:"data"`
	Sign string `form:"sign" binding:"required" json:"sign"`
}

type LubanQueryTxData struct {
	TxHash      string `json:"txHash"`
	BlockNumber int64  `json:"blockNumber"`
	Timestamp   int64  `json:"timestamp"`
}

//***************************************

type LubanBaseCallResp struct {
	Code    int64       `json:"code"`
	Msg     string      `json:"msg"`
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
}

type LubanCreateDIDCallReq struct {
	OwnerUID string `json:"ownerUid"`
}

type LubanDepositModelReq struct {
	DID  string `form:"did" binding:"required" json:"did"`
	Data string `form:"data" binding:"required" json:"data"`
}
