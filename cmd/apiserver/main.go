package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"
	apiserser "github.com/pyankovzhe/http-rest-api/internal/app/apiserver"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config file")
}
func main() {
	flag.Parse()
	config := apiserser.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}
	s := apiserser.New(config)

	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
