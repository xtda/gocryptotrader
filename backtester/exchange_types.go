// Package backtest to be written...
package backtest

import "github.com/thrasher-corp/gocryptotrader/currency"

// Exchange contains the currency pair, exchange fee, commission rate, maker/taker fee, and orderbook.
type Exchange struct {
	CurrencyPair   currency.Pair
	ExchangeFee    float64
	CommissionRate float64
	MakerFee       float64
	TakerFee       float64
	Orderbook      OrderBook
}
