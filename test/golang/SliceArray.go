package main

import "fmt"

func changeSlicePtr(slice *[]int)  {
	slice_temp := append(*slice, 1,2,3)
	fmt.Println("In changeSlice : ",slice_temp)
	slice = &slice_temp
}

func changeSlice(slice *[]int)  {
	(*slice)[0] = 10
}

func main() {
	// array test
	var array = []int{1,2,3,4}
	fmt.Println(array)

	// slice test
	fmt.Println("======== slice test ========")
	var slice []int
	fmt.Println("Original slice : ",slice)
	slice = append(slice, 1);
	fmt.Println("Append slice : " ,slice)

	// clone a pointer of slice, so if change the param will not change slice,just change the clone pointer of slice
	changeSlicePtr(&slice)
	fmt.Println("After changeSlicePtr : ",slice)

	changeSlice(&slice)
	fmt.Println("After changeSlice : ",slice)


	fmt.Println("======== slice references array ========")
	slice1 := array[:3]
	fmt.Println("Before array changes ；", slice1)
	array[0] = 10
	fmt.Println("After array changes : ",slice1)


}
