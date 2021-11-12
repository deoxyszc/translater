package service

import (
	"bufio"
	"io"
	"log"
	"os"
)

func ReadFile(fileName string) string {
	f, err := os.Open(fileName)
	if err != nil {
		log.Printf("打开文件失败:%#v", err)
	}
	defer f.Close()
	r := bufio.NewReader(f)
	var dataBlocks []byte
	buf := make([]byte, 102400) // 每次读取字节数
	for {
		n, err := r.Read(buf) // 读取字节数 n
		if err != nil {
			if err == io.EOF {
				// 判断文件读取结束
				break
			}
			log.Printf("打开文件失败:%#v", err)
		}
		dataBlocks = append(dataBlocks, buf[:n]...) // 注意有人这里[:n] 是读的字节数赋值，最后一次读取可能小于buf定义量
	}
	return string(dataBlocks)
}

func WriteFile(m map[string]string) {
	fp, err := os.OpenFile("zh_cn.json", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal(err)
	}
	defer fp.Close()
	data := MapToJson(m)
	_, err = fp.Write(data)
	if err != nil {
		log.Fatal(err)
	}
}
