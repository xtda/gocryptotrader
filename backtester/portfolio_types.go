// Package backtest to be written...
package backtest

import "github.com/thrasher-corp/gocryptotrader/currency"

// Portfolio struct contains the initialFunds, funds, holdings, transactions
// size manager, and risk manager.
type Portfolio struct {
	initialFunds float64
	funds        float64
	holdings     map[currency.Pair]Positions
	transactions []FillEvent
	sizeManager  SizeHandler
	riskManager  RiskHandler
}
