package main

import (
	"fmt"
	"strconv"
)

func main() {

	//the scope of result and err is valid till the below block only
	//if result,err := strconv.Atoi("45") is known as intialization statement.it is used to declare/initialise local variables
	//err==nil is boolean expression which returns true or false
	if result, err := strconv.Atoi("45"); err == nil {
		fmt.Println(result)
	} else {
		fmt.Println(err)
	}

}
