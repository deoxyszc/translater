package service

import (
	"bytes"
	"crypto/md5"
	"crypto/rand"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
	"net/url"
)

/* appid, key是百度api提供的；salt是随机生成的字符串，可以由字母和数字组成，其长度由saltlen配置 */
var (
	appid   string
	salt    string
	key     string
	saltlen int
)

func init() {
	saltlen = 10
	appid = "appid"
	salt = createSalt(saltlen)
	key = "key"
}

func createSalt(len int) string {
	var container string
	var str = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	b := bytes.NewBufferString(str)
	length := b.Len()
	bigInt := big.NewInt(int64(length))
	for i := 0; i < len; i++ {
		randomInt, _ := rand.Int(rand.Reader, bigInt)
		container += string(str[randomInt.Int64()])
	}
	return container
}

func Query(q string) {
	data := []byte(appid + q + salt + key)
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has)
	params := url.Values{}
	Url, _ := url.Parse("https://fanyi-api.baidu.com/api/trans/vip/translate")
	params.Set("q", q)
	params.Set("from", "en")
	params.Set("to", "zh")
	params.Set("appid", "20211109000994741")
	params.Set("salt", salt)
	params.Set("sign", md5str)
	//如果参数中有中文参数,这个方法会进行URLEncode
	Url.RawQuery = params.Encode()
	urlPath := Url.String()
	fmt.Println(urlPath)
	resp, _ := http.Get(urlPath)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	res := string(body)
	fmt.Println(string(res))
	dst := GetTranslateResult(res)
	fmt.Println(string(dst))
}
