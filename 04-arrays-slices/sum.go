package main

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	lengthofNumbers := len(numbersToSum)
	sums := make([]int, lengthofNumbers) // make allows you to create a slice with starting capacity of lengthofNumbers
	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
	}
	return sums
}
