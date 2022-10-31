package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

func examplePanic() {
	var data = []byte(`{"status": 200}`)
	var result map[string]interface{}

	if err := json.Unmarshal(data, &result); err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("%T\n", result["status"]) // float64
	var status = result["status"].(int)  // 类型断言错误
	fmt.Println("Status value: ", status)
}

// 将 decode 后需要的 float 值转为 int 使用
func example2() {
	var data = []byte(`{"status": 200}`)
	var result map[string]interface{}

	if err := json.Unmarshal(data, &result); err != nil {
		log.Fatalln(err)
	}

	var status = uint64(result["status"].(float64))
	fmt.Println("Status value: ", status)
}

// 使用 Decoder 类型来 decode JSON 数据，明确表示字段的值类型
func example3() {
	var data = []byte(`{"status": 200}`)
	var result map[string]interface{}

	var decoder = json.NewDecoder(bytes.NewReader(data))
	decoder.UseNumber()

	if err := decoder.Decode(&result); err != nil {
		log.Fatalln(err)
	}

	var status, _ = result["status"].(json.Number).Int64()
	fmt.Println("Status value: ", status)
}

// 你可以使用 string 来存储数值数据，在 decode 时再决定按 int 还是 float 使用
// 将数据转为 decode 为 string
func example4() {
	var data = []byte(`{"status": 200}`)
	var result map[string]interface{}

	var decoder = json.NewDecoder(bytes.NewReader(data))
	decoder.UseNumber()

	if err := decoder.Decode(&result); err != nil {
		log.Fatalln(err)
	}

	var status uint64
	err := json.Unmarshal([]byte(result["status"].(json.Number).String()), &status)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Status value: ", status)
}

// 使用 struct 类型将你需要的数据映射为数值型
func example5() {
	var data = []byte(`{"status": 200}`)

	var result struct {
		Status uint64 `json:"status"`
	}

	err := json.NewDecoder(bytes.NewReader(data)).Decode(&result)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("Result: %+v\n", result)
}

// 可以使用 struct 将数值类型映射为 json.RawMessage 原生数据类型
// 适用于如果 JSON 数据不着急 decode 或 JSON 某个字段的值类型不固定等情况
func example6() {
	records := [][]byte{
		[]byte(`{"status":200, "tag":"one"}`),
		[]byte(`{"status":"ok", "tag":"two"}`),
	}

	for idx, record := range records {
		var result struct {
			StatusCode uint64
			StatusName string
			Status     json.RawMessage `json:"status"`
			Tag        string          `json:"tag"`
		}

		err := json.NewDecoder(bytes.NewBuffer(record)).Decode(&result)
		if err != nil {
			log.Fatalln(err)
		}

		var code uint64
		err = json.Unmarshal(result.Status, &code)
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Printf("[%v] result => %+v\n", idx, result)
	}
}
