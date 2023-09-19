package main

import (
	"fmt"
)

func main() {
	var grades [5]int
	fmt.Println(grades)

	var fruites [5]string
	fmt.Println(fruites)

	var a [3]int = [3]int{10, 20, 30}
	b := [3]int{10, 20, 30}
	fmt.Println(a)
	fmt.Println(b)

	c := [...]int{10, 20, 30, 40} //implicitly calculate length of array
	fmt.Println(c)
	c[3] = 100
	fmt.Println(c)
	//len() length of arr
	fmt.Println(len(c))
	//index
	fmt.Println(c[2])

	//looping in arr using range key
	for i, element := range c {
		fmt.Println(i, "=>", element)
	}

	//multi dimentional arr
	arr := [3][2]int{{1, 2}, {2, 3}, {3, 4}}
	fmt.Println(arr[1][1])

	//slice
	sl := []int{10, 20, 30}
	fmt.Println(sl)

	slice_1 := c[0:2]
	fmt.Println(slice_1)

	sub_slice := slice_1[0:1]
	fmt.Println(sub_slice)

	//make func to create empty silce
	slice_m := make([]int, 5, 8)
	fmt.Println(slice_m)
	fmt.Println(len(slice_m))
	fmt.Println((cap(slice_m)))

	//slice and index numbers
	arr_1 := [5]int{10, 20, 30, 40, 50}
	slice_2 := arr_1[:3]
	fmt.Println(arr_1)
	fmt.Println(slice_2)
	slice_2[1] = 1000
	fmt.Println("After Modification")
	fmt.Println(arr_1)
	fmt.Println(slice_2)

	//appending to a slice
	fmt.Println(len(slice_2))
	fmt.Println((cap(slice_2)))
	slice_2 = append(slice_2, 50, 60, 100)
	fmt.Println(slice_2)
	fmt.Println(len(slice_2))
	fmt.Println((cap(slice_2)))

	slice_2 = append(slice_2, slice_1...)
	fmt.Println(slice_2)
	fmt.Println(len(slice_2))
	fmt.Println((cap(slice_2)))

	//delete ele in slice
	//copy from slice
	dest_slice := make([]int, 10)
	num := copy(dest_slice, slice_2)
	fmt.Println(dest_slice)
	fmt.Println("number of ele copied: ", num)

	arr5 := []int{-1, -2}
	for _, value := range arr5 {
		fmt.Println(value)
	}

	/*
				arr := [5]int{10,20,90,70,60}
				slice := arr[:3]
				slice[2] = 900
				fmt.Println(arr)
				fmt.Println(slice)

				arr := [5]int{10, 20, 90, 70, 60}
		        slice := arr[:3]
		        fmt.Println(cap(slice))

		        slice_2 := make([]int, 5, 20)
		        new_slice := append(slice, slice_2...)
		        fmt.Println(cap(new_slice))

	*/

}
