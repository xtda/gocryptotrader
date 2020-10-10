// Package backtest to be written...
package backtest

import (
	"time"

	"github.com/thrasher-corp/gocryptotrader/currency"
	"github.com/thrasher-corp/gocryptotrader/exchanges/order"
)

// Event contains the time and the currency pair.
type Event struct {
	Time         time.Time
	CurrencyPair currency.Pair
}

// Signal contains the event, amount, price and direction.
type Signal struct {
	Event
	Amount    float64
	Price     float64
	Direction order.Side
}
