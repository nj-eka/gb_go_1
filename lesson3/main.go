package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/nj-eka/gb_go_1/lesson3/calclib"
)

func main() {
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("calc: ")
		input, _ := reader.ReadString('\n')
		lexer := calclib.Lexer(input)
		interpreter := calclib.Interpreter(&lexer)
		result, err := interpreter.Result()
		if err == nil {
			fmt.Printf("res: %s\n", strconv.FormatFloat(result, 'f', -1, 64))
		} else {
			fmt.Printf("%s\n", err)
		}
	}
}
