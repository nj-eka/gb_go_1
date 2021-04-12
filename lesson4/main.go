package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

//https://en.wikipedia.org/wiki/Quicksort
func partition(a []int64, lo, hi int) int {
	p := a[hi]
	for j := lo; j < hi; j++ {
		if a[j] < p {
			a[j], a[lo] = a[lo], a[j]
			lo++
		}
	}

	a[lo], a[hi] = a[hi], a[lo]
	return lo
}

func quickSort(a []int64, lo, hi int) {
	if lo > hi {
		return
	}

	p := partition(a, lo, hi)
	quickSort(a, lo, p-1)
	quickSort(a, p+1, hi)
}

func main() {
	inputNums := []int64{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		num, err := strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			break
		}
		inputNums = append(inputNums, num)
	}
	quickSort(inputNums, 0, len(inputNums)-1)
	for _, e := range inputNums {
		fmt.Println(e)
	}
}
