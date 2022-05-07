package service

import (
	"errors"
	"fmt"
	"github.com/jinzhu/copier"
	"lj-chain-center/common/log"
	"lj-chain-center/dao"
	"lj-chain-center/model"
	"lj-chain-center/pkg"
	"time"
)

type CoreService struct {
	companyDIDDao  *dao.CompanyDIDDao
	personDIDDao   *dao.PersonDIDDao
	antBassService *AntBassService
	lubanService   *LubanService
}

func NewCoreService() *CoreService {
	return &CoreService{dao.NewCompanyDIDDao(), dao.NewPersonDIDDao(), NewAntBassService(), NewLubanService()}
}

//创建公司DID
func (service *CoreService) CreateCompanyDID(req *model.CreateDIDModelReq) ([]*model.CreateDIDModelResp, error) {
	//ant bass chain
	antBassDID, err := service.createCompanyDIDSub(pkg.ANT_BASS_CHAIN_TYPE, req)
	if err != nil {
		log.Errorf(err, "CreateCompanyDID service.createCompanyDID error")
		return nil, errors.New(pkg.HANDLE_ERROR)
	}
	//luban chain
	lubanDID, err := service.createCompanyDIDSub(pkg.LUBAN_CHAIN_TYPE, req)
	if err != nil {
		log.Errorf(err, "CreateCompanyDID service.createCompanyDID error")
		return nil, errors.New(pkg.HANDLE_ERROR)
	}

	list := make([]*model.CreateDIDModelResp, 0)
	antBassItem := &model.CreateDIDModelResp{
		BizID:     req.BizID,
		Identity:  req.Identify,
		ChainType: pkg.ANT_BASS_CHAIN_TYPE,
		DID:       antBassDID,
	}
	list = append(list, antBassItem)

	lubanItem := &model.CreateDIDModelResp{
		BizID:     req.BizID,
		Identity:  req.Identify,
		ChainType: pkg.LUBAN_CHAIN_TYPE,
		DID:       lubanDID,
	}
	list = append(list, lubanItem)

	return list, nil
}

func (service *CoreService) createCompanyDIDSub(chainType int, req *model.CreateDIDModelReq) (string, error) {
	timeStr := fmt.Sprintf(`%v`, time.Now().UnixNano())
	req.OwnerUID = "LJ" + timeStr
	log.Infof("createCompanyDIDSub req.OwnerUID:%v", req.OwnerUID)
	//1:ant bass chain 2:luban chain
	item, err := service.companyDIDDao.GetCompanyDIDByIdentity(req.Identify, chainType)
	if err != nil {
		log.Errorf(err, "createCompanyDIDSub service.companyDIDDao.GetCompanyDIDByIdentity error")
		return "", errors.New(pkg.HANDLE_ERROR)
	}
	if item.DID != "" {
		return item.DID, nil

	}
	did := ""
	if chainType == pkg.ANT_BASS_CHAIN_TYPE {
		did, err = service.antBassService.CreateAntBassDID(pkg.COMPANY_TYPE, req)
		if err != nil {
			log.Errorf(err, "createCompanyDIDSub service.antBassService.CreateAntBassDID error")
			return "", errors.New(pkg.HANDLE_ERROR)
		}
	} else {
		did, err = service.lubanService.CreateLubanDID(req)
		if err != nil {
			log.Errorf(err, "createCompanyDIDSub service.lubanClient.CreateLubanDID error")
			return "", errors.New(pkg.HANDLE_ERROR)
		}
	}
	log.Infof("createCompanyDIDSub createCompanyDIDSub chainType %v, did:%v", chainType, did)
	item.BizID = req.BizID
	item.Identity = req.Identify
	item.ChainType = chainType
	item.OwnerUID = req.OwnerUID
	item.DID = did
	item.Status = pkg.VALID
	if err := service.companyDIDDao.AddCompanyDID(item); err != nil {
		log.Errorf(err, "createCompanyDIDSub service.companyDIDDao.AddCompanyDID error")
		return "", errors.New(pkg.HANDLE_ERROR)
	}

	return item.DID, nil
}

//创建个人DID
func (service *CoreService) CreatePersonDID(req *model.CreateDIDModelReq) ([]*model.CreateDIDModelResp, error) {
	//ant bass chain
	antBassDID, err := service.createPersonDIDSub(pkg.ANT_BASS_CHAIN_TYPE, req)
	if err != nil {
		log.Errorf(err, "CreatePersonDID service.createCompanyDID error")
		return nil, errors.New(pkg.HANDLE_ERROR)
	}
	lubanDID, err := service.createPersonDIDSub(pkg.LUBAN_CHAIN_TYPE, req)
	if err != nil {
		log.Errorf(err, "CreatePersonDID service.createCompanyDID error")
		return nil, errors.New(pkg.HANDLE_ERROR)
	}

	list := make([]*model.CreateDIDModelResp, 0)
	antBassItem := &model.CreateDIDModelResp{
		BizID:     req.BizID,
		Identity:  req.Identify,
		ChainType: pkg.ANT_BASS_CHAIN_TYPE,
		DID:       antBassDID,
	}
	list = append(list, antBassItem)

	lubanItem := &model.CreateDIDModelResp{
		BizID:     req.BizID,
		Identity:  req.Identify,
		ChainType: pkg.LUBAN_CHAIN_TYPE,
		DID:       lubanDID,
	}
	list = append(list, lubanItem)

	return list, nil
}

func (service *CoreService) createPersonDIDSub(chainType int, req *model.CreateDIDModelReq) (string, error) {
	timeStr := fmt.Sprintf(`%v`, time.Now().UnixNano())
	req.OwnerUID = "LJ" + timeStr
	log.Infof("createPersonDIDSub req.OwnerUID:%v", req.OwnerUID)
	//1:ant bass chain 2:luban chain
	item, err := service.personDIDDao.GetPersonDIDByIdentity(req.Identify, chainType)
	if err != nil {
		log.Errorf(err, "createPersonDIDSub service.personDIDDao.GetPersonDIDByIdentity error")
		return "", errors.New(pkg.HANDLE_ERROR)
	}
	if item.DID == "" {
		did := ""
		if chainType == pkg.ANT_BASS_CHAIN_TYPE {
			did, err = service.antBassService.CreateAntBassDID(pkg.PERSON_TYPE, req)
			if err != nil {
				log.Errorf(err, "createPersonDIDSub service.antBassService.CreateAntBassDID error")
				return "", errors.New(pkg.HANDLE_ERROR)
			}
		} else {
			did, err = service.lubanService.CreateLubanDID(req)
			if err != nil {
				log.Errorf(err, "createPersonDIDSub service.lubanClient.CreateLubanDID error")
				return "", errors.New(pkg.HANDLE_ERROR)
			}
		}
		log.Infof("createPersonDIDSub createCompanyDIDSub chainType %v, did:%v", chainType, did)
		item.BizID = req.BizID
		item.Identity = req.Identify
		item.ChainType = chainType
		item.OwnerUID = req.OwnerUID
		item.DID = did
		item.Status = pkg.VALID
		if err := service.personDIDDao.AddPersonDID(item); err != nil {
			log.Errorf(err, "createPersonDIDSub service.personDIDDao.AddPersonDID error")
			return "", errors.New(pkg.HANDLE_ERROR)
		}
	}

	return item.DID, nil
}

func (service *CoreService) GetCompanyDIDList(identity string) ([]*model.CompanyDIDModel, error) {
	entityList, err := service.companyDIDDao.GetCompanyDIDListByIdentity(identity)
	if err != nil {
		log.Errorf(err, "GetCompanyDIDList service.companyDIDDao.GetCompanyDIDListByIdentity error")
		return nil, errors.New(pkg.HANDLE_ERROR)
	}
	resList := make([]*model.CompanyDIDModel, 0)
	for _, entity := range entityList {
		item := &model.CompanyDIDModel{}
		err = copier.Copy(item, entity)
		if err != nil {
			log.Errorf(err, "GetCompanyDIDList copier.Copy error")
		}
		resList = append(resList, item)
	}

	return resList, nil
}

func (service *CoreService) GetPersonDIDList(identity string) ([]*model.PersonDIDModel, error) {
	entityList, err := service.personDIDDao.GetPersonDIDListByIdentity(identity)
	if err != nil {
		log.Errorf(err, "GetPersonDIDList service.personDIDDao.GetPersonDIDListByIdentity error")
		return nil, errors.New(pkg.HANDLE_ERROR)
	}
	resList := make([]*model.PersonDIDModel, 0)
	for _, entity := range entityList {
		item := &model.PersonDIDModel{}
		err = copier.Copy(item, entity)
		if err != nil {
			log.Errorf(err, "GetPersonDIDList copier.Copy error")
		}
		resList = append(resList, item)
	}

	return resList, nil
}
