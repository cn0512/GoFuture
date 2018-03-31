#make core

build:
	go install -v -x -a -buildmode=shared runtime sync/atomic #构建核心基本库
	go install -v -x -a -buildmode=shared -linkshared #构建GO动态库

quoteserver:
	go build -linkshared -o ./Instance/bin/quote.exe ./Instance/quote/q.go
tradeserver:
	go build -linkshared -o ./Instance/bin/trade.exe ./Instance/trade/trade.go
influxserver:
	go build -o ./Instance/bin/influx.exe ./Instance/influxdb/*.go
chatserver:
	go build -o ./Instance/bin/chat.exe ./Instance/chat/Server/*.go
robot:
	go build -o ./Instance/bin/robot.exe ./Instance/chat/robot/*.go