package main

import (
	"encoding/json"
	"fmt"
)

type course1 struct {
	Name string `json:"name"` //Name,Price and Tags should have N,P,T capital otherwise there would be empty json object if we are not writing `json:"name"`

	Price int `json:"-"`

	Tags []string `json:"tags,omitempty"`
}

func main() {

	//var courses []course1

	//var map1 = make(map[string]interface{})

	var map1 []map[string]interface{}
	jsonData := []byte(`[{"name":"ReactJS","tags":["frontend"]},{"name":"NodeJs","tags":["backend"]},{"name":"AngularJS"}]`)

	checkJson := json.Valid(jsonData)

	fmt.Println(checkJson)

	//map1 := make(map[string]interface{})

	err := json.Unmarshal(jsonData, &map1)

	if err != nil {

		fmt.Println(err)
	}

	fmt.Printf("%v\n", map1)

	for k, v := range map1 {
		fmt.Println(k, v)
	}
}
