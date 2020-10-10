// Package backtest to be written...
package backtest

import (
	"github.com/thrasher-corp/gocryptotrader/currency"
	"github.com/thrasher-corp/gocryptotrader/exchanges/order"
)

// IsOrder returns true.
func (o *Order) IsOrder() bool {
	return true
}

// SetDirection sets the direction on the order.
func (o *Order) SetDirection(s order.Side) {
	o.Direction = s
}

// GetDirection returns the side of the order.
func (o *Order) GetDirection() order.Side {
	return o.Direction
}

// SetAmount sets the amount on the order.
func (o *Order) SetAmount(i float64) {
	o.Amount = i
}

// GetAmount returns the amount of the order.
func (o *Order) GetAmount() float64 {
	return o.Amount
}

// Pair returns the currency pair.
func (o *Order) Pair() currency.Pair {
	return o.CurrencyPair
}

// Cancel sets the order status to pending cancel.
func (o *Order) Cancel() {
	o.Status = order.PendingCancel
}

// GetStatus returns the status of the order.
func (o *Order) GetStatus() order.Status {
	return o.Status
}

// SetID sets the id of the order.
func (o *Order) SetID(id int) {
	o.id = id
}

// ID returns the id.
func (o *Order) ID() int {
	return o.id
}

// Limit returns the limit of the order.
func (o *Order) Limit() float64 {
	return o.limit
}

// IsLeveraged returns whether or not the order is leveraged.
func (o *Order) IsLeveraged() bool {
	return o.leverage > 0.0
}

// Leverage returns the leverage amount.
func (o *Order) Leverage() float64 {
	return o.leverage
}

// SetLeverage sets the leverage amount.
func (o *Order) SetLeverage(l float64) {
	o.leverage = l
}
