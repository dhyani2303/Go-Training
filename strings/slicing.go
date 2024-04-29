package main

import "fmt"

func main() {

	// Slicing a string is efficient because it reuses the same backing array
	// Slicing returns bytes not runes

	s1 := "abcdefg"

	fmt.Println(s1[0:3]) //abc will be answer because they are ascii values

	//since s2 does not contain ascii code the bytes corresponding to 0 and 1 are returned
	s2 := "中文维基是世界上"
	fmt.Println(s2[0:2]) // -> � - the unicode representation of bytes from index 0 and 1.

	// returning a slice of runes
	// 1st step: converting string to rune slice

	rs := []rune(s2)

	fmt.Printf("%T\n", rs)

	// 2st step: slicing the rune slice
	fmt.Println(rs[0:3]) // [20013 25991 32500] these are unicode corresponding to the three characters
	fmt.Println(string(rs[0:3]))

}
