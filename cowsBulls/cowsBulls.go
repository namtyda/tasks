//Secret number: 4271
//Opponents try: 1234
//Answer: 1 bull and 2 cows
package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func main() {
	var secretNumber string
	go func() {
		secretNumber = genSecret()
	}()
	for {
		var input string
		fmt.Scan(&input)

		if uniqueInput(input) == false {
			fmt.Print("numbers must be uique\n")
			continue
		}
		if len(input) != 4 {
			fmt.Print("len must be 4\n")
			continue
		}
		if input == secretNumber {
			fmt.Print("You won\n")
			return
		}
		bulls := countBullsRecursive(secretNumber, input)
		fmt.Printf("bull:%d cows:%d\n", bulls, countCowsRecursive(secretNumber, input)-bulls)
		fmt.Printf("secret number %s\n", secretNumber)
	}
}

func genSecret() string {
	rand.Seed(time.Now().UnixNano())
	result := make([]string, 0, 4)
	result = append(result, strconv.Itoa(rand.Intn(10)))

cont:
	for len(result) != 4 {
		randInt := strconv.Itoa(rand.Intn(10))
		for _, v := range result {
			if randInt == v {
				fmt.Println(result)
				continue cont
			}
		}
		result = append(result, randInt)
		continue
	}
	return strings.Join(result, "")
}

func contains(input string, searchLiter string) bool {
	if len(input) == 0 {
		return false
	}
	if string(input[0]) == searchLiter {
		return true
	}
	return contains(input[1:], searchLiter)
}

func uniqueInput(input string) bool {
	for i := 0; i < len(input)-1; i++ {
		if contains(input[i+1:], input[i:i+1]) {
			return false
		}
	}
	return true
}

func countBulls(secretNumber string, inputNumber string) int {
	bulls := 0
	for i, v := range secretNumber {
		if v == rune(inputNumber[i]) {
			bulls++
		}
	}
	return bulls
}

//countBulls("4216","4836") == ...countBulls("216","836") + 1
//2 == ...1 + 1

//countBulls("2216","4836") == countBulls("216","836")
//Recursive variant countBulls
func countBullsRecursive(secretNumber string, inputNumber string) int {
	if len(secretNumber) == 0 && len(inputNumber) == 0 {
		return 0
	}
	if secretNumber[0] == inputNumber[0] {
		return countBullsRecursive(secretNumber[1:], inputNumber[1:]) + 1
	}
	return countBullsRecursive(secretNumber[1:], inputNumber[1:])
}

// 4712  4892
//Recursive variant countCows
func countCowsRecursive(secretNumber string, inputNumber string) int {
	if len(inputNumber) == 0 {
		return 0
	}
	if contains(secretNumber, inputNumber[:1]) {
		return countCowsRecursive(secretNumber, inputNumber[1:]) + 1
	}
	return countCowsRecursive(secretNumber, inputNumber[1:])
}

func countCows(secretNumber string, inputNumber string) int {
	cows := 0
	for iSecret, vSecret := range secretNumber {
		for iInput, vInput := range inputNumber {
			if vSecret == vInput && iSecret != iInput {
				cows++
			}
		}
	}
	return cows
}
