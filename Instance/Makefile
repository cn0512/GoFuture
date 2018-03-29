#make core
build:
	go install -v -x -a -buildmode=shared runtime sync/atomic #构建核心基本库
	go install -v -x -a -buildmode=shared -linkshared #构建GO动态库
quoteserver:
	go build -linkshared -o ./bin/quote.exe ./quote/q.go
trade:
	go build -linkshared -o ./bin/trade.exe ./trade/trade.go
influx:
	go build -o ./bin/influx.exe ./influxdb/influx.go
chatserver:
	go build -o ./bin/chat.exe ./chat/Server/*.go
robot:
	go build -o ./bin/robot.exe ./chat/robot/*.go
