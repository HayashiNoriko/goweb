package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type Signature struct {
	Name string
}

func main() {
	// 1. 序列化
	signature := Signature{Name: "test"}

	msg, err := json.Marshal(signature)
	if err != nil {
		fmt.Printf("JSON marshal error: %s\n", err)
	}

	fmt.Printf("JSON message: %s\n", msg)

	// 2. 反序列化
	signature2 := new(Signature)
	delivery := []byte(`{"Name":"test2"}`)

	// 用 Unmarshal 也可以
	// if err := json.Unmarshal(delivery, signature2); err != nil {
	// 	fmt.Printf("JSON decode error: %s\n", err)
	// }
	// fmt.Printf("Signature: %+v\n", signature2)

	// 但 machinery 框架中用了更高级的方法：
	// json.Decoder.Decode，适用于流式处理（如读取文件、网络流），可以多次调用 Decode 逐个解码多个 JSON 对象
	decoder := json.NewDecoder(bytes.NewReader(delivery))
	// 让数字以 json.Number 类型处理，避免精度丢失
	decoder.UseNumber()
	if err := decoder.Decode(signature2); err != nil {
		fmt.Printf("JSON decode error: %s\n", err)
	}

	fmt.Printf("Signature: %+v\n", signature2)
}
