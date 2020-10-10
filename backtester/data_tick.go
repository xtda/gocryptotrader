// Package backtest to be written.
package backtest

import (
	"errors"

	"github.com/thrasher-corp/gocryptotrader/exchanges/ticker"
)

// DataFromTick struct contains the list of ticker prices.
type DataFromTick struct {
	ticks []*ticker.Price

	Data
}

// Load will load the ticks from the exchange and convert them to
// backtest ticks and set the stream.
func (d *DataFromTick) Load() error {
	if len(d.ticks) == 0 {
		return errors.New("no tick data provided")
	}

	data := make([]DataEventHandler, len(d.ticks))
	for i := range d.ticks {
		data[i] = &Tick{
			Event: Event{
				Time:         d.ticks[i].LastUpdated,
				CurrencyPair: d.ticks[i].Pair,
			},
			Ask: d.ticks[i].Ask,
			Bid: d.ticks[i].Bid,
		}
	}
	d.stream = data
	d.SortStream()
	return nil
}
