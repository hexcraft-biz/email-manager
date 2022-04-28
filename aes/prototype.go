package aes

import (
	"crypto/aes"
	"github.com/hexcraft-biz/email-manager/misc"
)

func GenSalt() ([]byte, error) {
	return misc.GenRandomSalt(aes.BlockSize)
}

func Encrypt(key []byte, text string) ([]byte, error) {
	if c, err := aes.NewCipher(key); err != nil {
		return nil, err
	} else {
		out := make([]byte, len(text))
		c.Encrypt(out, []byte(text))
		return out, nil
	}
}

func Decrypt(key []byte, enc []byte) (string, error) {
	if c, err := aes.NewCipher(key); err != nil {
		return "", err
	} else {
		pt := make([]byte, len(enc))
		c.Decrypt(pt, enc)
		return string(pt[:]), nil
	}
}
