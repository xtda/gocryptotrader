// Package backtest to be written.
package backtest

// DP defaults to 8.
const DP = 8

// BackTest struct contains handlers for
// data, strategy, portfolio, orderbook, exchange and event queue.
type BackTest struct {
	data       DataHandler
	strategy   StrategyHandler
	portfolio  PortfolioHandler
	orderbook  OrderBook
	exchange   ExecutionHandler
	statistic  StatisticHandler
	eventQueue []EventHandler
}
