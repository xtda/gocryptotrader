package backtest

import gctorder "github.com/thrasher-corp/gocryptotrader/exchanges/order"

type Portfolio struct {
	initialFunds float64
	funds        float64
	Holdings     Position
	OrderBook    []OrderEvent
	sizeManager  SizeHandler
	riskManager  RiskHandler
}

type PortfolioHandler interface {
	OnSignal(SignalEvent, DataHandler) (*Order, error)
	OnFill(*Order) (*Order, error)
	IsInvested() (Position, bool)
	IsLong() (Position, bool)
	IsShort() (Position, bool)
	Update(DataEvent)

	InitialFunds() float64
	SetInitialFunds(float64)
	Funds() float64
	SetFunds(float64)

	Value() float64

	Reset() error

	Order(price float64, amount float64, side gctorder.Side)

	Position() Position
}