package entity

type PersonDIDEntity struct {
	ID        int64  `db:"id" json:"id"`
	BizID     string `db:"biz_id" json:"bizId"`
	Identity  string `db:"identity" json:"identity"`
	ChainType int    `db:"chain_type" json:"chainType"`
	OwnerUID  string `db:"owner_uid" json:"ownerUid"`
	DID       string `db:"did" json:"did"`
	Status    int    `db:"status" json:"status"`
}

type CompanyDIDEntity struct {
	ID        int64  `db:"id" json:"id"`
	BizID     string `db:"biz_id" json:"bizId"`
	Identity  string `db:"identity" json:"identity"`
	ChainType int    `db:"chain_type" json:"chainType"`
	OwnerUID  string `db:"owner_uid" json:"ownerUid"`
	DID       string `db:"did" json:"did"`
	Status    int    `db:"status" json:"status"`
}
