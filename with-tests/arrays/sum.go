package main

func Sum(numbers []int) int {
	var sum int

	for _, number := range numbers {
		sum += number
	}

	return sum
}

// SumAll takes a variable number of slices of integers and returns a slice containing
// the sum of each provided slice. Each element in the returned slice corresponds to
// the sum of the respective input slice.
//
// Example:
//
//	SumAll([]int{1,2}, []int{0,9}) // returns []int{3, 9}
func SumAll(numbersToSum ...[]int) []int {
	var sums []int

	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int

	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}

	return sums
}
