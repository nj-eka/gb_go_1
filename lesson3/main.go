package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	calc "github.com/nj-eka/gb_go_1/lesson3/calc"
	prime "github.com/nj-eka/gb_go_1/lesson3/prime"
)

func main() {
	fmt.Println("Select: [1, 2]")
	fmt.Println("1 - simple calculator")
	fmt.Println("2 - primes")

	var choice string
	fmt.Scan(&choice)
	switch choice {
	case "1":
		for {
			reader := bufio.NewReader(os.Stdin)
			fmt.Print("calc: ")
			input, _ := reader.ReadString('\n')
			lexer := calc.Lexer(input)
			interpreter := calc.Interpreter(&lexer)
			result, err := interpreter.Result()
			if err == nil {
				fmt.Printf("res: %s\n", strconv.FormatFloat(result, 'f', -1, 64))
			} else {
				fmt.Printf("%s\n", err)
			}
		}
	case "2":
		fmt.Println("Input int number up to which you want to find primes: ")
		var n int
		if _, err := fmt.Scanf("%d", &n); err == nil {
			primes := prime.FindPrimesByBruteForce(n)
			for _, p := range primes {
				fmt.Println(p)
			}
		}
	default:
		fmt.Println("Unknown operation")
	}
}
