package crypto

import (
	"encoding/base64"
	"testing"
)

func TestCrypto(t *testing.T) {
	c := New()

	str := "床前明月光，疑似地上霜，举头望明月，低头思故乡"

	text, err := c.Encrypt([]byte(str))
	if err != nil {
		panic(err)
	}

	t.Log("密文：", base64.StdEncoding.EncodeToString(text))

	text, err = c.Encrypt(text)
	if err != nil {
		panic(err)
	}

	t.Log("明文：", string(text))

	if str != string(text) {
		t.Fail()
	}
}
