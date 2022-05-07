package ipfs

import (
	shell "github.com/ipfs/go-ipfs-api"
	my_common "go-chain-center/client/common"
	"go-chain-center/rest/pkg"
	"testing"
)

func TestIpfsService_UnloadFile(t *testing.T) {
	service := newIpfsService4Test()
	filePath := "/Users/vincecfl/go/src/go-flash-chain/module/ipfs/file/ccc.txt"
	service.UnloadFile(filePath)
}

func TestIpfsService_UnloadDir(t *testing.T) {
	service := newIpfsService4Test()
	dirPath := "/Users/vincecfl/go/src/go-flash-chain/module/ipfs/file"
	service.UnloadDir(dirPath)
}

func newIpfsService4Test() *IpfsService {
	pkg.Init("../../conf/config-dev.yaml")
	service := NewIpfsService(shell.NewShell(my_common.IpfsLocalUrl))
	return service
}
