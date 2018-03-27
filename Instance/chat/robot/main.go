package main

/*
	聊天机器人，负责转发策略触发信息
*/

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	cons "github.com/cn0512/GoFuture/constant"

	mq "github.com/cn0512/GoServerFrame/Core/MQ/Redis"
	"github.com/gorilla/websocket"
)

var c *websocket.Conn
var cmdChan chan string

type MsgCmd struct {
}

func (m *MsgCmd) Call(msg string) {
	fmt.Println("Call:", msg)
	cmdChan <- msg
}

func main() {

	log.Println("robot is init...")

	//[1]链接到ChatSvr
	cmdChan = make(chan string, 1)
	c = Connect()
	go func() {
		for {
			select {
			case t := <-cmdChan:
				c.SetWriteDeadline(time.Now().Add(writeWait))
				err := c.WriteMessage(websocket.TextMessage, []byte(t))
				if err != nil {
					log.Fatal("write:", err)
				}
			}
		}

	}()

	//[2]订阅策略消息
	sub, err_sub := mq.NewSub(cons.Topic_ChatSvr_StrategyCmd, &MsgCmd{})
	if err_sub != nil {
		panic(err_sub)
	}
	defer sub.Shutdown()
	log.Println("subing...", cons.Topic_ChatSvr_StrategyCmd)

	//[3]
	log.Println("robot is runing...")

	qt := make(chan os.Signal, 1)
	signal.Notify(qt, os.Interrupt, os.Kill)
	<-qt
}
