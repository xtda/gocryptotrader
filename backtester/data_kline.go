// Package backtest to be written.
package backtest

import (
	"errors"

	"github.com/thrasher-corp/gocryptotrader/exchanges/kline"
)

// DataFromKline struct contains the kline item and data.
type DataFromKline struct {
	Item kline.Item

	Data
}

// Load will load the data from the kline exchange.
func (d *DataFromKline) Load() error {
	if len(d.Item.Candles) == 0 {
		return errors.New("no candle data provided")
	}

	data := make([]DataEventHandler, len(d.Item.Candles))
	for i := range d.Item.Candles {
		data[i] = &Candle{
			Event: Event{
				Time: d.Item.Candles[i].Time, CurrencyPair: d.Item.Pair,
			},
			Open:   d.Item.Candles[i].Open,
			High:   d.Item.Candles[i].High,
			Low:    d.Item.Candles[i].Low,
			Close:  d.Item.Candles[i].Close,
			Volume: d.Item.Candles[i].Volume,
		}
	}
	d.stream = data
	d.SortStream()
	return nil
}

// StreamOpen returns the open price for each candle in the stream.
func (d *DataFromKline) StreamOpen() []float64 {
	ret := make([]float64, len(d.stream))
	for x := range d.stream[d.offset:] {
		ret[x] = d.stream[x].(*Candle).Open
	}
	return ret
}

// StreamHigh returns the high price for each candle in the stream.
func (d *DataFromKline) StreamHigh() []float64 {
	ret := make([]float64, len(d.stream))
	for x := range d.stream[d.offset:] {
		ret[x] = d.stream[x].(*Candle).High
	}
	return ret
}

// StreamLow returns the low price for each candle in the stream.
func (d *DataFromKline) StreamLow() []float64 {
	ret := make([]float64, len(d.stream))
	for x := range d.stream[d.offset:] {
		ret[x] = d.stream[x].(*Candle).Low
	}
	return ret
}

// StreamClose returns the closing price for each candle in the stream.
func (d *DataFromKline) StreamClose() []float64 {
	ret := make([]float64, len(d.stream))
	for x := range d.stream[d.offset:] {
		ret[x] = d.stream[x].(*Candle).Close
	}
	return ret
}

// StreamVol returns the volume of each candle in the stream.
func (d *DataFromKline) StreamVol() []float64 {
	ret := make([]float64, len(d.stream))
	for x := range d.stream[d.offset:] {
		ret[x] = d.stream[x].(*Candle).Volume
	}
	return ret
}
