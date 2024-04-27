package  main

func Sum(arr []int) int {
	sum := 0
	for _ , x := range arr{
		sum += x
	}
	return sum
}

// func SumAll(numbersToSum ...[]int) []int {
// 	lengthOfNumbers := len(numbersToSum)
// 	sums := make([]int, lengthOfNumbers)

// 	for i, numbers := range numbersToSum {
// 		sums[i] = Sum(numbers)
// 	}

// 	return sums
// }

func SumAll(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return sums
}

// func SumTail(arr []int) int {
// 	return arr[len(arr) - 1]
// }

// func SumAllTails(numbersToSum ...[]int) []int {
// 	var sums []int
// 	for _, numbers := range numbersToSum {
// 		sums = append(sums, SumTail(numbers))
// 	}
// 	return sums
// }

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