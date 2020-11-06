package des

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
)

var iv []byte = []byte("PhoenixY")

func paddingText(src []byte, blockSize int) []byte {
	paddingText := blockSize - len(src)%blockSize
	padText := bytes.Repeat([]byte{byte(paddingText)}, paddingText)
	newText := append(src, padText...)
	return newText
}

func unpaddingText(src []byte) []byte {
	unpadnum := int(src[len(src)-1])
	newText := src[:len(src)-unpadnum]//注意此处是左闭右开*********
	return newText
}

//EnDes is a function
func EnDes(src []byte, key []byte) []byte {
	block, err := des.NewCipher(key) //
	if err != nil {
		panic(err)
	}
	src = paddingText(src, block.BlockSize())
	// iv := []byte("PhoenixY")
	blockMode := cipher.NewCBCEncrypter(block, iv)

	blockMode.CryptBlocks(src, src)
	return src
}

//DeDes is a function
func DeDes(src, key []byte) []byte {
	block, err := des.NewCipher(key)
	if err != nil {
		panic(err)
	}
	blockMode := cipher.NewCBCDecrypter(block, iv)

	blockMode.CryptBlocks(src, src)
	newText := unpaddingText(src)
	return newText
}
