package greeting

import (
	"rsc.io/quote"
)

func SayHello() string {
	return quote.Hello()
}

func SayGo() string {
	return quote.Go()
}
