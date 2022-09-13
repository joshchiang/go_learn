package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Set the default values
	fmt.Println("Bucket Sort")
	data := []int{89, 34, 23, 78, 67, 100, 66, 29, 79, 55, 78, 88, 92, 96, 96, 23}

	bucketsort(data)
}
func bucketsort(data []int) {
	const max_scroe int = 100
	var dataSize = len(data)
	var bucket [max_scroe + 1]int
	var bucketRep [max_scroe + 1]int

	//桶子歸零
	for i := 0; i <= max_scroe; i++ {
		bucket[i] = 0
		bucketRep[i] = 0
		// fmt.Println(i)
	}
	// fmt.Println(len(bucket))
	// // 放入對應
	for i := 0; i < dataSize; i++ {
		if bucket[data[i]] > 0 {
			bucketRep[data[i]] = bucketRep[data[i]] + 1
			fmt.Println("Has Same: " + strconv.Itoa(bucket[data[i]]) + ":" + strconv.Itoa(bucketRep[data[i]]))
		} else {
			bucket[data[i]] = data[i]
			bucketRep[data[i]] = 1
		}
	}
	// print sort
	for i := 0; i <= max_scroe; i++ {
		//print has vaule
		if bucket[i] != 0 {
			fmt.Println("bucket: " + strconv.Itoa(bucket[i]))
			fmt.Println("bucket: " + strconv.Itoa(bucketRep[i]))

		}
	}
}
