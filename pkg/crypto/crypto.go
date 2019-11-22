package crypto

import (
	"crypto/aes"
	"crypto/cipher"
)

// Crypto 密码
type Crypto interface {
	Encrypt(plainText []byte) ([]byte, error)
	Decrypt(cipherText []byte) ([]byte, error)
}

// Option 密码选项
type Option struct {
	key []byte
	iv  []byte
}

type crypto struct {
	opt Option
}

// 默认
var nopOpt = Option{
	key: []byte("heeyapp-2019-crypto-$$$###&&&***"),
	iv:  []byte("heeyapp-crypto-8"),
}

// New 创建密码实例
func New(opts ...Option) Crypto {
	if len(opts) == 0 {
		return &crypto{
			opt: nopOpt,
		}
	}

	return &crypto{
		opt: opts[0],
	}
}

// 加密明文
func (c *crypto) Encrypt(plainText []byte) ([]byte, error) {
	block, err := aes.NewCipher(c.opt.key)
	if err != nil {
		return nil, err
	}

	stream := cipher.NewCTR(block, c.opt.iv)
	cipherText := make([]byte, len(plainText))
	stream.XORKeyStream(cipherText, plainText)

	return cipherText, nil
}

// 解密密文
func (c *crypto) Decrypt(cipherText []byte) ([]byte, error) {
	block, err := aes.NewCipher(c.opt.key)
	if err != nil {
		return nil, err
	}

	stream := cipher.NewCTR(block, c.opt.iv)
	plainText := make([]byte, len(cipherText))
	stream.XORKeyStream(plainText, cipherText)

	return plainText, nil
}
