package main

import (
	"encoding/json"
	"fmt"
)

//字符串 --> 结构体
type Message struct {
	Uid  string
	Name string
}

//字段
type Field struct {
	Len []uint
}

func main() {
	strMsg := []byte(`{"Uid":"101","Name":"Beyond"}`)

	var stMsg Message
	err := json.Unmarshal(strMsg, &stMsg)
	fmt.Println(err)
	fmt.Println(strMsg)

	strNumber := `[1,32]`
	var stField Field
	err = json.Unmarshal([]byte(strNumber), &stField.Len)
	fmt.Println(err)
	fmt.Println(stField)

	strJson := `[
				{"name":"张三","age":20,"address":["youku.com","https://github.com/"]},
				{"name":"李四","age":20,"address":["taobao.com","https://github.com/"]},
				{"name":"王五","age":20,"address":["bilibili.com","https://github.com/"]}
	]`
	mapSlice := make([]map[string]interface{}, 0)
	err = json.Unmarshal([]byte(strJson), &mapSlice)
	fmt.Println(err)
	if err != nil {
		fmt.Println("反序列化失败", err)
	} else {
		fmt.Println(mapSlice)
		for _, mapVal := range mapSlice {
			for key, val := range mapVal {
				fmt.Println(key, val)
			}
		}
	}

	strArticle := `[{"Id":100,"Title":"汉书"},{"Id":200,"Title":"楚辞"},{"Id":300,"Title":"春秋","Test":100}]`
	var slcArticle []map[string]interface{}
	err = json.Unmarshal([]byte(strArticle), &slcArticle)
	if err != nil {
		fmt.Println("反序列化出错", err)
	}
	for i, mapArticle := range slcArticle {
		fmt.Println("第", i, "个值是", mapArticle, mapArticle["Id"], mapArticle["Title"])
	}
}
