package entity

type LubanBlock struct {
	BlockNum        int64  `json:"blockNum"`
	BlockHash       string `json:"blockHash"`
	Timestamp       int64  `json:"timestamp"`
	TxCount         int64  `json:"txCount"`
	ParentHash      string `json:"parentHash"`
	Version         string `json:"version"`
	GasUsed         string `json:"gasUsed"`
	StateRoot       string `json:"stateRoot"`
	ReceiptRoot     string `json:"receiptRoot"`
	TransactionRoot string `json:"transactionRoot"`
}

type LubanTransaction struct {
	TxHash         string `json:"txHash"`
	TxTypeNum      int64  `json:"txTypeNum"`
	TxTypeName     string `json:"txTypeName"`
	Timestamp      int64  `json:"timestamp"`
	BlockNum       int64  `json:"blockNum"`
	BlockHash      string `json:"blockHash"`
	Nonce          string `json:"nonce"`
	FromAddress    string `json:"fromAddress"`
	ToAddress      string `json:"toAddress"`
	GasUsed        string `json:"gasUsed"`
	Sign           string `json:"sign"`
	TxOriginalData string `json:"txOriginalData"`
	TxParseData    string `json:"txParseData"`
}

type LubanRealData struct {
	BlockNum    int64 `json:"blockNum"`
	BizAmount   int64 `json:"bizAmount"`
	RunDay      int64 `json:"runDay"`
	ChainStatus int64 `json:"chainStatus"`
}

type LubanTxData struct {
	Datetime int64 `json:"datetime"`
	TxAmount int64 `json:"txAmount"`
}
