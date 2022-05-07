package dao

import (
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"lj-chain-center/common/log"
	"lj-chain-center/common/util"
	"lj-chain-center/entity"
	"lj-chain-center/pkg"
)

type CompanyDIDDao struct {
	db *sqlx.DB
}

func NewCompanyDIDDao() *CompanyDIDDao {
	return &CompanyDIDDao{
		db: pkg.DB,
	}
}

func (dao *CompanyDIDDao) GetCompanyDIDByIdentity(identity string, chainType int) (*entity.CompanyDIDEntity, error) {
	sql := fmt.Sprintf(`select id, biz_id, identity, chain_type, owner_uid, did, status `)
	sql += fmt.Sprintf(`from cc_company_did `)
	sql += fmt.Sprintf(`where identity='%v' and chain_type=%v and status=1 `, identity, chainType)
	log.Infof("GetCompanyDIDByIdentity sql:%v", sql)

	rows, err := dao.db.Queryx(sql)
	if err != nil {
		log.Errorf(err, "GetCompanyDIDByIdentity dao.db.Queryx error")
		return nil, errors.New(pkg.HANDLE_ERROR)
	}
	defer rows.Close()
	item := &entity.CompanyDIDEntity{}
	if rows.Next() {
		rows.StructScan(item)
	}

	return item, nil
}

func (dao *CompanyDIDDao) AddCompanyDID(item *entity.CompanyDIDEntity) error {
	log.Infof("AddCompanyDID item:%v", util.ToJSONStr(item))
	sql := fmt.Sprintf(`insert into cc_company_did(biz_id, identity, chain_type, owner_uid, did, status) `)
	sql += fmt.Sprintf(`values('%v', '%v', %v, '%v', '%v', %v)`, item.BizID, item.Identity, item.ChainType, item.OwnerUID, item.DID, item.Status)
	log.Infof("AddCompanyDID sql:%v", sql)
	if _, err := dao.db.Exec(sql); err != nil {
		log.Errorf(err, "AddCompanyDID dao.db.Exec error")
		return errors.New(pkg.HANDLE_ERROR)
	}
	return nil
}

func (dao *CompanyDIDDao) GetCompanyDIDByID(did string, chainType int) (*entity.CompanyDIDEntity, error) {
	sql := fmt.Sprintf(`select id, biz_id, identity, chain_type, owner_uid, did, status `)
	sql += fmt.Sprintf(`from cc_company_did `)
	sql += fmt.Sprintf(`where did='%v' and chain_type=%v and status=1 `, did, chainType)
	log.Infof("GetCompanyDIDByID sql:%v", sql)

	rows, err := dao.db.Queryx(sql)
	if err != nil {
		log.Errorf(err, "GetCompanyDIDByID dao.db.Queryx error")
		return nil, errors.New(pkg.HANDLE_ERROR)
	}
	defer rows.Close()
	item := &entity.CompanyDIDEntity{}
	if rows.Next() {
		rows.StructScan(item)
	}

	return item, nil
}

func (dao *CompanyDIDDao) GetCompanyDIDListByIdentity(identity string) ([]*entity.CompanyDIDEntity, error) {
	sql := fmt.Sprintf(`select id, biz_id, identity, chain_type, owner_uid, did, status `)
	sql += fmt.Sprintf(`from cc_company_did `)
	sql += fmt.Sprintf(`where identity='%v' and status=1 `, identity)
	sql += fmt.Sprintf(`order by chain_type asc `)
	log.Infof("GetCompanyDIDListByIdentity sql:%v", sql)

	rows, err := dao.db.Queryx(sql)
	if err != nil {
		log.Errorf(err, "GetCompanyDIDListByIdentity dao.db.Queryx error")
		return nil, errors.New(pkg.HANDLE_ERROR)
	}
	defer rows.Close()
	list := make([]*entity.CompanyDIDEntity, 0)
	for rows.Next() {
		item := &entity.CompanyDIDEntity{}
		rows.StructScan(item)
		list = append(list, item)
	}

	return list, nil
}
