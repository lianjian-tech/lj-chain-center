package ipfs

import (
	"bytes"
	"github.com/ipfs/go-ipfs-api"
	"go-chain-center/common/errno"
	"go-chain-center/common/log"
	"io/ioutil"
	"os"
)

type IpfsService struct {
	shell *shell.Shell
}

func NewIpfsService(shell *shell.Shell) *IpfsService {
	return &IpfsService{shell: shell}
}

func (service *IpfsService) UnloadFile(filePath string) (string, *errno.Errno) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Errorf(err, "open file error")
		return "", errno.HandleError
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Errorf(err, "read all error")
		return "", errno.HandleError
	}
	hash, err := service.shell.Add(bytes.NewBuffer(data))
	if err != nil {
		log.Errorf(err, "add error")
		return "", errno.HandleError
	}
	log.Infof("hash: %v", hash)
	return hash, nil
}

func (service *IpfsService) UnloadDir(dirPath string) (string, *errno.Errno) {
	hash, err := service.shell.AddDir(dirPath)
	if err != nil {
		log.Errorf(err, "add dir error")
		return "", errno.HandleError
	}
	log.Infof("hash: %v", hash)
	return hash, nil
}
