package main

import (
	"fmt"
	"reflect"
)

type resume struct {
	Name string `json:"name" doc:"我的名字"`
}

func findDoc(x interface{}) map[string]string {
	t := reflect.TypeOf(x).Elem()
	doc := make(map[string]string)

	for i := 0; i < t.NumField(); i++ {
		doc[t.Field(i).Tag.Get("json")] = t.Field(i).Tag.Get("doc")
	}

	return doc

}

func main() {
	var x resume
	doc := findDoc(&x)
	fmt.Println(doc)
	fmt.Printf("name字段为：%s\n", doc["name"])
}
