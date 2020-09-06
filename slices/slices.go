package slices

// Sum takes a slice of integers and adds them up to a single result.
func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}

	return sum
}

// SumAll takes any number of slices of integers and adds each slice up to return a slice of all the sums.
func SumAll(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return sums
}

// SumAllTails takes any number of integer slices and sums all the elements of each slice excluding the first one
// and returns a slice with a total for each slice passed in.
func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
			continue
		}
		tail := numbers[1:]
		sums = append(sums, Sum(tail))
	}

	return sums
}
