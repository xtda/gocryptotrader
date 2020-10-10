// Package backtest to be written.
package backtest

import (
	"time"

	"github.com/thrasher-corp/gocryptotrader/currency"
	"github.com/thrasher-corp/gocryptotrader/exchanges/order"
)

// IsEvent returns true.
func (e *Event) IsEvent() bool {
	return true
}

// GetTime returns the time of the event.
func (e *Event) GetTime() time.Time {
	return e.Time
}

// Pair returns the currency pair.
func (e *Event) Pair() currency.Pair {
	return e.CurrencyPair
}

// IsSignal returns true.
func (s *Signal) IsSignal() bool {
	return true
}

// SetDirection sets the direction of the side.
func (s *Signal) SetDirection(st order.Side) {
	s.Direction = st
}

// GetDirection returns the side for the signal.
func (s *Signal) GetDirection() order.Side {
	return s.Direction
}

// Pair returns the currency pair for the signal.
func (s *Signal) Pair() currency.Pair {
	return s.CurrencyPair
}

// SetAmount sets the amount on the signal.
func (s *Signal) SetAmount(f float64) {
	s.Amount = f
}

// GetAmount gets the amount from the signal.
func (s *Signal) GetAmount() float64 {
	return s.Amount
}

// GetPrice returns the price from the signal.
func (s *Signal) GetPrice() float64 {
	return s.Price
}

// SetPrice sets the price on the signal.
func (s *Signal) SetPrice(f float64) {
	s.Price = f
}
