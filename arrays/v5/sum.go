package main

import (
	"fmt"
)

// func Sum(numbers [5]int) int{
// 	result := 0

// 	for i:=0;i<5;i++ {
// 		result += numbers[i]
// 	}
// 	return result
// }

func Sum(numbers []int) int {
	result := 0
	for _, number := range numbers {
		result += number
	}
	return result
}
func SumAll(numbersToSum ...[]int) []int {
	numbersLength := len(numbersToSum)
	sums := make([]int, numbersLength)

	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
	}
	return sums
}
func TailSumAll(numbersToSum ...[]int) []int {
	var sums []int
	
	for _, numbers := range numbersToSum {
		if(len(numbers)==0){
		sums=[]int{0}
	}else{

		tail := numbers[1:] //numbers should not be an empty array when reaching this step
		sums = append(sums, Sum(tail))
	}
	}
	return sums
}

func main() {
	val := TailSumAll([]int{}, []int{3, 6, 7})
	fmt.Println(val)

}
