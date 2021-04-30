package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/nj-eka/gb_go_1/lesson9/config"
)

func main() {
	os.Setenv("SERVER_HOST", "http://geekbrains.com")
	os.Setenv("SERVER_PORT", "80")
	os.Setenv("SERVER_PATH", "/index.html")

	configPath := flag.String("config", "config.yaml", "Config file path")
	flag.Parse()
	cfg := config.GetConfig(*configPath)
	fmt.Println(*cfg)
}
