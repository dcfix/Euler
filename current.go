package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {

	primes, err := getPrimes("primes.txt")
	if err != nil {
		fmt.Println("Error loading pre calculate prime numbers")
	}
	fmt.Println("got my primes")
	fmt.Println(primes[2:5])

	target := 600851475143
	highest := 0
	square := int(math.Sqrt(float64(target)))

	for _, p := range primes {
		if p < square {
			if target%p == 0 {
				highest = p
			}

		}

	}
	fmt.Println(highest)
}

func getPrimes(path string) ([]int, error) {

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	var myPrimes []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		val, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, err
		}
		myPrimes = append(myPrimes, val)
	}
	return myPrimes, scanner.Err()
}
