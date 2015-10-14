package main

import "fmt"

func main() {
	// sieve of Eratosthenes - let's find some primes!
	// I was going to prime the pump with the number 2 anyway,
	// so why not add a few more?

	primes := []int{2}

	// any number that is not divisible by one of our primes is a prime
	// we'll skip all of the even numbers, because we know that they're not prime.

	isPrime := true
	for i := 3; i < 1000000; i += 2 {
		for _, prime := range primes {
			if i%prime == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			primes = append(primes, i)
			// fmt.Println(i)
		}
		isPrime = true
	}

	for _, prime := range primes {
		fmt.Println(prime)
	}

}
