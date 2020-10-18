// Package backtest to be written...
package backtest

import (
	"errors"
	"math"

	"github.com/thrasher-corp/gocryptotrader/currency"
	gctorder "github.com/thrasher-corp/gocryptotrader/exchanges/order"
)

// EvaluateOrder evaluates the risk tolerance if the order is leveraged.
// TODO implement risk manager.
func (r *Risk) EvaluateOrder(order OrderEvent, _ DataEventHandler, _ map[currency.Pair]Positions) (*Order, error) {
	retOrder := order.(*Order)

	if order.IsLeveraged() {
		// handle risk
	}
	return retOrder, nil
}

// SizeOrder determines the size of the order
// TODO implement risk manager.
func (s *Size) SizeOrder(order OrderEvent, _ DataEventHandler, _ PortfolioHandler) (*Order, error) {
	retOrder := order.(*Order)

	if (s.DefaultSize == 0) || (s.DefaultValue == 0) {
		return nil, errors.New("no defaultSize or defaultValue set")
	}

	switch retOrder.GetDirection() {
	case gctorder.Buy:
	case gctorder.Sell:
		retOrder.SetAmount(s.setDefaultSize(retOrder.Price))
	}
	return retOrder, nil
}

func (s *Size) setDefaultSize(price float64) float64 {
	if (s.DefaultSize * price) > s.DefaultValue {
		return math.Floor(s.DefaultValue / price)
	}
	return s.DefaultSize
}
