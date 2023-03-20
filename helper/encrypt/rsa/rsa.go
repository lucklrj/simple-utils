package rsa

import (
	"encoding/base64"

	"github.com/wumansgy/goEncrypt"
)

func Encode(content []byte, publicKey []byte) (cryptText string, err error) {

	contentEncrypt, err := goEncrypt.RsaEncrypt(content, publicKey)
	if err != nil {
		return "", err
	} else {
		return base64.StdEncoding.EncodeToString(contentEncrypt), nil
	}
}
func Decode(content string, privateKey []byte) (cryptText string, err error) {
	contentByte, err := base64.StdEncoding.DecodeString(content)
	if err != nil {
		return "", err
	}
	contentDecode, err := goEncrypt.RsaDecrypt(contentByte, privateKey)
	return string(contentDecode), err
}
