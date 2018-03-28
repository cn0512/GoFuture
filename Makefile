#make core

build:
	go install -v -x -a -buildmode=shared runtime sync/atomic #构建核心基本库
	go install -v -x -a -buildmode=shared -linkshared #构建GO动态库

quote:
	go build -v -x -linkshared ./Instance/quote/q.go

trade:
	go build -v -x -linkshared ./trade/trade.go