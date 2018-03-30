package main

/*
	sub topic,recv tick data and save to influx
*/

import (
	"fmt"
	"os"
	"os/signal"

	cons "github.com/cn0512/GoFuture/constant"
	influx "github.com/cn0512/GoServerFrame/Core/DB/InfluxDB"
	mq "github.com/cn0512/GoServerFrame/Core/MQ/Redis"
)

type MsgCmd struct {
}

func (m *MsgCmd) Call(msg string) {
	fmt.Println("Call:", msg)
}

var c *influx.InfluxClient

func main() {
	//sub mq
	sub, err_sub := mq.NewSub(cons.Topic_Quote_Tick, &MsgCmd{})
	if err_sub != nil {
		panic(err_sub)
	}
	defer sub.Shutdown()

	//influx
	//uri := fmt.Sprintf("http://%s:%d", "localhost", 8086)
	cfg := &ServerCfg{}
	Parse("./influx.yaml", cfg)
	c = influx.New(cfg.Addr)
	c.Ping()
	c.Query()

	qt := make(chan os.Signal, 1)
	signal.Notify(qt, os.Interrupt, os.Kill)
	<-qt
}
