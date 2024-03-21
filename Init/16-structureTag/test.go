package main

import (
	"fmt"
	"reflect"
)

type Human struct {
	Name string `info:"Jay" doc:"My name"`
	Sex  int    `asn1:"application" doc:"3"`
}

func findTag(x interface{}) {
	t := reflect.TypeOf(x)
	for i := 0; i < t.NumField(); i++ {
		taginfo := t.Field(i).Tag.Get("info")
		tagdoc := t.Field(i).Tag.Get("doc")
		tagAsn1 := t.Field(i).Tag.Get("asn1")
		fmt.Println("info: ", taginfo, "doc:", tagdoc, "asn1: ", tagAsn1)

	}
}

func main() {
	h := Human{"1", 0}
	findTag(h)
}
