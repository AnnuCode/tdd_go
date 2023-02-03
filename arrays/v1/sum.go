package main

// func Sum(numbers [5]int) int{
// 	result := 0

// 	for i:=0;i<5;i++ {
// 		result += numbers[i]
// 	}
// 	return result
// }
 
//using range instead of for loop

func Sum(numbers []int) int{
	result := 0
	for _,number := range numbers{
		result += number
	}
	return result
}
