package main

func Repeat(x string) string {

	var result string
	for i := 0; i < 5; i++ {
		result += x
	}
	return result
}

func main() {
	print(Repeat("a"))
}