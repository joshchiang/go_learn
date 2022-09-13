package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Quick Sort")
	data := []int{89, 34, 23, 78, 67, 100, 66, 29, 79, 55, 78, 88, 92, 96, 96, 23}
	dataLen := len(data)

	left := 0
	right := dataLen - 1

	quick(data, right, left)

	for i := 0; i < dataLen; i++ {
		fmt.Println("test data:" + strconv.Itoa(data[i]))
	}
}

func quick(data []int, right int, left int) {
	for j := 0; j <= right; j++ {
		if data[left] < data[j] {
			// tmpI = data[i]
			tmpJ := data[j]
			tmpRige := data[right]

			data[right] = tmpJ
			data[j] = tmpRige
			right--

			// fmt.Println("Change :" + strconv.Itoa(j) + ":" + strconv.Itoa(tmpJ) + "- R :" + strconv.Itoa(right) + ":" + strconv.Itoa(tmpRige))
		}
		if j == right {
			tmpi := data[left]
			tmpRige := data[right]

			data[right] = tmpi
			data[j] = tmpRige
		}
	}
}
