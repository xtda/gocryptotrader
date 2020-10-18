// Package backtest to be written...
package backtest

import (
	"math"
)

// ExecuteOrder creates and returns a new fill with the event, amount and
// price from the order event. From the exchange, it will set the direction, commission,
// exchange fee, and calculates the cost.
func (e *Exchange) ExecuteOrder(order OrderEvent, data DataHandler) (*Fill, error) {
	f := &Fill{
		Event: Event{
			Time:         order.GetTime(),
			CurrencyPair: order.Pair(),
		},
		Amount: order.GetAmount(),
		Price:  data.Latest().LatestPrice(),
	}
	f.Direction = order.GetDirection()
	f.Commission = e.calculateCommission(f.Amount, f.Price)
	f.ExchangeFee = e.calculateExchangeFee()
	f.Cost = e.calculateCost(f.Commission, f.ExchangeFee)

	return f, nil
}

func (e *Exchange) calculateCommission(amount, price float64) float64 {
	return math.Floor(amount*price*e.CommissionRate*10000) / 10000
}

func (e *Exchange) calculateExchangeFee() float64 {
	return e.ExchangeFee
}

func (e *Exchange) calculateCost(commission, fee float64) float64 {
	return commission + fee
}
