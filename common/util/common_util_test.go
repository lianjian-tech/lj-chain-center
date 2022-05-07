package util

import (
	"fmt"
	"go-chain-center/model"
	"go-chain-center/pkg"
	"testing"
)

func TestGetDatetime(t *testing.T) {
	value := GetDatetime(pkg.YYYY_MM_DD_HH_MM_SS)
	fmt.Println("value:", value)
}

const (
	privateKey = "MIIEvAIBADANBgkqhkiG9w0BAQEFAASCBKYwggSiAgEAAoIBAQCN2E9TDuyGLm0MmsjT7gXww36Pf+nTj8vMQYMCVdGqx2fuXxAGp+3SZ5X63Xmo76X3ybv6Ya+bDhvJh/+RrwV+eafHvbOF+sprn/y6h03bYEorfLAVPoPjCDuchsGAL8MPbPsSXY1+ATVZSpAh4orAs4IPJQHo+UGAqVv8ZQPFqVIjzt/6512txKdIUrkkHeGlqnN3Ga+qj2Lt2T0eh1lH88RkxZ0JSvvSRdfHwzDnb6j6Mn5WeTciALT9eowh96Q9HLKGRx49/8rtIeyzQEhNEFBKXW+4oyHn4H+DNmdf+7tFLqHPXGxlYJ6gNARnmQEB1M6jrd/O44ef/w8QvU1RAgMBAAECggEABvDp2M+nJV3KOk6suyovojt+yvjgcVjdCS6/gXK/otLFNDX0q8615ZDqFu+Vj33CqmB10WfMVxMWqwL7r55X73t2iWCwE0NVenTGwj3ywmRw3LXqS+/WNlrlx+AUQfiUXAHiO2DWlq7qdGcgrHfQgGOIXsNZ15Oy79/Zv30V9SskB7vSgsN0EjIGRdniJlGuembNldFJnJbU3ZZEmv4wNRFG13empcO263gREbY3abFzG4M8WD/U4zjAT/cAXYl0N2hpRGZw8nQ5z+LqTld+zQerHcyuqtXy7wDAvG+QqzKUcr8M3JA3LGf/rgaJmrS5xYvXm9IpaHz4n6ULheEvbQKBgQDjZMhaTUJfvyh78zpRNSBv8jzB3XNTkjxSkAv6KNnM5/4b4rmknbIto05kWyWoubBqbznZlDmIzngrE5Rws4dIO6TNIKBzpegpG5G75+R4ADufN0LUbdBZ2q9Y53DfpHSfeVT5XDHFVMkrjpZxICjX+0rnNQsqAevcR9BHrR+XWwKBgQCfsG3gZJttE7ihKiMFz98gbpi4ftazHVFuXgfWiTYxg10KEzW1u6ICsaRPR/Pkn6s5W3KgzwG5vLE1AItRv+AOBwGKfbOmc5kq3q0m4CENAVu+m4qk6OiDKoNP3PcDexa+IpSj7kBOaLR/wlTiSWhMmx/iMMoLuIB7KRy5ldl5wwKBgB594S/atE/Kfk9Aqk1BZnwca3FDlLdcIKKhljmRSMNlSHnMyT/9tWrRPy3N2xro63MaPxOiAKULqql8EqeaK0XQWaT3/wwC70Kz47LiwYwl29jZvChoIUwrOuAdCNG0Z78Ksg5OAo6HdAJXOG0q9s8fb0fZoQmdAFeleZRi7p2JAoGAF5qFczCLxMCwfqnyyU4jhfCzBDpDU/BWdADUGcLoTsNDAlU6kCUu2kUVq15QT4/GSFEbDWfrsl+Qtevhq/C3lns0oBrabhUSI8MMRv3EBN8Zh4AIKz39+D1VcX2QyWPtRado3x49RRm27Fd2IoiFMSWQvedGRhs8anf2NfxPS+8CgYA2mnx5p5NJbDO23q/pLcGy7nT2vcbZMkl5JpoykZCf3Mt0bLip2DjqbR1EeFECJN2EGTE0VwJfStEhzSm3vzpSZjtTf+ns5bxH/4DcgC5uw4VZEt5PcN2Y1C3o6McXxrpQR+2+eb02IQfhd+k8tmQz4ypUDBpvkcu8lsyvktTKWA=="
	publicKey  = "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAjdhPUw7shi5tDJrI0+4F8MN+j3/p04/LzEGDAlXRqsdn7l8QBqft0meV+t15qO+l98m7+mGvmw4byYf/ka8Ffnmnx72zhfrKa5/8uodN22BKK3ywFT6D4wg7nIbBgC/DD2z7El2NfgE1WUqQIeKKwLOCDyUB6PlBgKlb/GUDxalSI87f+uddrcSnSFK5JB3hpapzdxmvqo9i7dk9HodZR/PEZMWdCUr70kXXx8Mw52+o+jJ+Vnk3IgC0/XqMIfekPRyyhkcePf/K7SHss0BITRBQSl1vuKMh5+B/gzZnX/u7RS6hz1xsZWCeoDQEZ5kBAdTOo63fzuOHn/8PEL1NUQIDAQAB"
)

func TestSignMd5(t *testing.T) {
	req := &model.LubanDeposit4KxReq{Data: "测试数据"}
	sign, err := SignMd5(req, privateKey)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("sign:%v", sign)
}

func TestVerifySignMd5(t *testing.T) {
	req := &model.LubanDeposit4KxReq{Data: "测试数据"}
	sign, err := SignMd5(req, privateKey)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("sign:%v", sign)
	req.Sign = sign
	flag := VerifySignMd5(req, publicKey)
	t.Logf("flag:%v", flag)

}

func initConfig() {
	pkg.Init("../../conf/config-dev.yaml")
}
