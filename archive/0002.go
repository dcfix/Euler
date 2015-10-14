package main

import "fmt"

func main() {

	first := 1
	second := 1
	sum := 0
	for second < 4000000 {
		second = first + second
		first = second - first
		if second%2 == 0 {
			sum += second
			fmt.Println(second)
		}
	}

	fmt.Println(sum)

}
