#make core
go build -o ./bin/influx.exe ./influxdb/influx.go
go build -o ./bin/chat.exe ./chat/Server/
go build -o ./bin/robot.exe ./chat/robot/

pause