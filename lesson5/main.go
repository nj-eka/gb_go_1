package main

import (
	"fmt"
	"log"
	"runtime"
	"time"
)

func FibonacciFun() func() int {
	var first, second int = 0, 1
	return func() int {
		res := first
		first, second = second, first+second
		return res
	}
}

func FibonacciDefered() func() int {
	first, second := 1, 1
	return func() int {
		defer func() {
			first = first + second
			first, second = second, first
		}()
		return first
	}
}

func FibonacciCached(n int, cache map[int]int) int {
	defer timeTrack(time.Now(), "")
	if value, found := cache[n]; found {
		return value
	}
	cache[n] = FibonacciCached(n-1, cache) + FibonacciCached(n-2, cache)
	return cache[n]
}

// https://blog.stathat.com/2012/10/10/time_any_function_in_go.html
func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	if len(name) == 0 {
		name = funcName(2)
	}
	log.Printf("function %s took %v to complete.", name, elapsed)
}

func funcName(callerNum int) string {
	pc, _, _, _ := runtime.Caller(callerNum)
	return runtime.FuncForPC(pc).Name()
}

func main() {
	var n int
	for {
		fmt.Println("Enter number (int) for Fibonacci: ")
		if c, e := fmt.Scanf("%d", &n); e != nil && c != 1 {
			fmt.Printf("Invalid input %v - try again.\n", n)
			continue
		}
		break
	}

	{
		defer timeTrack(time.Now(), "FibonacciFun")
		f := FibonacciFun()
		for i := 0; i < int(n)-1; i++ {
			f()
		}
		fmt.Println("Fibonacci number = ", f())
	}

	{
		defer timeTrack(time.Now(), "FibonacciDefered")
		f := FibonacciDefered()
		for i := 0; i < int(n)-1; i++ {
			f()
		}
		fmt.Println("Fibonacci number = ", f())
	}

	{
		defer timeTrack(time.Now(), "FibonacciCached")
		fmt.Println("Число Фибоначчи:", FibonacciCached(n, map[int]int{1: 1, 2: 1}))
	}

}
