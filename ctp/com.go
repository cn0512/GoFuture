package ctp

import (
	"github.com/cn0512/GoFuture"
)

const (
	Broker_id    = "9999"
	Investor_id  = "025458"
	Pass_word    = "test123"
	Market_front = "tcp://180.168.146.187:10011"
	Trade_front  = "tcp://180.168.146.187:10001"
)

var CTP CtpCfg

type CtpCfg struct {
	BrokerID   string
	InvestorID string
	Password   string

	MdFront string
	MdApi   GoFuture.CThostFtdcMdApi

	TraderFront string
	TraderApi   GoFuture.CThostFtdcTraderApi

	MdRequestID     int
	TraderRequestID int
}

func (g *CtpCfg) GetMdRequestID() int {
	g.MdRequestID += 1
	return g.MdRequestID
}

func (g *CtpCfg) GetTraderRequestID() int {
	g.TraderRequestID += 1
	return g.TraderRequestID
}
