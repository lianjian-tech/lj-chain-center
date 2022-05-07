package dao

import (
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"lj-chain-center/common/log"
	"lj-chain-center/entity"
	"lj-chain-center/pkg"
)

type PersonDIDDao struct {
	db *sqlx.DB
}

func NewPersonDIDDao() *PersonDIDDao {
	return &PersonDIDDao{
		db: pkg.DB,
	}
}

func (dao *PersonDIDDao) GetPersonDIDByIdentity(identity string, chainType int) (*entity.PersonDIDEntity, error) {
	sql := fmt.Sprintf(`select id, biz_id, identity, chain_type, owner_uid, did, status `)
	sql += fmt.Sprintf(`from cc_person_did `)
	sql += fmt.Sprintf(`where identity='%v' and chain_type=%v and status=1 `, identity, chainType)
	log.Infof("GetPersonDIDByIdentity sql:%v", sql)

	rows, err := dao.db.Queryx(sql)
	if err != nil {
		log.Errorf(err, "GetPersonDIDByIdentity dao.db.Queryx error")
		return nil, errors.New(pkg.HANDLE_ERROR)
	}
	defer rows.Close()
	item := &entity.PersonDIDEntity{}
	if rows.Next() {
		rows.StructScan(item)
	}

	return item, nil
}

func (dao *PersonDIDDao) AddPersonDID(item *entity.PersonDIDEntity) error {
	sql := fmt.Sprintf(`insert into cc_person_did(biz_id, identity, chain_type, owner_uid, did, status) `)
	sql += fmt.Sprintf(`values('%v', '%v', %v, '%v', '%v', %v)`, item.BizID, item.Identity, item.ChainType, item.OwnerUID, item.DID, item.Status)
	log.Infof("AddPersonDID sql:%v", sql)
	if _, err := dao.db.Exec(sql); err != nil {
		log.Errorf(err, "AddPersonDID dao.db.Exec error")
		return errors.New(pkg.HANDLE_ERROR)
	}
	return nil
}

func (dao *PersonDIDDao) GetPersonDIDListByIdentity(identity string) ([]*entity.PersonDIDEntity, error) {
	sql := fmt.Sprintf(`select id, biz_id, identity, chain_type, owner_uid, did, status `)
	sql += fmt.Sprintf(`from cc_person_did `)
	sql += fmt.Sprintf(`where identity='%v' and status=1 `, identity)
	sql += fmt.Sprintf(`order by chain_type asc `)

	log.Infof("GetPersonDIDListByIdentity sql:%v", sql)

	rows, err := dao.db.Queryx(sql)
	if err != nil {
		log.Errorf(err, "GetPersonDIDListByIdentity dao.db.Queryx error")
		return nil, errors.New(pkg.HANDLE_ERROR)
	}
	defer rows.Close()
	list := make([]*entity.PersonDIDEntity, 0)
	for rows.Next() {
		item := &entity.PersonDIDEntity{}
		rows.StructScan(item)
		list = append(list, item)
	}

	return list, nil
}
