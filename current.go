package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {

	content, err := ioutil.ReadFile("primes.txt")

	if err != nil {
		//Do something
	}
	primes := strings.Split(string(content), "\n")

	answer := 0

}
