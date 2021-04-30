package main

import (
	"flag"
	"fmt"
	"net/url"
	"os"

	config "github.com/nj-eka/gb_go_1/lesson8/config"
)

func main() {

	// env
	fmt.Println(os.Args)
	// os.Setenv("cmd", os.Args[0])
	// environs := config.EnvironsMap()
	// fmt.Println(environs["cmd"])
	// fmt.Println(config.GetEnval("cmd", ""))
	os.Setenv("url", "http://localhost")
	// url_env := config.GetEnval("url", "")

	//flags https://golang.org/pkg/flag/
	fmt.Println(flag.Args()) //Arg(i)

	var u = &url.URL{}
	flag.Var(&config.URLValue{u}, "url", "URL to parse")
	flag.Parse()
	// fs := flag.NewFlagSet("ExampleValue", flag.ExitOnError)
	// fs.Var(&config.URLValue{u}, "url", "URL to parse")
	// fs.Parse([]string{"-url", "https://golang.org/pkg/flag/"})
	fmt.Printf("{scheme: %q, host: %q, path: %q}\n", u.Scheme, u.Host, u.Path)

}
