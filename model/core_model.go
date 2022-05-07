package model

type CreateDIDModelReq struct {
	BaseAuthModelReq
	BizID    string `form:"bizId" binding:"required" json:"bizId"`
	Identify string `form:"identify" binding:"required" json:"identify"`
	OwnerUID string `form:"ownerUid" json:"ownerUid"`
}

type CreateDIDModelResp struct {
	BizID     string `json:"bizId"`
	Identity  string `json:"identity"`
	ChainType int    `json:"chainType"`
	DID       string `json:"did"`
}

type PersonDIDModel struct {
	BizID     string `json:"bizId"`
	Identity  string `json:"identity"`
	ChainType int    `json:"chainType"`
	DID       string `json:"did"`
}

type CompanyDIDModel struct {
	BizID     string `json:"bizId"`
	Identity  string `json:"identity"`
	ChainType int    `json:"chainType"`
	DID       string `json:"did"`
}

type QueryCompanyDIDListReq struct {
	BaseAuthModelReq
	Identify string `form:"identify" binding:"required" json:"identify"`
}

type QueryPersonDIDListReq struct {
	BaseAuthModelReq
	Identify string `form:"identify" binding:"required" json:"identify"`
}
