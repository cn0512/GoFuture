package ctp

import (
	"flag"
	"log"
	"os"
	ctp "tblive/src/ctp/base"
	"time"
)

var (
	broker_id    = flag.String("BrokerID", "9999", "经纪公司编号,SimNow BrokerID统一为：9999")
	investor_id  = flag.String("InvestorID", "<InvestorID>", "交易用户代码")
	pass_word    = flag.String("Password", "<Password>", "交易用户密码")
	market_front = flag.String("MarketFront", "tcp://180.168.146.187:10031", "行情前置,SimNow的测试环境: tcp://180.168.146.187:10031")
	trade_front  = flag.String("TradeFront", "tcp://180.168.146.187:10030", "交易前置,SimNow的测试环境: tcp://180.168.146.187:10030")
)

var CTP GoCTPClient

type GoCTPClient struct {
	BrokerID   string
	InvestorID string
	Password   string

	MdFront string
	MdApi   ctp.CThostFtdcMdApi

	TraderFront string
	TraderApi   ctp.CThostFtdcTraderApi

	MdRequestID     int
	TraderRequestID int
}

func (g *GoCTPClient) GetMdRequestID() int {
	g.MdRequestID += 1
	return g.MdRequestID
}

func (g *GoCTPClient) GetTraderRequestID() int {
	g.TraderRequestID += 1
	return g.TraderRequestID
}

func NewDirectorCThostFtdcTraderSpi(v interface{}) ctp.CThostFtdcTraderSpi {
	return ctp.NewDirectorCThostFtdcTraderSpi(v)
}

type GoCThostFtdcTraderSpi struct {
	Client GoCTPClient
}

func (p *GoCThostFtdcTraderSpi) OnRspError(pRspInfo ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
	log.Println("GoCThostFtdcTraderSpi.OnRspError.")
	p.IsErrorRspInfo(pRspInfo)
}

func (p *GoCThostFtdcTraderSpi) OnFrontDisconnected(nReason int) {
	log.Printf("GoCThostFtdcTraderSpi.OnFrontDisconnected: %#v", nReason)
}

func (p *GoCThostFtdcTraderSpi) OnHeartBeatWarning(nTimeLapse int) {
	log.Printf("GoCThostFtdcTraderSpi.OnHeartBeatWarning: %#v", nTimeLapse)
}

func (p *GoCThostFtdcTraderSpi) OnFrontConnected() {
	log.Println("GoCThostFtdcTraderSpi.OnFrontConnected.")
	p.ReqUserLogin()
}

func (p *GoCThostFtdcTraderSpi) ReqUserLogin() {
	log.Println("GoCThostFtdcTraderSpi.ReqUserLogin.")

	req := ctp.NewCThostFtdcReqUserLoginField()
	req.SetBrokerID(p.Client.BrokerID)
	req.SetUserID(p.Client.InvestorID)
	req.SetPassword(p.Client.Password)

	iResult := p.Client.TraderApi.ReqUserLogin(req, p.Client.GetTraderRequestID())

	if iResult != 0 {
		log.Println("发送用户登录请求: 失败.")
	} else {
		log.Println("发送用户登录请求: 成功.")
	}
}

func (p *GoCThostFtdcTraderSpi) IsFlowControl(iResult int) bool {
	return ((iResult == -2) || (iResult == -3))
}

func (p *GoCThostFtdcTraderSpi) IsErrorRspInfo(pRspInfo ctp.CThostFtdcRspInfoField) bool {
	// 如果ErrorID != 0, 说明收到了错误的响应
	bResult := (pRspInfo.GetErrorID() != 0)
	if bResult {
		log.Printf("ErrorID=%v ErrorMsg=%v\n", pRspInfo.GetErrorID(), pRspInfo.GetErrorMsg())
	}
	return bResult
}

func (p *GoCThostFtdcTraderSpi) OnRspUserLogin(pRspUserLogin ctp.CThostFtdcRspUserLoginField, pRspInfo ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {

	log.Println("GoCThostFtdcTraderSpi.OnRspUserLogin.")
	if bIsLast && !p.IsErrorRspInfo(pRspInfo) {

		log.Printf("获取当前交易日  : %#v\n", p.Client.TraderApi.GetTradingDay())
		log.Printf("获取用户登录信息: %#v %#v %#v\n", pRspUserLogin.GetFrontID(), pRspUserLogin.GetSessionID(), pRspUserLogin.GetMaxOrderRef())

		// // 保存会话参数
		// FRONT_ID = pRspUserLogin->FrontID;
		// SESSION_ID = pRspUserLogin->SessionID;
		// int iNextOrderRef = atoi(pRspUserLogin->MaxOrderRef);
		// iNextOrderRef++;
		// sprintf(ORDER_REF, "%d", iNextOrderRef);
		// sprintf(EXECORDER_REF, "%d", 1);
		// sprintf(FORQUOTE_REF, "%d", 1);
		// sprintf(QUOTE_REF, "%d", 1);
		// ///获取当前交易日
		// cerr << "获取当前交易日 = " << pMdApi->GetTradingDay() << endl;
		// ///投资者结算结果确认
		p.ReqSettlementInfoConfirm()
	}
}

func (p *GoCThostFtdcTraderSpi) ReqSettlementInfoConfirm() {
	req := ctp.NewCThostFtdcSettlementInfoConfirmField()

	req.SetBrokerID(p.Client.BrokerID)
	req.SetInvestorID(p.Client.InvestorID)

	iResult := p.Client.TraderApi.ReqSettlementInfoConfirm(req, p.Client.GetTraderRequestID())

	if iResult != 0 {
		log.Println("投资者结算结果确认: 失败.")
	} else {
		log.Println("投资者结算结果确认: 成功.")
	}
}

func (p *GoCThostFtdcTraderSpi) OnRspSettlementInfoConfirm(pSettlementInfoConfirm ctp.CThostFtdcSettlementInfoConfirmField, pRspInfo ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
	//cerr << "--->>> " << "OnRspSettlementInfoConfirm" << endl
	log.Println("GoCThostFtdcTraderSpi.OnRspSettlementInfoConfirm.")
	if bIsLast && !p.IsErrorRspInfo(pRspInfo) {
		///请求查询合约
		p.ReqQryInstrument()
	}
}

func (p *GoCThostFtdcTraderSpi) ReqQryInstrument() {
	req := ctp.NewCThostFtdcQryInstrumentField()

	var id string = "cu1612"
	req.SetInstrumentID(id)

	for {
		iResult := p.Client.TraderApi.ReqQryInstrument(req, p.Client.GetTraderRequestID())

		if !p.IsFlowControl(iResult) {
			log.Printf("--->>> 请求查询合约: 成功 %#v\n", iResult)
			//break
		} else {
			log.Printf("--->>> 请求查询合约: 受到流控 %#v\n", iResult)
			time.Sleep(1 * time.Second)
		}
	}
}

func (p *GoCThostFtdcTraderSpi) OnRspQryInstrument(pInstrument ctp.CThostFtdcInstrumentField, pRspInfo ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
	log.Println("GoCThostFtdcTraderSpi.OnRspQryInstrument: ", pInstrument.GetInstrumentID(), pInstrument.GetExchangeID(),
		pInstrument.GetInstrumentName(), pInstrument.GetExchangeInstID(), pInstrument.GetProductID(), pInstrument.GetProductClass(),
		pInstrument.GetDeliveryYear(), pInstrument.GetDeliveryMonth(), pInstrument.GetMaxMarketOrderVolume(), pInstrument.GetMinMarketOrderVolume(),
		pInstrument.GetMaxLimitOrderVolume(), pInstrument.GetMinLimitOrderVolume(), pInstrument.GetVolumeMultiple(), pInstrument.GetPriceTick(),
		pInstrument.GetCreateDate(), pInstrument.GetOpenDate(), pInstrument.GetExpireDate(), pInstrument.GetStartDelivDate(), pInstrument.GetEndDelivDate(),
		nRequestID, bIsLast)
	if bIsLast && !p.IsErrorRspInfo(pRspInfo) {

		///请求查询合约
		p.ReqQryTradingAccount()
	}
}

func (p *GoCThostFtdcTraderSpi) ReqQryTradingAccount() {
	req := ctp.NewCThostFtdcQryTradingAccountField()
	req.SetBrokerID(p.Client.BrokerID)
	req.SetInvestorID(p.Client.InvestorID)

	for {
		iResult := p.Client.TraderApi.ReqQryTradingAccount(req, p.Client.GetTraderRequestID())

		if !p.IsFlowControl(iResult) {
			log.Printf("--->>> 请求查询资金账户: 成功 %#v\n", iResult)
			//break
		} else {
			log.Printf("--->>> 请求查询资金账户: 受到流控 %#v\n", iResult)
			time.Sleep(1 * time.Second)
		}
	}
}

func (p *GoCThostFtdcTraderSpi) OnRspQryTradingAccount(pTradingAccount ctp.CThostFtdcTradingAccountField, pRspInfo ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {

	log.Println("GoCThostFtdcTraderSpi.OnRspQryTradingAccount.")

	if bIsLast && !p.IsErrorRspInfo(pRspInfo) {
		///请求查询投资者持仓
		p.ReqQryInvestorPosition()
	}
}

func (p *GoCThostFtdcTraderSpi) ReqQryInvestorPosition() {

	req := ctp.NewCThostFtdcQryInvestorPositionField()
	req.SetBrokerID(p.Client.BrokerID)
	req.SetInvestorID(p.Client.InvestorID)
	req.SetInstrumentID("cu1612")

	for {
		iResult := p.Client.TraderApi.ReqQryInvestorPosition(req, p.Client.GetTraderRequestID())

		if !p.IsFlowControl(iResult) {
			log.Printf("--->>> 请求查询投资者持仓: 成功 %#v\n", iResult)
			//break;
		} else {
			log.Printf("--->>> 请求查询投资者持仓: 受到流控 %#v\n", iResult)
			time.Sleep(1 * time.Second)
		}
	}
}

func (p *GoCThostFtdcTraderSpi) OnRspQryInvestorPosition(pInvestorPosition ctp.CThostFtdcInvestorPositionField, pRspInfo ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
	log.Println("GoCThostFtdcTraderSpi.OnRspQryInvestorPosition.")

	if bIsLast && !p.IsErrorRspInfo(pRspInfo) {
		// ///报单录入请求
		// ReqOrderInsert();
		// //执行宣告录入请求
		// ReqExecOrderInsert();
		// //询价录入
		// ReqForQuoteInsert();
		// //做市商报价录入
		// ReqQuoteInsert();
	}
}

func init() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	log.SetPrefix("CTP: ")
}

func main() {

	if len(os.Args) < 2 {
		log.Fatal("usage: ./ctp_trader_example -BrokerID 9999 -InvestorID 000000 -Password 000000 -MarketFront tcp://180.168.146.187:10010 -TradeFront tcp://180.168.146.187:10000")
	}

	flag.Parse()

	CTP = GoCTPClient{
		BrokerID:   *broker_id,
		InvestorID: *investor_id,
		Password:   *pass_word,

		MdFront: *market_front,
		MdApi:   ctp.CThostFtdcMdApiCreateFtdcMdApi(),

		TraderFront: *trade_front,
		TraderApi:   ctp.CThostFtdcTraderApiCreateFtdcTraderApi(),

		MdRequestID:     0,
		TraderRequestID: 0,
	}

	pTraderSpi := ctp.NewDirectorCThostFtdcTraderSpi(&GoCThostFtdcTraderSpi{Client: CTP})

	CTP.TraderApi.RegisterSpi(pTraderSpi)                         // 注册事件类
	CTP.TraderApi.SubscribePublicTopic(0 /*THOST_TERT_RESTART*/)  // 注册公有流
	CTP.TraderApi.SubscribePrivateTopic(0 /*THOST_TERT_RESTART*/) // 注册私有流
	CTP.TraderApi.RegisterFront(CTP.TraderFront)
	CTP.TraderApi.Init()

	CTP.TraderApi.Join()
	CTP.TraderApi.Release()
}
