package main

import (
	"fmt"
	"os"

	config "github.com/nj-eka/gb_go_1/lesson8/config"
)

func main() {

	os.Setenv("cmd", os.Args[0])
	environs := config.EnvironsMap()
	fmt.Println(environs["cmd"])
	fmt.Println(config.GetEnval("cmd", ""))

}
