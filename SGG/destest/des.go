package destest

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"fmt"
)

func paddingText1(src []byte, blockSize int) []byte {
	padding := blockSize - len(src)%blockSize //所填值、数量
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	newText := append(src, padText...)
	return newText
}
func unpaddingText1(src []byte) []byte {
	len := len(src)
	num := int(src[len-1])
	newText := src[:len-num]
	fmt.Println(newText)
	return newText
}

// EncryptDES1 is a function
func EncryptDES1(src, key []byte) []byte {
	// func NewCipher(key []byte) (cipher.Block, error)
	Block, err := des.NewCipher(key)
	if err != nil {
		panic(err)
	}
	src = paddingText1(src, Block.BlockSize()) //填充操作
	iv := []byte("aaaabbbb")
	BlockMode := cipher.NewCBCEncrypter(Block, iv)
	BlockMode.CryptBlocks(src, src)
	return src
}

// DecrypteDES1 is a function
func DecrypteDES1(src, key []byte) []byte {
	Block, err := des.NewCipher(key)
	if err != nil {
		panic(err)
	}
	iv := []byte("aaaabbbb")
	BlockMode := cipher.NewCBCDecrypter(Block, iv)
	BlockMode.CryptBlocks(src, src)
	newText := unpaddingText1(src)
	return newText

}

//填充分组函数
func paddingText(src []byte, blockSize int) []byte { //src -->原始数据
	//blockSize 每个分组数据长度
	//1.求最后一个分组要填充多少个字节
	padding := blockSize - len(src)%blockSize //求要填多少个
	//2.创建切片并初始化，字节值为padding,字节数为padding
	//bytes.Repeat(b[]byte ,count int)[]byte==>把切片b赋值count个然后合成新的字节切片返回
	padText := bytes.Repeat([]byte{byte(padding)}, padding) //初始化切片，{}强制转换
	//3.切片与src连接
	//append()
	// append的用法有两种：
	// slice = append(slice, elem1, elem2)
	// slice = append(slice, anotherSlice...)

	// 第一种用法中，第一个参数为slice,后面可以添加多个参数。
	// 如果是将两个slice拼接在一起，则需要使用第二种用法，在第二个slice的名称后面加三个点
	// ，而且这时候append只支持两个参数，不支持任意个数的参数。
	newText := append(src, padText...) //切片后加...代表是切片  newText为连接后data
	return newText

}

//删除末尾填充字节函数
func unpaddingText(src []byte) []byte {
	//1.求处理的长度
	len := len(src)
	//2.取出末尾字符整型值
	number := int(src[len-1]) //转换为int类型
	//3.将末尾num个字节删掉
	newText := src[:len-number] //左闭右开-->[0,len-number)***

	return newText
}

//No Error

//使用des完成对称加密
//[!]DES函数中默认const blockSize == 8
func encryptDES(src, key []byte) []byte { //key :与src等长
	//1.创建并返回一个使用DES算法的cipher.Block接口
	// func NewCipher(key []byte) (cipher.Block, error)
	//[!]Block接口代表一个使用特定密钥的底层块加/解密器。它提供了加密和解密独立数据块的能力。
	//type Block interface{
	// 	BlockSize()int
	// 	Encrypt(dst,src []byte)
	// 	Decrypt(dst,src []byte)
	// }
	block, err := des.NewCipher(key)
	if err != nil {
		panic(err)
		//panic 能够改变程序的控制流，函数调用panic 时会立刻停止执行函数的其他代码
	}
	//2.对最后一个明文分组补充数据	
	src = paddingText(src, block.BlockSize()) //调用之前写的函数填充，block.BlockSize()一般为8
	//3.创建一个链式的，底层使用DES加密的接口*************************[!][!]
	iv := []byte("aaaabbbb")
	blockMode := cipher.NewCBCEncrypter(block, iv)
	//func NewCBCEncrypter()
	// func NewCBCEncrypter(b Block, iv []byte) BlockMode
	// 返回一个密码分组链接模式的、底层用b加密的BlockMode接口，初始向量iv的长度必须等于b的块尺寸。********

	//4.encryter 连续数据块
	//dst := make([]byte, len(src))   //求内存***************此处报错**********
	blockMode.CryptBlocks(src, src) //studygolang/pkgdoc   1.dst，src
	return src
}

//使用DES解密
func decrypteDES(src []byte, key []byte) []byte {
	//1.创建并返回一个DES算法cipher.Block接口
	block, err := des.NewCipher(key)
	if err != nil {
		panic(err) //panic 函数
	}

	//2.创建一个密码分组为链接模式的，底层使用DES解密的BlockMode接口
	iv := []byte("aaaabbbb") //同上的初始化向量
	blockMode := cipher.NewCBCDecrypter(block, iv)
	//3.解密

	blockMode.CryptBlocks(src, src) //可以覆盖同一个内存地址
	//4.去填充内容-->func unpaddingText
	newText := unpaddingText(src)
	return newText
	//return src
}
