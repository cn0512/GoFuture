package main

/*
	redis/pub 是广播模式；负载均衡的MQ，如NSQ
*/

import (
	"log"
	"os"
	"os/signal"

	cons "github.com/cn0512/GoFuture/constant"
	mq "github.com/cn0512/GoServerFrame/Core/MQ/Redis"
)

type MsgCmd struct {
}

func (m *MsgCmd) Call(msg string) {
	log.Println("Call:", msg)
}

func main() {
	//订阅CURD消息
	sub, err_sub := mq.NewSub(cons.Topic_CURD, &MsgCmd{})
	if err_sub != nil {
		panic(err_sub)
	}
	defer sub.Shutdown()
	log.Println("subing...", cons.Topic_CURD)

	qt := make(chan os.Signal, 1)
	signal.Notify(qt, os.Interrupt, os.Kill)
	<-qt
}
