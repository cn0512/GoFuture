package main

import (
	ctp "github.com/cn0512/GoFuture/ctp"

	cons "github.com/cn0512/GoFuture/constant"
)

func main() {

	//CTP Thread
	ctp.Start(cons.Topic_Quote_Tick)
}
