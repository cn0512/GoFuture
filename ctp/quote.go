package ctp

import (
	"encoding/json"
	//"flag"
	"log"
	//"os"
	ps "github.com/aalness/go-redis-pubsub"
	"github.com/cn0512/GoFuture"
	mq "github.com/cn0512/GoServerFrame/Core/MQ/Redis"
	Utils "github.com/cn0512/GoServerFrame/Core/Utils"
)

var pub ps.Publisher
var topic string

func NewDirectorCThostFtdcMdSpi(v interface{}) GoFuture.CThostFtdcMdSpi {

	return GoFuture.NewDirectorCThostFtdcMdSpi(v)
}

type GoCThostFtdcMdSpi struct {
	Client CtpCfg
}

func (p *GoCThostFtdcMdSpi) OnRspError(pRspInfo GoFuture.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
	log.Println("GoCThostFtdcMdSpi.OnRspError.")
	p.IsErrorRspInfo(pRspInfo)
}

func (p *GoCThostFtdcMdSpi) OnFrontDisconnected(nReason int) {
	log.Printf("GoCThostFtdcMdSpi.OnFrontDisconnected: %#v\n", nReason)
}

func (p *GoCThostFtdcMdSpi) OnHeartBeatWarning(nTimeLapse int) {
	log.Printf("GoCThostFtdcMdSpi.OnHeartBeatWarning: %v", nTimeLapse)
}

func (p *GoCThostFtdcMdSpi) OnFrontConnected() {
	log.Println("GoCThostFtdcMdSpi.OnFrontConnected.")
	p.ReqUserLogin()
}

func (p *GoCThostFtdcMdSpi) ReqUserLogin() {
	log.Println("GoCThostFtdcMdSpi.ReqUserLogin.")

	req := GoFuture.NewCThostFtdcReqUserLoginField()
	req.SetBrokerID(p.Client.BrokerID)
	req.SetUserID(p.Client.InvestorID)
	req.SetPassword(p.Client.Password)

	iResult := p.Client.MdApi.ReqUserLogin(req, p.Client.GetMdRequestID())

	if iResult != 0 {
		log.Println("发送用户登录请求: 失败.")
	} else {
		log.Println("发送用户登录请求: 成功.")
	}
}

func (p *GoCThostFtdcMdSpi) IsErrorRspInfo(pRspInfo GoFuture.CThostFtdcRspInfoField) bool {
	// 如果ErrorID != 0, 说明收到了错误的响应
	bResult := (pRspInfo.GetErrorID() != 0)
	if bResult {
		log.Printf("ErrorID=%v ErrorMsg=%v\n", pRspInfo.GetErrorID(), pRspInfo.GetErrorMsg())
	}
	return bResult
}

func (p *GoCThostFtdcMdSpi) OnRspUserLogin(pRspUserLogin GoFuture.CThostFtdcRspUserLoginField, pRspInfo GoFuture.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {

	if bIsLast && !p.IsErrorRspInfo(pRspInfo) {

		//log.Printf("获取当前版本信息: %#v\n", GoFuture.CThostFtdcTraderApiGetApiVersion())
		log.Printf("获取当前交易日期: %#v\n", p.Client.MdApi.GetTradingDay())
		log.Printf("获取用户登录信息: %#v %#v %#v\n", pRspUserLogin.GetLoginTime(), pRspUserLogin.GetSystemName(), pRspUserLogin.GetSessionID())

		//ppInstrumentID := []string{"cu1610", "cu1611", "cu1612", "cu1701", "cu1702", "cu1703", "cu1704", "cu1705", "cu1706"}
		ppInstrumentID := []string{"rb1805", "au1805"}

		p.SubscribeMarketData(ppInstrumentID)
		p.SubscribeForQuoteRsp(ppInstrumentID)
	}
}

func (p *GoCThostFtdcMdSpi) SubscribeMarketData(name []string) {

	iResult := p.Client.MdApi.SubscribeMarketData(name)

	if iResult != 0 {
		log.Println("发送行情订阅请求: 失败.")
	} else {
		log.Println("发送行情订阅请求: 成功.")
	}
}

func (p *GoCThostFtdcMdSpi) SubscribeForQuoteRsp(name []string) {

	iResult := p.Client.MdApi.SubscribeForQuoteRsp(name)

	if iResult != 0 {
		log.Println("发送询价订阅请求: 失败.")
	} else {
		log.Println("发送询价订阅请求: 成功.")
	}
}

func (p *GoCThostFtdcMdSpi) OnRspSubMarketData(pSpecificInstrument GoFuture.CThostFtdcSpecificInstrumentField, pRspInfo GoFuture.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
	log.Printf("GoCThostFtdcMdSpi.OnRspSubMarketData: %#v %#v %#v\n", pSpecificInstrument.GetInstrumentID(), nRequestID, bIsLast)
	p.IsErrorRspInfo(pRspInfo)
}

func (p *GoCThostFtdcMdSpi) OnRspSubForQuoteRsp(pSpecificInstrument GoFuture.CThostFtdcSpecificInstrumentField, pRspInfo GoFuture.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
	log.Printf("GoCThostFtdcMdSpi.OnRspSubForQuoteRsp: %#v %#v %#v\n", pSpecificInstrument.GetInstrumentID(), nRequestID, bIsLast)
	p.IsErrorRspInfo(pRspInfo)
}

func (p *GoCThostFtdcMdSpi) OnRspUnSubMarketData(pSpecificInstrument GoFuture.CThostFtdcSpecificInstrumentField, pRspInfo GoFuture.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
	log.Printf("GoCThostFtdcMdSpi.OnRspUnSubMarketData: %#v %#v %#v\n", pSpecificInstrument.GetInstrumentID(), nRequestID, bIsLast)
	p.IsErrorRspInfo(pRspInfo)
}

func (p *GoCThostFtdcMdSpi) OnRspUnSubForQuoteRsp(pSpecificInstrument GoFuture.CThostFtdcSpecificInstrumentField, pRspInfo GoFuture.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
	log.Printf("GoCThostFtdcMdSpi.OnRspUnSubForQuoteRsp: %#v %#v %#v\n", pSpecificInstrument.GetInstrumentID(), nRequestID, bIsLast)
	p.IsErrorRspInfo(pRspInfo)
}

func (p *GoCThostFtdcMdSpi) OnRtnDepthMarketData(pDepthMarketData GoFuture.CThostFtdcDepthMarketDataField) {

	log.Println("GoCThostFtdcMdSpi.OnRtnDepthMarketData: ", pDepthMarketData.GetTradingDay(),
		pDepthMarketData.GetInstrumentID(),
		pDepthMarketData.GetExchangeID(),
		pDepthMarketData.GetExchangeInstID(),
		pDepthMarketData.GetLastPrice(),
		pDepthMarketData.GetPreSettlementPrice(),
		pDepthMarketData.GetPreClosePrice(),
		pDepthMarketData.GetPreOpenInterest(),
		pDepthMarketData.GetOpenPrice(),
		pDepthMarketData.GetHighestPrice(),
		pDepthMarketData.GetLowestPrice(),
		pDepthMarketData.GetVolume(),
		pDepthMarketData.GetTurnover(),
		pDepthMarketData.GetOpenInterest())

	//log.Printf("GoCThostFtdcMdSpi.OnRtnDepthMarketData: %+v\n", &pDepthMarketData)
	buf, _ := json.Marshal(pDepthMarketData)
	pub.Publish(topic, buf)
}

func (p *GoCThostFtdcMdSpi) OnRtnForQuoteRsp(pForQuoteRsp GoFuture.CThostFtdcForQuoteRspField) {
	log.Printf("GoCThostFtdcMdSpi.OnRtnForQuoteRsp: %#v\n", pForQuoteRsp)
}

func init() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	log.SetPrefix("CTP: ")
}

func Start(tp string) {

	//[1]MQ
	pub = mq.NewPub()
	defer pub.Shutdown()
	Utils.Logout("Pub`s init")
	topic = tp

	//[2]
	CTP = CtpCfg{
		BrokerID:   Broker_id,
		InvestorID: Investor_id,
		Password:   Pass_word,

		MdFront: Market_front,
		MdApi:   GoFuture.CThostFtdcMdApiCreateFtdcMdApi(),

		TraderFront: Trade_front,
		TraderApi:   GoFuture.CThostFtdcTraderApiCreateFtdcTraderApi(),

		MdRequestID:     0,
		TraderRequestID: 0,
	}

	log.Printf("客户端配置: %+#v\n", CTP)

	pMdSpi := GoFuture.NewDirectorCThostFtdcMdSpi(&GoCThostFtdcMdSpi{Client: CTP})

	CTP.MdApi.RegisterSpi(pMdSpi)
	CTP.MdApi.RegisterFront(CTP.MdFront)
	CTP.MdApi.Init()

	CTP.MdApi.Join()
	CTP.MdApi.Release()
}
