// Package backtest to be written...
package backtest

import (
	"github.com/shopspring/decimal"
	"github.com/thrasher-corp/gocryptotrader/exchanges/order"
)

// SetDirection sets the direction on the fill.
func (f *Fill) SetDirection(s order.Side) {
	f.Direction = s
}

// GetDirection returns the side of the fill.
func (f *Fill) GetDirection() order.Side {
	return f.Direction
}

// SetAmount sets the amount on the fill.
func (f *Fill) SetAmount(i float64) {
	f.Amount = i
}

// GetAmount returns the amount of the fill.
func (f *Fill) GetAmount() float64 {
	return f.Amount
}

// GetPrice returns the price of the fill.
func (f *Fill) GetPrice() float64 {
	return f.Price
}

// GetCommission returns the commission of the fill.
func (f *Fill) GetCommission() float64 {
	return f.Commission
}

// GetExchangeFee returns the exchange fee of the fill.
func (f *Fill) GetExchangeFee() float64 {
	return f.ExchangeFee
}

// GetCost returns the cost of the fill.
func (f *Fill) GetCost() float64 {
	return f.Cost
}

// Value returns the value.
func (f *Fill) Value() float64 {
	amount := decimal.NewFromFloat(f.Amount)
	price := decimal.NewFromFloat(f.Price)
	value, _ := amount.Mul(price).Round(DP).Float64()
	return value
}

// NetValue returns the net value.
func (f *Fill) NetValue() float64 {
	amount := decimal.NewFromFloat(f.Amount)
	price := decimal.NewFromFloat(f.Price)
	cost := decimal.NewFromFloat(f.Cost)

	if f.Direction == order.Buy {
		netValue, _ := amount.Mul(price).Add(cost).Round(DP).Float64()
		return netValue
	}
	netValue, _ := amount.Mul(price).Sub(cost).Round(DP).Float64()
	return netValue
}
