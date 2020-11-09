package aes

import (
	"fmt"
	"testing"
)

func TestAes(t *testing.T)  {
	val := "1234567890"
	fmt.Println("加密前：",val)
	v,err := EncryptToken(val)
	if err != nil{
		panic(fmt.Sprintf("err：%s",err.Error()))
	}
	fmt.Println("加密后:",v)
	s,err := DecryptToken(v)
	if err != nil{
		panic(fmt.Sprintf("err：%s",err.Error()))
	}
	fmt.Println("解密后:",s)
}