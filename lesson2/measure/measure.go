package measure

import (
	"fmt"
	"math"
)

const maxInputAttempts int = 2

func promptFloat(msg string, attempts int) (float64, error) {
	var value float64
	var err error = nil
	for i := 0; i < attempts; i++ {
		fmt.Print(msg)
		n, err := fmt.Scanf("%f\n", &value)
		if err == nil && n == 1 {
			break
		}
	}
	return value, err
}

func promptInt(msg string, attempts int) (int, error) {
	var value int
	var err error = nil
	for i := 0; i < attempts; i++ {
		fmt.Print(msg)
		n, err := fmt.Scanf("%d\n", &value)
		if err == nil && n == 1 {
			break
		}
	}
	return value, err
}

func RectangleS() {
	a, err := promptFloat("Input value (float, > 0) of rectangle side [a]: ", maxInputAttempts)
	if err != nil {
		return
	}
	b, err := promptFloat("Input value (float, > 0) of rectangle side [b]: ", maxInputAttempts)
	if err != nil {
		return
	}
	fmt.Printf("Area = %f\n", a*b)
}

func CircleDR() {
	s, err := promptFloat("Input value (float, > 0) of circle area: ", maxInputAttempts)
	if err != nil {
		return
	}
	r := math.Sqrt(s / math.Pi)
	fmt.Printf("D = %f\nL = %f", 2 * r, 2 * r * math.Pi)
}

func NumberParts() {
	n, err := promptInt("Input number [0, 999]: ", maxInputAttempts)
	if err != nil {
		return
	}
	fmt.Printf("%d = %d + %d + %d", n, n / 100, (n / 10) % 10, n % 10
}