package main

import (
	"encoding/json"
	"fmt"
)

// encoding to json also known as marshal means converting any data type to json
//`json:""` are used to give define what will key appear once converted to json
//`json:"-"` indicates that the particular value has to be excluded from json object. i.e price wont be present in json object
//omitempty is used for the case where we want to exclude the field if that particular field is empty

type course struct {
	Name string `json:"name"` //Name,Price and Tags should have N,P,T capital otherwise there would be empty json object if we are not writing `json:"name"`

	Price int `json:"-"`

	Tags []string `json:"tags,omitempty"`
}

func main() {

	defer func() {

		r := recover()

		if r != nil {
			fmt.Println(r)

		}
	}()

	//slice of course
	courses := []course{

		{
			Name: "ReactJS", Price: 200, Tags: []string{"frontend"},
		},
		{
			Name: "NodeJs", Price: 300, Tags: []string{"backend"},
		},
		{
			Name: "AngularJS", Price: 400, Tags: nil,
		},
	}
	fmt.Printf("%T %v\n", courses, len(courses))

	fmt.Println(courses)

	jsonData, err := json.Marshal(courses)

	if err != nil {
		panic(err)
	}

	//MarshalIndent is used to make code more readable, first argument is the type ,second is prefix which appears in front of every line, 3rd is indent \t is the ident used in json to make it more readable
	//jsonData1, err1 := json.MarshalIndent(courses, "prefix", "\t")

	//output with prefix
	//prefix  {
	//	prefix          "name": "ReactJS",
	//	prefix          "tags": [
	//	prefix                  "frontend"
	//	prefix          ]
	//	prefix  },
	//	prefix  {
	//		prefix          "name": "NodeJs",
	//		prefix          "tags": [
	//		prefix                  "backend"
	//		prefix          ]
	//		prefix  },
	//	prefix  {
	//		prefix          "name": "AngularJS"
	//		prefix  }
	//prefix]

	//jsonData1, err1 := json.MarshalIndent(courses, "", "\t")
	//if err1 != nil {
	//	panic(err1)
	//}

	fmt.Println(jsonData) //[]byte slice is returned
	fmt.Println(string(jsonData))
	fmt.Printf("%s", jsonData)
	//fmt.Println(string(jsonData1))

	//encoding integer to json
	integer := 10

	jsonData, err = json.Marshal(integer)
	fmt.Println(string(jsonData))

	//encoding float to json
	float := 10.20
	jsonData, err = json.Marshal(float)
	fmt.Println(string(jsonData))

	//encoding slice to json
	slice := []int{1, 2, 3}

	jsonData, err = json.Marshal(slice)

	fmt.Println(string(jsonData))

	//encoding map to json

	map1 := make(map[string]string)

	map1["key"] = "value"

	jsonData, err = json.Marshal(map1)

	fmt.Println(string(jsonData))
}
