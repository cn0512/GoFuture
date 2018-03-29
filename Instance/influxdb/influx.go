package main

/*
	sub topic,recv tick data and save to influx
*/

import (
	"fmt"
	"os"
	"os/signal"

	cons "github.com/cn0512/GoFuture/constant"
	mq "github.com/cn0512/GoServerFrame/Core/MQ/Redis"
)

type MsgCmd struct {
}

func (m *MsgCmd) Call(msg string) {
	fmt.Println("Call:", msg)
}

func main() {
	//sub mq
	sub, err_sub := mq.NewSub(cons.Topic_Quote_Tick, &MsgCmd{})
	if err_sub != nil {
		panic(err_sub)
	}
	defer sub.Shutdown()

	qt := make(chan os.Signal, 1)
	signal.Notify(qt, os.Interrupt, os.Kill)
	<-qt
}
