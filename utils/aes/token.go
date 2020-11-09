package aes

import (
	"bytes"
	"compress/gzip"
	"crypto/md5"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
	"time"
)

type tokenObj struct {
	Tm   int64 `json:"tm"`
	Hx   string `json:"hx"`
	Data string `json:"data"`
}

const tokenKey = "token密钥"
const tokenHx = "tokenHx"
const tokenTime = 3600 //token有效期
//生成token
func EncryptToken(data string) (string, error) {
	var err error
	now := time.Now().Unix()
	has := fmt.Sprintf("%x", md5.Sum([]byte(data +tokenHx+strconv.FormatInt(now, 10))))
	obj := &tokenObj{
		Tm:   now,
		Data: data,
		Hx:   has,
	}
	bys, _ := json.Marshal(obj)
	token, ok := EncryptWithAES(tokenKey, string(bys))
	if !ok {
		return "", fmt.Errorf("AES加密失败")
	}
	//字符串压缩
	var b bytes.Buffer
	gz,err := gzip.NewWriterLevel(&b,gzip.BestCompression)
	if err != nil{
		return "",err
	}
	_, err = gz.Write([]byte(token))
	if err != nil {
		return "", err
	}
	err = gz.Flush()
	if err != nil {
		return "", err
	}
	err = gz.Close()
	if err != nil {
		return "", err
	}
	token = base64.StdEncoding.EncodeToString(b.Bytes())
	return token, nil
}

//解密token
func DecryptToken(token string)(data string,err error){
	bys, err := base64.StdEncoding.DecodeString(token)
	if err != nil{
		return "",err
	}
	rdata := bytes.NewReader(bys)
	r, err := gzip.NewReader(rdata)
	if err != nil{
		return "",err
	}
	s, err := ioutil.ReadAll(r)
	if err != nil{
		return "",err
	}
	defer func() {
		if recover() != nil{
			err = fmt.Errorf("aes反序列化恐慌")
		}
	}()
	token,ok := DecryptWithAES(tokenKey,string(s) + "sdasdsdf")
	if err != nil{
		return "",err
	}
	if !ok{
		return "",fmt.Errorf("aes反序列化失败")
	}
	obj := &tokenObj{}
	err = json.Unmarshal([]byte(token),obj)
	if err != nil{
		return "",err
	}
	//时间校验
	if tm := time.Now().Unix() - obj.Tm;tm<0||tm>= tokenTime {
		return "",fmt.Errorf("token过期")
	}
	//哈希校验
	hax := fmt.Sprintf("%x", md5.Sum([]byte(obj.Data +tokenHx+strconv.FormatInt(obj.Tm, 10))))
	if hax != obj.Hx{
		return "",fmt.Errorf("哈希校验失败")
	}
	return obj.Data,nil
}
