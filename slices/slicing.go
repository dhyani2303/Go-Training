package main

import "fmt"

func main() {

	// arrays, slices and strings are sliceable
	// slicing doesn't modify the array or the slice, it returns a new one

	// declaring an [5]int array
	a := [5]int{1, 2, 3, 4, 5}

	// a slice expression is formed by specifying a start or a low bound and a stop or high bound like  a[start:stop].
	// this selects a range of elements which includes the element at index start, but excludes the element at index stop.

	// slicing an array returns a slice, not an array

	b := a[1:3]

	fmt.Printf("Type %T,Value %v\n", b, b)

	//slicing in slice
	s1 := []int{2, 3, 4, 5, 6}

	s2 := s1[1:3]

	fmt.Println(s2)

	// a missing low index defaults to zero; a missing high index defaults to the length of the sliced operand.
	fmt.Println(s1[2:]) // => [3 4 5 6], same as s1[2:len(s1)]
	fmt.Println(s1[:3]) // => [1 2 3], same as s1[0:3]
	fmt.Println(s1[:])  // => [1 2 3 4 5 6], same with s1[0:len(s1)]

	//slicing with append
	s1 = append(s1[:4], 100)
	fmt.Println(s1)

	s1 = append(s1[:4], 1200) // 100 will be overriden to 1200

	fmt.Println(s1)

}
