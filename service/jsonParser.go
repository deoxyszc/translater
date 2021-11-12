package service

import (
	"bytes"
	"encoding/json"
	"log"
)

type Result struct {
	Src string `json:"src"`
	Dst string `json:"dst"`
}

type Results struct {
	From        string   `json:"from"`
	To          string   `json:"to"`
	TransResult []Result `json:"trans_result"`
}

//解析json字符串成 map
func JsonStringToMap(jsonStr string) (m map[string]string, err error) {
	ret := map[string]string{}
	unmarsha1Err := json.Unmarshal([]byte(jsonStr), &ret)
	if unmarsha1Err != nil {
		return nil, unmarsha1Err
	}
	//fmt.Println(unmarsha1Err, "转换结果", ret)
	return ret, nil
}

func GetTranslateResult(result string) map[string]string {
	var res Results
	json.Unmarshal([]byte(result), &res)
	ret := map[string]string{}
	for _, result := range res.TransResult {
		ret[result.Src] = result.Dst
	}
	return ret
}

func MapToJson(m map[string]string) []byte {
	data, err := json.Marshal(m)
	if err != nil {
		log.Fatal(err)
	}
	var out bytes.Buffer
	err = json.Indent(&out, data, "", "\t")
	if err != nil {
		log.Fatal(err)
	}
	return out.Bytes()
}
