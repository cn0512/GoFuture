package ctp

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/cn0512/GoFuture"
	"gopkg.in/yaml.v2"
)

/*
const (
	Broker_id    = "9999"
	Investor_id  = "025458"
	Pass_word    = "123456"
	Market_front = "tcp://180.168.146.187:10011"
	Trade_front  = "tcp://180.168.146.187:10001"
)
*/
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

type YamlCfg struct {
	BrokerID    string   `yaml:"BrokerID"`
	InvestorID  string   `yaml:"InvestorID"`
	Password    string   `yaml:"Password"`
	MdFront     string   `yaml:"MdFront"`
	TraderFront string   `yaml:"TraderFront"`
	Ins         []string `yaml:"Ins"`
}

type QuoteTick struct {
	TradingDay         string `json:"TradingDay"`
	InstrumentID       string `json:"InstrumentID"`
	ExchangeID         string `json:"ExchangeID"`
	ExchangeInstID     string `json:"ExchangeInstID"`
	LastPrice          string `json:"LastPrice"`
	PreSettlementPrice string `json:"PreSettlementPrice"`
	PreClosePrice      string `json:"PreClosePrice"`
	PreOpenInterest    string `json:"PreOpenInterest"`
	OpenPrice          string `json:"OpenPrice"`
	HighestPrice       string `json:"HighestPrice"`
	LowestPrice        string `json:"LowestPrice"`
	Turnover           string `json:"Turnover"`
	OpenInterest       string `json:"OpenInterest"`
	ClosePrice         string `json:"ClosePrice"`
	SettlementPrice    string `json:"SettlementPrice"`
	UpperLimitPrice    string `json:"UpperLimitPrice"`
	LowerLimitPrice    string `json:"LowerLimitPrice"`
	PreDelta           string `json:"PreDelta"`
	CurrDelta          string `json:"CurrDelta"`
	UpdateTime         string `json:"UpdateTime"`
	UpdateMillisec     string `json:"UpdateMillisec"`

	BidPrice1  string `json:"BidPrice1"`
	BidVolume1 string `json:"BidVolume1"`
	AskPrice1  string `json:"AskPrice1"`
	AskVolume1 string `json:"AskVolume1"`

	BidPrice2  string `json:"BidPrice2"`
	BidVolume2 string `json:"BidVolume2"`
	AskPrice2  string `json:"AskPrice2"`
	AskVolume2 string `json:"AskVolume2"`

	BidPrice3  string `json:"BidPrice3"`
	BidVolume3 string `json:"BidVolume3"`
	AskPrice3  string `json:"AskPrice3"`
	AskVolume3 string `json:"AskVolume3"`

	BidPrice4  string `json:"BidPrice4"`
	BidVolume4 string `json:"BidVolume4"`
	AskPrice4  string `json:"AskPrice4"`
	AskVolume4 string `json:"AskVolume4"`

	BidPrice5  string `json:"BidPrice5"`
	BidVolume5 string `json:"BidVolume5"`
	AskPrice5  string `json:"AskPrice5"`
	AskVolume5 string `json:"AskVolume5"`

	AveragePrice string `json:"AveragePrice"`
	ActionDay    string `json:"ActionDay"`
}

func (g *CtpCfg) GetMdRequestID() int {
	g.MdRequestID += 1
	return g.MdRequestID
}

func (g *CtpCfg) GetTraderRequestID() int {
	g.TraderRequestID += 1
	return g.TraderRequestID
}

func Parse(file string, cfg interface{}) {
	buf, err := ioutil.ReadFile(file)
	if err != nil {
		log.Panic(err)
	}
	//fmt.Println(buf)
	err = yaml.Unmarshal(buf, cfg)
	if err != nil {
		log.Panic(err)
	}
}

func (q *QuoteTick) Create(data GoFuture.CThostFtdcDepthMarketDataField) {
	q.TradingDay = ""
}

func (q *QuoteTick) Encode() []byte {
	buf, err := json.Marshal(*q)
	if err != nil {
		log.Fatalln(err)
	}
	return buf
}

func (q *QuoteTick) Decode(buf []byte) error {
	err := json.Unmarshal(buf, q)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}
