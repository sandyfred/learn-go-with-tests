package main

func Sum(numbers []int) (sum int) {
	for _, number := range numbers {
		sum += number
	}
	return
}

func SumAll(numbersToSum ...[]int) []int {
	var sum []int

	for _, numbers := range numbersToSum {
		sum = append(sum, Sum(numbers))
	}

	return sum
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sum []int

	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sum = append(sum, 0)
		} else {
			tail := numbers[1:]
			sum = append(sum, Sum(tail))
		}
	}

	return sum
}
