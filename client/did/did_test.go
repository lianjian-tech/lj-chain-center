package did

import (
	"testing"
)

func TestCreateDid(t *testing.T) {
	privateKey, address, err := CreateKey()
	if err != nil {
		t.Error("CreateKey error ")
		return
	}
	t.Logf("privateKey:%v, address:%v", privateKey, address)
}
