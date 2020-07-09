package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)
//  下面代码输出什么
//  float64
func main() {
	json_str := []byte(`{"age":1}`)
	var value map[string]interface{}
	json.Unmarshal(json_str, &value)
	age := value["age"]
	fmt.Println(reflect.TypeOf(age))

}
