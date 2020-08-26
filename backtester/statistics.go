package backtest

import (
	"errors"
	"fmt"
	"time"

	"github.com/shopspring/decimal"
	"gonum.org/v1/gonum/stat"
)

func (s *Statistic) Update(d DataEventHandler, p PortfolioHandler) {
	if s.initialBuy == 0 {
		s.initialBuy = p.InitialFunds() / d.LatestPrice()
	}

	e := equityPoint{}
	e.timestamp = d.GetTime()
	e.equity = p.Value()

	e.buyAndHoldValue = s.initialBuy * d.LatestPrice()

	if len(s.equity) > 0 {
		e = s.calcEquityReturn(e)
	}

	if len(s.equity) > 0 {
		e = s.calcDrawdown(e)
	}

	if e.equity >= s.high.equity {
		s.high = e
	}
	if e.equity <= s.low.equity {
		s.low = e
	}

	s.equity = append(s.equity, e)
}

func (s *Statistic) TrackEvent(e EventHandler) {
	s.eventHistory = append(s.eventHistory, e)
}

func (s *Statistic) Events() []EventHandler {
	return s.eventHistory
}

func (s *Statistic) TrackTransaction(f FillEvent) {
	s.transactionHistory = append(s.transactionHistory, f)
}

func (s *Statistic) Transactions() []FillEvent {
	return s.transactionHistory
}

func (s *Statistic) Reset() {
	s.eventHistory = nil
	s.transactionHistory = nil
	s.equity = nil
	s.high = equityPoint{}
	s.low = equityPoint{}
}

func (s *Statistic) ReturnResult() Results {
	results := Results{
		TotalEvents:       len(s.Events()),
		TotalTransactions: len(s.Transactions()),
		SharpieRatio:      s.SharpRatio(0),
	}
	for v := range s.Transactions() {
		results.Transactions = append(results.Transactions, resultTransactions{
			time:      s.Transactions()[v].GetTime(),
			direction: s.Transactions()[v].GetDirection(),
			price:     s.Transactions()[v].GetPrice(),
			amount:    s.Transactions()[v].GetAmount(),
		})
	}
	return results
}

func (s *Statistic) PrintResult() {
	fmt.Printf("Counted %d total events.\n", len(s.Events()))

	fmt.Printf("Counted %d total transactions:\n", len(s.Transactions()))
	for k, v := range s.Transactions() {
		fmt.Printf("%d. Transaction: %v Action: %v Price: %f Amount %f\n", k+1, v.GetTime().Format("2006-01-02 15:04:05.999999999"), v.GetDirection(), v.GetPrice(), v.GetAmount())
	}

	result, _ := s.TotalEquityReturn()

	fmt.Println("TotalEquity:", result, "MaxDrawdown:", s.MaxDrawdown())
}

func (s *Statistic) TotalEquityReturn() (r float64, err error) {
	firstEquityPoint, ok := s.firstEquityPoint()
	if !ok {
		return r, errors.New("could not calculate totalEquityReturn, no equity points found")
	}
	firstEquity := decimal.NewFromFloat(firstEquityPoint.equity)

	lastEquityPoint, _ := s.lastEquityPoint()
	lastEquity := decimal.NewFromFloat(lastEquityPoint.equity)

	totalEquityReturn := lastEquity.Sub(firstEquity).Div(firstEquity)
	total, _ := totalEquityReturn.Round(DP).Float64()
	return total, nil
}

func (s *Statistic) MaxDrawdown() float64 {
	_, ep := s.maxDrawdownPoint()
	return ep.drawdown
}

func (s *Statistic) MaxDrawdownTime() time.Time {
	_, ep := s.maxDrawdownPoint()
	return ep.timestamp
}

func (s *Statistic) MaxDrawdownDuration() (d time.Duration) {
	i, ep := s.maxDrawdownPoint()

	if len(s.equity) == 0 {
		return d
	}

	maxPoint := equityPoint{}
	for index := i; index >= 0; index-- {
		if s.equity[index].equity > maxPoint.equity {
			maxPoint = s.equity[index]
		}
	}

	d = ep.timestamp.Sub(maxPoint.timestamp)
	return d
}

func (s *Statistic) SharpRatio(riskfree float64) float64 {
	var equityReturns = make([]float64, len(s.equity))

	for i, v := range s.equity {
		equityReturns[i] = v.equityReturn
	}
	mean, stddev := stat.MeanStdDev(equityReturns, nil)

	sharp := (mean - riskfree) / stddev
	return sharp
}

func (s *Statistic) SortinoRatio(riskfree float64) float64 {
	var equityReturns = make([]float64, len(s.equity))

	for i, v := range s.equity {
		equityReturns[i] = v.equityReturn
	}
	mean := stat.Mean(equityReturns, nil)

	// sortino uses the stddev of only the negativ returns
	var negReturns = []float64{}
	for _, v := range equityReturns {
		if v < 0 {
			negReturns = append(negReturns, v)
		}
	}
	stdDev := stat.StdDev(negReturns, nil)

	sortino := (mean - riskfree) / stdDev
	return sortino
}

func (s *Statistic) ViewEquityHistory() {
	fmt.Println(s.equity)
}

func (s *Statistic) firstEquityPoint() (ep equityPoint, ok bool) {
	if len(s.equity) == 0 {
		return ep, false
	}
	ep = s.equity[0]

	return ep, true
}

func (s *Statistic) lastEquityPoint() (ep equityPoint, ok bool) {
	if len(s.equity) == 0 {
		return ep, false
	}
	ep = s.equity[len(s.equity)-1]

	return ep, true
}

func (s *Statistic) calcEquityReturn(e equityPoint) equityPoint {
	last, ok := s.lastEquityPoint()
	if !ok {
		e.equityReturn = 0
		return e
	}

	lastEquity := decimal.NewFromFloat(last.equity)
	currentEquity := decimal.NewFromFloat(e.equity)

	if lastEquity.Equal(decimal.Zero) {
		e.equityReturn = 1
		return e
	}

	equityReturn := currentEquity.Sub(lastEquity).Div(lastEquity)
	e.equityReturn, _ = equityReturn.Round(DP).Float64()

	return e
}

func (s *Statistic) calcDrawdown(e equityPoint) equityPoint {
	if s.high.equity == 0 {
		e.drawdown = 0
		return e
	}

	lastHigh := decimal.NewFromFloat(s.high.equity)
	equity := decimal.NewFromFloat(e.equity)

	if equity.GreaterThanOrEqual(lastHigh) {
		e.drawdown = 0
		return e
	}

	drawdown := equity.Sub(lastHigh).Div(lastHigh)
	e.drawdown, _ = drawdown.Round(DP).Float64()

	return e
}

func (s *Statistic) maxDrawdownPoint() (i int, ep equityPoint) {
	if len(s.equity) == 0 {
		return 0, ep
	}

	var maxDrawdown = 0.0
	var index = 0

	for i, ep := range s.equity {
		if ep.drawdown < maxDrawdown {
			maxDrawdown = ep.drawdown
			index = i
		}
	}

	return index, s.equity[index]
}