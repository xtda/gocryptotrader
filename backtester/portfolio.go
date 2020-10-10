// Package backtest to be written...
package backtest

import (
	"errors"

	"github.com/shopspring/decimal"
	"github.com/thrasher-corp/gocryptotrader/currency"
	gctorder "github.com/thrasher-corp/gocryptotrader/exchanges/order"
)

// SetSizeManager sets the size manager from the side handler.
func (p *Portfolio) SetSizeManager(size SizeHandler) {
	p.sizeManager = size
}

// Reset resets the portfolio to the default state.
func (p *Portfolio) Reset() {
	p.funds = 0
	p.holdings = nil
	p.transactions = nil
}

// OnSignal returns an order based on the signal given.
func (p *Portfolio) OnSignal(signal SignalEvent, data DataHandler) (*Order, error) {
	var limit float64

	if signal.GetDirection() == "" {
		return &Order{}, errors.New("invalid Direction")
	}

	currAmount := p.holdings[signal.Pair()].amount
	currFunds := p.Funds()
	currPrice := data.Latest().LatestPrice()

	if (signal.GetDirection() == gctorder.Sell || signal.GetDirection() == gctorder.Ask) && currAmount <= signal.GetAmount() {
		return &Order{}, errors.New("no holdings to sell")
	}

	if (signal.GetDirection() == gctorder.Buy || signal.GetDirection() == gctorder.Bid) && currFunds <= signal.GetPrice()*currPrice {
		return &Order{}, errors.New("not enough funds to buy")
	}

	initialOrder := &Order{
		Event: Event{
			Time:         signal.GetTime(),
			CurrencyPair: signal.Pair(),
		},
		Direction: signal.GetDirection(),
		Price:     signal.GetPrice(),
		Amount:    signal.GetAmount(),
		OrderType: gctorder.Market,
		limit:     limit,
	}

	latest := data.Latest()
	sizedOrder, err := p.sizeManager.SizeOrder(initialOrder, latest, p)
	if err != nil {
		return nil, err
	}

	order, err := p.riskManager.EvaluateOrder(sizedOrder, latest, p.holdings)
	if err != nil {
		return nil, err
	}

	return order, nil
}

// OnFill returns a fill based on the fill event.
func (p *Portfolio) OnFill(fill FillEvent, _ DataHandler) (*Fill, error) {
	if p.holdings == nil {
		p.holdings = make(map[currency.Pair]Positions)
	}

	if pos, ok := p.holdings[fill.Pair()]; ok {
		pos.Update(fill)
		p.holdings[fill.Pair()] = pos
	} else {
		pos := Positions{}
		pos.Create(fill)
		p.holdings[fill.Pair()] = pos
	}

	if fill.GetDirection() == gctorder.Buy {
		p.funds -= fill.NetValue()
	} else if fill.GetDirection() == gctorder.Sell || fill.GetDirection() == gctorder.Ask {
		p.funds += fill.NetValue()
	}

	p.transactions = append(p.transactions, fill)

	return fill.(*Fill), nil
}

// IsInvested determines if the portfolio contains the currency pair with an amount positive or negative.
func (p *Portfolio) IsInvested(pair currency.Pair) (pos Positions, ok bool) {
	pos, ok = p.holdings[pair]
	if ok && (pos.amount != 0) {
		return pos, true
	}
	return pos, false
}

// IsLong determines if portfolio contains the currency pair and the position amount is greater then zero.
func (p *Portfolio) IsLong(pair currency.Pair) (pos Positions, ok bool) {
	pos, ok = p.holdings[pair]
	if ok && (pos.amount > 0) {
		return pos, true
	}
	return pos, false
}

// IsShort determines if portfolio contains the currency pair and the position amount is less then zero.
func (p *Portfolio) IsShort(pair currency.Pair) (pos Positions, ok bool) {
	pos, ok = p.holdings[pair]
	if ok && (pos.amount < 0) {
		return pos, true
	}
	return pos, false
}

// Update will set the pair and it's position to the portfolio's holdings.
func (p *Portfolio) Update(d DataEventHandler) {
	if pos, ok := p.IsInvested(d.Pair()); ok {
		pos.UpdateValue(d)
		p.holdings[d.Pair()] = pos
	}
}

// SetInitialFunds sets the initial funds available.
func (p *Portfolio) SetInitialFunds(initial float64) {
	p.initialFunds = initial
}

// InitialFunds gets the initialFunds sets on the portfolio.
func (p *Portfolio) InitialFunds() float64 {
	return p.initialFunds
}

// SetFunds sets the funds on the portfolio.
func (p *Portfolio) SetFunds(funds float64) {
	p.funds = funds
}

// Funds gets the funds from the portfolio.
func (p *Portfolio) Funds() float64 {
	return p.funds
}

// Value sums up the market value from all of currency pairs in the holdings.
func (p *Portfolio) Value() float64 {
	holdingValue := decimal.NewFromFloat(0)
	for x := range p.holdings {
		marketValue := decimal.NewFromFloat(p.holdings[x].marketValue)
		holdingValue = holdingValue.Add(marketValue)
	}

	funds := decimal.NewFromFloat(p.funds)
	value, _ := funds.Add(holdingValue).Round(4).Float64()
	return value
}

// ViewHoldings returns the map currency pairs and their positions.
func (p *Portfolio) ViewHoldings() map[currency.Pair]Positions {
	return p.holdings
}
