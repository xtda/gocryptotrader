package backtest

import (
	"time"

	"github.com/thrasher-corp/gocryptotrader/currency"
	"github.com/thrasher-corp/gocryptotrader/exchanges/order"
)

type Event struct {
	Time         time.Time
	CurrencyPair currency.Pair
}
type Candle struct {
	Event
	Timestamp time.Time
	Open      float64
	Close     float64
	Low       float64
	High      float64
	Volume    float64
}

type DataEvent struct {
	Metrics map[string]float64
}

type Tick struct {
	Event
	DataEvent
	Bid float64
	Ask float64
}

type Signal struct {
	Event
	Amount    float64
	Price     float64
	Direction order.Side
}