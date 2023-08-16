package main

import "fmt"


func main()  {
	arr1 := [3]int {1, 2, 3}

	fmt.Printf("arr1 长度 %d, 容量 %d", len(arr1), cap(arr1))

	arr2 := [3]int{1, 2}
	fmt.Printf("arr2 长度 %d, 容量 %d", len(arr2), cap(arr2))

	// 每个元素全部int类型的零值
	arr3 := [3]int{}
	for idx, val := range arr3 {
		println("idx ,val " , idx, val)
	}

	arr4 := make([]int, 3, 5)
	arr4 = append(arr4, 1)

	for _, val := range arr4 {
		println("arr4 val", val)
	}

	arr5 := make([]int, 5, 10)
	arr5 = append(arr5, 2)
	arr6 := arr5[2:4]
	for idx := range(arr6) {
		println(arr6[idx])
	}

	arr6[1] = 2;
	// arr5也被子切片影响了
	for idx, val := range(arr5) {
		println(arr5[idx], val)
	}
}