package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"path"
	"strings"

	envconfig "github.com/kelseyhightower/envconfig"
	yaml "gopkg.in/yaml.v2"
)

type Config struct {
	Server struct {
		Host string `envconfig:"SERVER_HOST", yaml:"host", json:"host"`
		Port string `envconfig:"SERVER_PORT", yaml:"port", json:"port"`
		Path string `envconfig:"SERVER_PATH", yaml:"path", json:"path"`
	} `yaml:"server" json:"server"`
}

func processError(err error) {
	if err != nil {
		log.Fatal(err)
		// fmt.Println(err)
		// os.Exit(2)
	}
}

func readFile(cfg *Config, configPath string) {

	b, err := ioutil.ReadFile(configPath)
	processError(err)

	configType := strings.ToLower(path.Ext(configPath))[1:]

	switch configType {
	case "yaml", "yml":
		err = yaml.Unmarshal(b, &cfg)
	case "json", "jsn":
		err = json.Unmarshal(b, &cfg)
	default:
		log.Printf("Unexpected config file type %s", configPath)
	}
	processError(err)
}

func readEnv(cfg *Config) {
	err := envconfig.Process("", cfg)
	processError(err)
}

func GetConfig(configPath string) *Config {
	var cfg Config // set defaults
	readFile(&cfg, configPath)
	readEnv(&cfg)
	//	readCli(&cfg)
	return &cfg
}
