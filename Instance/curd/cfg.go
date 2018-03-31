package main

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

/*
topic_pub: "tp_remotelog"
topic_sub: "nil"
pub_msg_list: "get_log_list"
pub_msg_detail: "get_log_detail"
sub_channel: "log.remote.dhms"
http: "localhost:8081"
nsq_tcp: "localhost:4150"
nsq_http: "localhost:4151"
*/

type SvcCfg struct {
	Topic_pub      string `yaml:"topic_pub"`
	Topic_sub      string `yaml:"topic_sub"`
	Pub_msg_list   string `yaml:"pub_msg_list"`
	Pub_msg_detail string `yaml:"pub_msg_detail"`
	Http           string `yaml:"http"`
	Nsq_tcp        string `yaml:"nsq_tcp"`
	Nsq_http       string `yaml:"nsq_http"`
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
