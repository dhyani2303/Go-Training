package main

import "fmt"

func main() {

	nums := []int{1, 2, 3, 4, 5}

	nums = append(nums, 10)

	fmt.Printf("%#v\n", nums)

	fmt.Println(nums)

	nums = append(nums, 20, 30, 40)

	nums2 := []int{1, 2, 3}

	nums = append(nums, nums2...) //ellipses operator is used to copy all the elements of one slice to another
	fmt.Printf("%#v\n", nums)

	//copying a slice to another.
	//if both the slices dont have the same number of elements then minimum length from both is taken

	src := []int{3, 4, 5}
	//dst := make([]int, len(src))
	dst := make([]int, 2) // [3,4] is copied to destination

	nn := copy(dst, src) // it returns the number of elements copied

	fmt.Println("Source :", src, "Dest:", dst, nn)

}
