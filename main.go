package main

import (
	"log"
	"translater/service"
)

func main() {
	fileString := service.ReadFile("G:/minecraft/translater/en_us.json")
	jsonMap, err := service.JsonStringToMap(fileString)
	if err != nil {
		log.Fatalf("parse json error! %s", fileString)
	}
	//fmt.Println(jsonMap)
	q := ""
	for _, value := range jsonMap {
		q += (string(value) + "\n")
	}
	jsonMapTrans := service.Query(q)
	jsonMapOutput := map[string]string{}
	for key := range jsonMap {
		jsonMapOutput[key] = jsonMapTrans[jsonMap[key]]
	}
	// mjson, _ := json.Marshal(jsonMapOutput)
	// mString := string(mjson)
	//fmt.Println(jsonMapOutput)
	service.WriteFile(jsonMapOutput)
}
