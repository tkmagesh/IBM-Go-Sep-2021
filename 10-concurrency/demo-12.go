package main

import "fmt"

/*
divide & conquer
	break the data in to two equal sets
	sum the data using 2 goroutintes
	get the sum results from the goroutines and print the final result
*/

func main() {
	data := []int{4, 1, 7, 3, 8, 5, 9, 3, 11, 15, 63, 87, 29, 40, 50, 55, 4, 1, 7, 3, 8, 5, 9, 3, 11, 15, 63, 87, 29, 40, 50, 55, 4, 1, 7, 3, 8, 5, 9, 3, 11, 15, 63, 87, 29, 40, 50, 55, 4, 1, 7, 3, 8, 5, 9, 3, 11, 15, 63, 87, 29, 40, 50, 55, 4, 1, 7, 3, 8, 5, 9, 3, 11, 15, 63, 87, 29, 40, 50, 55}

	// divide the data in to two equal sets
	dataSet1 := data[:len(data)/2]
	dataSet2 := data[len(data)/2:]
	resultCh := make(chan int)

	go func(dataSet1 []int) {
		resultCh <- sum(dataSet1...)
	}(dataSet1)

	go func(dataSet2 []int) {
		resultCh <- sum(dataSet2...)
	}(dataSet2)

	result := <-resultCh + <-resultCh
	fmt.Println("Result = ", result)

}

func sum(nos ...int) int {
	var sum int
	for _, no := range nos {
		sum += no
	}
	return sum
}
