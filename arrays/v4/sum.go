package main

import(
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
func SumAll(numbersToSum...[]int) []int{
	numbersLength := len(numbersToSum)
	sums := make([]int, numbersLength)

	for i,numbers := range numbersToSum {
		sums[i] = Sum(numbers)
	}
	return sums
}

func main(){
	val := SumAll([]int{1,3}, []int{3,6})
	fmt.Println(val)  //forgot to use fmt package to print the result of invoking this function
	
}
//what is the standard convention for writing names of the functions: when should the names start with capital letters? 