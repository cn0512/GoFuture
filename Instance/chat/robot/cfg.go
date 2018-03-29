package main

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type ServerCfg struct {
	Addr string `yaml:"addr"`
}

func Parse(file string, cfg interface{}) {
	buf, err := ioutil.ReadFile(file)
	if err != nil {
		log.Panic(err)
	}
	//fmt.Println(buf)
	err = yaml.Unmarshal(buf, cfg)
	if err != nil {
		log.Panic(err)
	}
}
