package main

import (
	q "tblive/src/ctp/quote"

	cons "tblive/src/constant"
)

func main() {

	//CTP Thread
	q.Start(cons.Topic_Quote_Tick)
}
