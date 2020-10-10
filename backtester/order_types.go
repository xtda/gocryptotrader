// Package backtest to be written...
package backtest

import (
	"sync"

	"github.com/thrasher-corp/gocryptotrader/exchanges/order"
)

// Order contains the event, id direction, status, price, amount, order type, limit and leverage.
type Order struct {
	Event
	id        int
	Direction order.Side
	Status    order.Status
	Price     float64
	Amount    float64
	OrderType order.Type
	limit     float64
	leverage  float64
}

// OrderBook contains the list of orders, the order history and the counter.
type OrderBook struct {
	counter int
	orders  []OrderEvent
	history []OrderEvent

	m sync.Mutex
}
