// Package backtest to be written...
package backtest

import (
	"github.com/thrasher-corp/gocryptotrader/exchanges/order"
)

// Fill struct contains the event, direction, amount,
// price, commission, exchange fee, and cost.
type Fill struct {
	Event
	// Exchange    string
	Direction   order.Side
	Amount      float64
	Price       float64
	Commission  float64
	ExchangeFee float64
	Cost        float64
}
