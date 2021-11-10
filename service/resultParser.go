package service

import "encoding/json"

type Result struct {
	Src string `json:"src"`
	Dst string `json:"dst"`
}

type Results struct {
	From        string   `json:"from"`
	To          string   `json:"to"`
	TransResult []Result `json:"trans_result"`
}

func GetTranslateResult(result string) string {
	var res Results
	json.Unmarshal([]byte(result), &res)
	ret := res.TransResult[0].Dst
	return ret
}
