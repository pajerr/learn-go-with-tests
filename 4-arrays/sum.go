package main

import "fmt"

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		fmt.Println("numbers:", numbers)
		sums = append(sums, Sum(numbers))
		fmt.Println("sums after:", sums)
	}

	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		fmt.Println("numbers:", numbers)
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			fmt.Println("tail:", tail)
			sums = append(sums, Sum(tail))
			fmt.Println("sums after:", sums)
		}
	}

	return sums
}

/*
// variadic function
func SumAll(numbersToSum ...[]int) []int {
	lengthOfNumbers := len(numbersToSum)
	//make allows you to create a slice with a starting capacity of the len of the numbersToSum we need to work through
	sums := make([]int, lengthOfNumbers)

	// range over the slice and create new slices that have sum of each slice to sums slice
	for i, numbers := range numbersToSum {
		//sums[i] determines which sliece we loop from slice of slices
		sums[i] = Sum(numbers)
	}

	return sums
}
*/
