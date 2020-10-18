// Package backtest to be written.
package backtest

import (
	"sort"

	"github.com/shopspring/decimal"
)

// Load specified data into Candle format
// this is empty and loader types will have their own implementation.
func (d *Data) Load() error {
	return nil
}

// Reset loaded data to blank state.
func (d *Data) Reset() {
	d.latest = nil
	d.offset = 0
	d.stream = nil
}

// Stream will return entire data list.
func (d *Data) Stream() []DataEventHandler {
	return d.stream
}

// Next will return the next event in the list and also shift the offset one.
func (d *Data) Next() (dh DataEventHandler, ok bool) {
	if len(d.stream) <= d.offset {
		return nil, false
	}

	ret := d.stream[d.offset]
	d.offset++
	d.latest = ret
	return ret, true
}

// History will return all previous data events that have happened.
func (d *Data) History() []DataEventHandler {
	return d.stream[:d.offset]
}

// Latest will return latest data event.
func (d *Data) Latest() DataEventHandler {
	return d.latest
}

// List will return an array of data events.
func (d *Data) List() []DataEventHandler {
	return d.stream[d.offset:]
}

// SortStream will sort the data events by time.
func (d *Data) SortStream() {
	sort.Slice(d.stream, func(i, j int) bool {
		b1 := d.stream[i]
		b2 := d.stream[j]

		if b1.GetTime().Equal(b2.GetTime()) {
			return b1.Pair().String() < b2.Pair().String()
		}
		return b1.GetTime().Before(b2.GetTime())
	})
}

// StreamOpen to be implemented.
func (d *Data) StreamOpen() []float64 {
	return []float64{}
}

// StreamHigh to be implemented.
func (d *Data) StreamHigh() []float64 {
	return []float64{}
}

// StreamLow to be implemented.
func (d *Data) StreamLow() []float64 {
	return []float64{}
}

// StreamClose to be implemented.
func (d *Data) StreamClose() []float64 {
	return []float64{}
}

// StreamVol to be implemented.
func (d *Data) StreamVol() []float64 {
	return []float64{}
}

// Offset returns the offset.
func (d *Data) Offset() int {
	return d.offset
}

// DataType returns the DataType candle, default is iota.
func (c *Candle) DataType() DataType {
	return DataTypeCandle
}

// LatestPrice returns the price at closing for the candle.
func (c *Candle) LatestPrice() float64 {
	return c.Close
}

// LatestPrice returns the latest rounded price.
func (t *Tick) LatestPrice() float64 {
	bid := decimal.NewFromFloat(t.Bid)
	ask := decimal.NewFromFloat(t.Ask)
	diff := decimal.New(2, 0)
	latest, _ := bid.Add(ask).Div(diff).Round(DP).Float64()
	return latest
}

// DataType returns the DataTypeTick.
func (t *Tick) DataType() DataType {
	return DataTypeTick
}

// Spread returns the spread from the bid price to the asking price.
func (t *Tick) Spread() float64 {
	return t.Bid - t.Ask
}
