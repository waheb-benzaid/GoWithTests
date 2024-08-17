package main

func Sum(numbers []int) int {
	sum := 0

	// for i := 0; i <= len(numbers); i++ {
	// 	sum += i
	// }

	// lets intreduce range!

	for index, number := range numbers {
		print(index)
		sum += number
	}
	return sum
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	print(Sum(numbers))
}