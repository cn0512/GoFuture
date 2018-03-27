package main

import (
	"encoding/json"
	"log"
	"net/url"
	"time"

	cons "github.com/cn0512/GoFuture/constant"

	"github.com/gorilla/websocket"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

func Connect() *websocket.Conn {
	u := url.URL{Scheme: "ws", Host: cons.ChatSvrAddr, Path: "/ws"}
	log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	/*
		c.SetReadLimit(maxMessageSize)
		c.SetReadDeadline(time.Now().Add(pongWait))
		c.SetPongHandler(func(string) error { c.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	*/
	return c
}

type RobotMsg struct {
	Email string
	Nick  string
	Body  string
}

func MakeMsg(str string) string {
	msg := RobotMsg{
		Email: cons.Chat_Robot_Email,
		Nick:  cons.Chat_Robot_Name,
		Body:  str,
	}
	buf, _ := json.Marshal(msg)
	return string(buf)
}
