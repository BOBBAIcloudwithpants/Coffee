package main

import (
	"flag"
	"io/ioutil"
	"log"

	coffee "github.com/XMatrixStudio/Coffee/App"
	yaml "gopkg.in/yaml.v2"
)

func main() {
	// 加载配置文件
	configFile := flag.String("c", "config/config.yaml", "Where is your config file?")
	flag.Parse()
	data, err := ioutil.ReadFile(*configFile)
	if err != nil {
		log.Printf("Can't find the config file in %v", *configFile)
		return
	}
	log.Printf("Load the config file in %v", *configFile)
	conf := coffee.Config{}
	yaml.Unmarshal(data, &conf)
	coffee.RunServer(conf)
}
