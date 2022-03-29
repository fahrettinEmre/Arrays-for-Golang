package main

import (
	"fmt"
	"time"
)

func main() {
	// Arrays
	fmt.Printf("Today : %s\n\n", time.Now())

	var points = []float32{10.45, -30.345, 55.90, 60.0123}
	var names = [4]string{"sarlok", "sumi", "varlord", "khan"}
	var numbers [7]int
	var matrix = [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}

	for i := 0; i < len(points); i++ {
		fmt.Printf("%d is %f\n", i, points[i])
	}

	for i := 0; i < len(names); i++ {
		fmt.Printf("%s\n", names[i])
	}

	var j int
	for j = 0; j < len(numbers); j++ {
		numbers[j] = j * j
		fmt.Printf("%d\t", numbers[j])
	}

	fmt.Printf("\nSum = %f\n", sum(points))

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("%d\t", matrix[i][j])
		}
		fmt.Println("")
	}
}

func sum(nmbrs []float32) float32 {
	var toplam float32 = 0
	for i := 0; i < len(nmbrs); i++ {
		toplam += nmbrs[i]
	}
	return toplam
}
