// Package backtest to be written.
package backtest

// DataType unsigned int.
type DataType uint8

// DataTypeCandle defaults the datatype to iota.
const (
	DataTypeCandle DataType = iota
	DataTypeTick
)

// Candle contains the event, open, close, low, high, and volume.
type Candle struct {
	Event
	Open   float64
	Close  float64
	Low    float64
	High   float64
	Volume float64
}

// Tick contains the event, bid price and asking price.
type Tick struct {
	Event
	Bid float64
	Ask float64
}

// Orderbook contains the event and the list of bids/asks.
type Orderbook struct {
	Event
	Bids []float64
	Asks []float64
}

// Data contains the latest event and the stream of events.
type Data struct {
	latest DataEventHandler
	stream []DataEventHandler

	offset int
}
