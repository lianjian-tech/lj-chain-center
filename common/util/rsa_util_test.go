package util

import (
	"go-chain-center/pkg"
	"testing"
)

func TestRsaSignWithMd5(t *testing.T) {
	initConf()
	data := "我爱你 I Love U"
	privateKeyHex := "MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQDEOmwRjXZjT2EVMCMu5K0raXJCik4W0/rPBLAfU0+uHGi51iRN27yuhc0k1Yr5X8imnsoKbsroavKAywUcpBlPhn1eJpINVKsBOAevVxro7zmzlBtXVQXnWKq2MGnAcO/1PZ2BDKG1D3Q5rxqdaiZ/Hn/N0NvGc+ogRDfAUksQLMaccxR+AnVbckTbuNGvMiCQ28veMgstCtNRkl6+YQ7pCVwg5JgPuHO6rbcuUddmQXiZrg3jHEEcHFpV+inGR/9C4m7gm0I692bQ5RqSreA3ssbAwFMzFzg9isxphHkAPNpJO07YCFgJI/1mOSAfz1JFnXyHyREpPzhTaU/UbFMTAgMBAAECggEAEshIydFRkKXrc/y8kWOAE9SvBCF64gJ+Ukxvk0AFM3Qyrf8KUI2E3OufvDPYbpA1g4VVb3Q1ZdO+zIciQSCP76W+4EEu36nICIDVc3oR8FPdoQu/JjfULdW12WR+6tWkkkPIjK6p18vX4yGYmZKshqRjtU6VisRTl60P7Q+zcaO8oidYesxr43Mzeb4tJWn+Pqzx4rPlicuJVripxAXw19UMil/vPOOr4SmZWtcJlLICX+Ul3CnWatLwcfd6q5uHA6kt0o+D/V9S7EgmUkMfEybE6IPLKya8LxrKNpbe2xuaQ+1iCqv8ivOZmsFGmhBGtTn/J530LPPrnnhUdwgZUQKBgQDitZpDoYTx4ixmdE9Zy+J7js20w/V/zA9ASRI3QdfroFeHqG5WtB+0+r04XTgg6imPaifHWWKwXkcBOcSOdFTvkgFOo1U7Pv8XaGKPw89EVfAjPd7By6th3uno/InJv31S6s6la2yntT69NUhLwzhpJN9BnSBqA09l9rJoRs6vyQKBgQDdlKglC8r7vtmadThMCLpgTD1TOs/GxQubbfQoHIWsqLl8q9zoZkrjsH5pK2/3tbFaWuzCTSzpOyVTCIrKp+18XBwJHF31ztJN36F+BI6mEr0LmwMLglwSNBCc2ZyPGH2v8/dLRhUhSrUc/6W5k5sqCSrEybiJUe0F+HzVtFyx+wKBgAgK80cU8td186dK1wnTftxtS4TvMsVu4tegAAJv1o4oaRKJ+igfScSs5IucWwJn7EOHO5QJ6sNhBNnZR8zZ+a1qG5vDlaF9caVN8bkx9CdzeQp3BnEPnaviMaAvXWAH6BVJ+TzLG9azSl+CBWeSrGymfzkfruez6FwwrVGfuXNpAoGAd4BiAsDOc2ElmJDZkOAwLBlSBZ41Sr0gmm39fNiNG2ZPyzYhSNFKWGvqcSv1EgJSqSlagRwAObUlVxfHSq6wfR5sHFnJLvWhJFCUlxM8H8MBY83xIXbn0wVzZDRUpDQQLSIFX/pnEPJqpT4XEfPSfiQ9ha5JMP4XN94g7Vx9lZsCgYEAzbZZsf5CVVZ3vvTsdc2UBTcI9pKXZAfn3+5GdhhC2M6U57vhWGsWIX+zWM8X3xli7KuQdTHHZeud7/yOt3GK+iwxMu/tyBxD3xO4ih+Ng/7hqvZ16zwRSeEW5BLiirI16A2hQq0wn6D26J70vjN97WukEw8/HFRlYg6Ll1RBMMQ="
	result, err := RsaSignWithMd5(data, privateKeyHex)
	if err != nil {
		t.Errorf("RsaSignWithMd5 error")
		return
	}
	t.Logf("result:%v", result)
}

func TestRsaVerifySignWithMd5(t *testing.T) {
	initConf()
	originalData := "我爱你 I Love U"
	signData := "RyWn5XDZ1hSy260NcoUO55VVMbL2bthcAdOhluennZfvyhJmr48JVSx1bkI009ThknrMEdmEtuLrN2kj/ANLUdrIveNQsW3KcPL+9tjn9bOVk9X73/GQzZTUB/k2e60YIicMT6ng/ubOBIiMkcP7m7VYFskquk0rgnthJxbkEQXS8Qwx4myhf9/sDIpYxny/ag25xrjKMauTqvDwE9FLIOGA7NS/FmfhoZN55d1b6BPTYtbneZYUV9JrySM4/TaXx5CMet5LoZYL/Ot/D2dzj6K4as+z45iKSL4BVUK4VcbU2TCpFZHx+vebuVHCKOHWXV7z/Oj77RVC5VOw8MLFpg=="
	publicKeyHex := "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAxDpsEY12Y09hFTAjLuStK2lyQopOFtP6zwSwH1NPrhxoudYkTdu8roXNJNWK+V/Ipp7KCm7K6GrygMsFHKQZT4Z9XiaSDVSrATgHr1ca6O85s5QbV1UF51iqtjBpwHDv9T2dgQyhtQ90Oa8anWomfx5/zdDbxnPqIEQ3wFJLECzGnHMUfgJ1W3JE27jRrzIgkNvL3jILLQrTUZJevmEO6QlcIOSYD7hzuq23LlHXZkF4ma4N4xxBHBxaVfopxkf/QuJu4JtCOvdm0OUakq3gN7LGwMBTMxc4PYrMaYR5ADzaSTtO2AhYCSP9ZjkgH89SRZ18h8kRKT84U2lP1GxTEwIDAQAB"
	resultFlag := RsaVerifySignWithMd5(originalData, signData, publicKeyHex)
	t.Logf("resultFlag:%v", resultFlag)
}

func initConf() {
	pkg.Init("../../conf/config-dev.yaml")
}
