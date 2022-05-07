package middle

import (
	"lj-chain-center/pkg"
	"testing"
)

//cfda0a5abbe7c7d134a8334f46e47ec08edb87e990f24bba2f6863891999aaf6
//0xfad5aE00694C519C4b711eB5282ab53FB9D7F7c4
//0x336a4ed40f51e643d5f6c9ae26eff5026b888fcec1839104e9e31f636a8992422af23256a5ecd057733b8aa2bfd43abb58126e7e73a7d57300fd1a357e645c0e01
const (
	private = "cfda0a5abbe7c7d134a8334f46e47ec08edb87e990f24bba2f6863891999aaf6"
	address = "0xfad5aE00694C519C4b711eB5282ab53FB9D7F7c4"
)

func TestSignAndVerify(t *testing.T) {
	initConfig()
	data := "admin"
	hexDataHash, hexDataSign, err := sign(private, data)
	if err != nil {
		t.Errorf("sign error")
		return
	}

	t.Logf("hexDataHash:%v, hexDataSign:%v", hexDataHash, hexDataSign)

	match := verifySignSub(address, hexDataHash, hexDataSign)
	t.Logf("match:%v", match)
}

func initConfig() {
	pkg.Init("../../conf/config-local.yaml")
}
