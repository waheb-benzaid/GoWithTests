package main

func Sum(numbers []int) int {
	sum := 0

	for index, number := range numbers {
		print(index)
		sum += number
	}
	return sum
}

func SumAll(numberToSum ...[]int) []int {

	// lenghtOfNumber := len(numberToSum)
	// sumOfNumbers := make([]int, lenghtOfNumber)

	// for index, number := range numberToSum {
	// 	sumOfNumbers[index] = Sum(number)
	// }

	// return sumOfNumbers

	var sums []int

	for _, number := range numberToSum {
		sums = append(sums, Sum(number))
	}

	return sums
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	print(Sum(numbers))
}