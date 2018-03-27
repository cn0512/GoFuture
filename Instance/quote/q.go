package main

import (
	q "github.com/cn0512/GoFuture/ctp/quote"

	cons "github.com/cn0512/GoFuture/constant"
)

func main() {

	//CTP Thread
	q.Start(cons.Topic_Quote_Tick)
}
