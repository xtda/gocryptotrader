package binance

import (
	"testing"
	"time"

	"github.com/thrasher-corp/gocryptotrader/common"
	"github.com/thrasher-corp/gocryptotrader/core"
	"github.com/thrasher-corp/gocryptotrader/currency"
	exchange "github.com/thrasher-corp/gocryptotrader/exchanges"
	"github.com/thrasher-corp/gocryptotrader/exchanges/asset"
	"github.com/thrasher-corp/gocryptotrader/exchanges/kline"
	"github.com/thrasher-corp/gocryptotrader/exchanges/order"
	"github.com/thrasher-corp/gocryptotrader/exchanges/trade"
	"github.com/thrasher-corp/gocryptotrader/portfolio/withdraw"
)

// Please supply your own keys here for due diligence testing
const (
	apiKey                  = ""
	apiSecret               = ""
	canManipulateRealOrders = false
)

var b Binance

func areTestAPIKeysSet() bool {
	return b.ValidateAPICredentials()
}

func setFeeBuilder() *exchange.FeeBuilder {
	return &exchange.FeeBuilder{
		Amount:        1,
		FeeType:       exchange.CryptocurrencyTradeFee,
		Pair:          currency.NewPair(currency.BTC, currency.LTC),
		PurchasePrice: 1,
	}
}

func TestGetExchangeInfo(t *testing.T) {
	t.Parallel()
	_, err := b.GetExchangeInfo()
	if err != nil {
		t.Error(err)
	}
}

func TestFetchTradablePairs(t *testing.T) {
	t.Parallel()

	_, err := b.FetchTradablePairs(asset.Spot)
	if err != nil {
		t.Error("Binance FetchTradablePairs(asset asets.AssetType) error", err)
	}
}

func TestGetOrderBook(t *testing.T) {
	t.Parallel()

	_, err := b.GetOrderBook(OrderBookDataRequestParams{
		Symbol: "BTCUSDT",
		Limit:  10,
	})

	if err != nil {
		t.Error("Binance GetOrderBook() error", err)
	}
}

func TestGetRecentTrades(t *testing.T) {
	t.Parallel()

	_, err := b.GetRecentTrades(RecentTradeRequestParams{
		Symbol: "BTCUSDT",
		Limit:  15,
	})

	if err != nil {
		t.Error("Binance GetRecentTrades() error", err)
	}
}

func TestGetHistoricalTrades(t *testing.T) {
	t.Parallel()

	_, err := b.GetHistoricalTrades("BTCUSDT", 5, 0)
	if !mockTests && err == nil {
		t.Error("Binance GetHistoricalTrades() expecting error")
	}
	if mockTests && err == nil {
		t.Error("Binance GetHistoricalTrades() error", err)
	}
}

func TestExchangeHistory(t *testing.T) {
	t.Parallel()
	req := &trade.HistoryRequest{
		Pair:           currency.NewPairFromString("BTCUSDT"),
		TimestampStart: time.Now(),
	}
	_, err := b.GetExchangeHistory(req)
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetAggregatedTrades(t *testing.T) {
	t.Parallel()
	startTime := time.Unix(1588636800, 0)
	endTime := time.Unix(1588640400, 0)
	_, err := b.GetAggregatedTrades("BTCUSDT", 5, startTime, endTime)
	if err != nil {
		t.Error("Binance GetAggregatedTrades() error", err)
	}
}

func TestGetSpotKline(t *testing.T) {
	t.Parallel()
	_, err := b.GetSpotKline(KlinesRequestParams{
		Symbol:    "BTCUSDT",
		Interval:  kline.FiveMin.Short(),
		Limit:     24,
		StartTime: time.Unix(1577836800, 0).Unix() * 1000,
		EndTime:   time.Unix(1580515200, 0).Unix() * 1000,
	})
	if err != nil {
		t.Error("Binance GetSpotKline() error", err)
	}
}

func TestGetAveragePrice(t *testing.T) {
	t.Parallel()

	_, err := b.GetAveragePrice("BTCUSDT")
	if err != nil {
		t.Error("Binance GetAveragePrice() error", err)
	}
}

func TestGetPriceChangeStats(t *testing.T) {
	t.Parallel()

	_, err := b.GetPriceChangeStats("BTCUSDT")
	if err != nil {
		t.Error("Binance GetPriceChangeStats() error", err)
	}
}

func TestGetTickers(t *testing.T) {
	t.Parallel()

	_, err := b.GetTickers()
	if err != nil {
		t.Error("Binance TestGetTickers error", err)
	}
}

func TestGetLatestSpotPrice(t *testing.T) {
	t.Parallel()

	_, err := b.GetLatestSpotPrice("BTCUSDT")
	if err != nil {
		t.Error("Binance GetLatestSpotPrice() error", err)
	}
}

func TestGetBestPrice(t *testing.T) {
	t.Parallel()

	_, err := b.GetBestPrice("BTCUSDT")
	if err != nil {
		t.Error("Binance GetBestPrice() error", err)
	}
}

func TestQueryOrder(t *testing.T) {
	t.Parallel()

	_, err := b.QueryOrder("BTCUSDT", "", 1337)
	switch {
	case areTestAPIKeysSet() && err != nil:
		t.Error("QueryOrder() error", err)
	case !areTestAPIKeysSet() && err == nil && !mockTests:
		t.Error("QueryOrder() expecting an error when no keys are set")
	case mockTests && err != nil:
		t.Error("Mock QueryOrder() error", err)
	}
}

func TestOpenOrders(t *testing.T) {
	t.Parallel()

	_, err := b.OpenOrders("BTCUSDT")
	switch {
	case areTestAPIKeysSet() && err != nil:
		t.Error("OpenOrders() error", err)
	case !areTestAPIKeysSet() && err == nil && !mockTests:
		t.Error("OpenOrders() expecting an error when no keys are set")
	case mockTests && err != nil:
		t.Error("Mock OpenOrders() error", err)
	}
}

func TestAllOrders(t *testing.T) {
	t.Parallel()

	_, err := b.AllOrders("BTCUSDT", "", "")
	switch {
	case areTestAPIKeysSet() && err != nil:
		t.Error("AllOrders() error", err)
	case !areTestAPIKeysSet() && err == nil && !mockTests:
		t.Error("AllOrders() expecting an error when no keys are set")
	case mockTests && err != nil:
		t.Error("Mock AllOrders() error", err)
	}
}

// TestGetFeeByTypeOfflineTradeFee logic test
func TestGetFeeByTypeOfflineTradeFee(t *testing.T) {
	t.Parallel()

	var feeBuilder = setFeeBuilder()
	b.GetFeeByType(feeBuilder)
	if !areTestAPIKeysSet() {
		if feeBuilder.FeeType != exchange.OfflineTradeFee {
			t.Errorf("Expected %v, received %v", exchange.OfflineTradeFee, feeBuilder.FeeType)
		}
	} else {
		if feeBuilder.FeeType != exchange.CryptocurrencyTradeFee {
			t.Errorf("Expected %v, received %v", exchange.CryptocurrencyTradeFee, feeBuilder.FeeType)
		}
	}
}

func TestGetFee(t *testing.T) {
	t.Parallel()

	var feeBuilder = setFeeBuilder()

	if areTestAPIKeysSet() || mockTests {
		// CryptocurrencyTradeFee Basic
		if resp, err := b.GetFee(feeBuilder); resp != float64(0.1) || err != nil {
			t.Error(err)
			t.Errorf("GetFee() error. Expected: %f, Received: %f", float64(0), resp)
		}

		// CryptocurrencyTradeFee High quantity
		feeBuilder = setFeeBuilder()
		feeBuilder.Amount = 1000
		feeBuilder.PurchasePrice = 1000
		if resp, err := b.GetFee(feeBuilder); resp != float64(100000) || err != nil {
			t.Errorf("GetFee() error. Expected: %f, Received: %f", float64(100000), resp)
			t.Error(err)
		}

		// CryptocurrencyTradeFee IsMaker
		feeBuilder = setFeeBuilder()
		feeBuilder.IsMaker = true
		if resp, err := b.GetFee(feeBuilder); resp != float64(0.1) || err != nil {
			t.Errorf("GetFee() error. Expected: %f, Received: %f", float64(0.1), resp)
			t.Error(err)
		}

		// CryptocurrencyTradeFee Negative purchase price
		feeBuilder = setFeeBuilder()
		feeBuilder.PurchasePrice = -1000
		if resp, err := b.GetFee(feeBuilder); resp != float64(0) || err != nil {
			t.Errorf("GetFee() error. Expected: %f, Received: %f", float64(0), resp)
			t.Error(err)
		}
	}

	// CryptocurrencyWithdrawalFee Basic
	feeBuilder = setFeeBuilder()
	feeBuilder.FeeType = exchange.CryptocurrencyWithdrawalFee
	if resp, err := b.GetFee(feeBuilder); resp != float64(0.0005) || err != nil {
		t.Errorf("GetFee() error. Expected: %f, Received: %f", float64(0.0005), resp)
		t.Error(err)
	}

	// CyptocurrencyDepositFee Basic
	feeBuilder = setFeeBuilder()
	feeBuilder.FeeType = exchange.CyptocurrencyDepositFee
	if resp, err := b.GetFee(feeBuilder); resp != float64(0) || err != nil {
		t.Errorf("GetFee() error. Expected: %f, Received: %f", float64(0), resp)
		t.Error(err)
	}

	// InternationalBankDepositFee Basic
	feeBuilder = setFeeBuilder()
	feeBuilder.FeeType = exchange.InternationalBankDepositFee
	feeBuilder.FiatCurrency = currency.HKD
	if resp, err := b.GetFee(feeBuilder); resp != float64(0) || err != nil {
		t.Errorf("GetFee() error. Expected: %f, Received: %f", float64(0), resp)
		t.Error(err)
	}

	// InternationalBankWithdrawalFee Basic
	feeBuilder = setFeeBuilder()
	feeBuilder.FeeType = exchange.InternationalBankWithdrawalFee
	feeBuilder.FiatCurrency = currency.HKD
	if resp, err := b.GetFee(feeBuilder); resp != float64(0) || err != nil {
		t.Errorf("GetFee() error. Expected: %f, Received: %f", float64(0), resp)
		t.Error(err)
	}
}

func TestFormatWithdrawPermissions(t *testing.T) {
	t.Parallel()

	expectedResult := exchange.AutoWithdrawCryptoText + " & " + exchange.NoFiatWithdrawalsText
	withdrawPermissions := b.FormatWithdrawPermissions()
	if withdrawPermissions != expectedResult {
		t.Errorf("Expected: %s, Received: %s", expectedResult, withdrawPermissions)
	}
}

func TestGetActiveOrders(t *testing.T) {
	t.Parallel()

	var getOrdersRequest = order.GetOrdersRequest{
		Type: order.AnyType,
	}
	_, err := b.GetActiveOrders(&getOrdersRequest)
	if err == nil {
		t.Error("Expected: 'At least one currency is required to fetch order history'. received nil")
	}

	getOrdersRequest.Pairs = []currency.Pair{
		currency.NewPair(currency.LTC, currency.BTC),
	}

	_, err = b.GetActiveOrders(&getOrdersRequest)
	switch {
	case areTestAPIKeysSet() && err != nil:
		t.Error("GetActiveOrders() error", err)
	case !areTestAPIKeysSet() && err == nil && !mockTests:
		t.Error("GetActiveOrders() expecting an error when no keys are set")
	case mockTests && err != nil:
		t.Error("Mock GetActiveOrders() error", err)
	}
}

func TestGetOrderHistory(t *testing.T) {
	t.Parallel()

	var getOrdersRequest = order.GetOrdersRequest{
		Type: order.AnyType,
	}

	_, err := b.GetOrderHistory(&getOrdersRequest)
	if err == nil {
		t.Error("Expected: 'At least one currency is required to fetch order history'. received nil")
	}

	getOrdersRequest.Pairs = []currency.Pair{
		currency.NewPair(currency.LTC,
			currency.BTC)}

	_, err = b.GetOrderHistory(&getOrdersRequest)
	switch {
	case areTestAPIKeysSet() && err != nil:
		t.Error("GetOrderHistory() error", err)
	case !areTestAPIKeysSet() && err == nil && !mockTests:
		t.Error("GetOrderHistory() expecting an error when no keys are set")
	case mockTests && err != nil:
		t.Error("Mock GetOrderHistory() error", err)
	}
}

// Any tests below this line have the ability to impact your orders on the exchange. Enable canManipulateRealOrders to run them
// -----------------------------------------------------------------------------------------------------------------------------

func TestSubmitOrder(t *testing.T) {
	t.Parallel()

	if areTestAPIKeysSet() && !canManipulateRealOrders && !mockTests {
		t.Skip("API keys set, canManipulateRealOrders false, skipping test")
	}

	var orderSubmission = &order.Submit{
		Pair: currency.Pair{
			Delimiter: "_",
			Base:      currency.LTC,
			Quote:     currency.BTC,
		},
		Side:     order.Buy,
		Type:     order.Limit,
		Price:    1,
		Amount:   1000000000,
		ClientID: "meowOrder",
	}

	_, err := b.SubmitOrder(orderSubmission)
	switch {
	case areTestAPIKeysSet() && err != nil:
		t.Error("SubmitOrder() error", err)
	case !areTestAPIKeysSet() && err == nil && !mockTests:
		t.Error("SubmitOrder() expecting an error when no keys are set")
	case mockTests && err != nil:
		t.Error("Mock SubmitOrder() error", err)
	}
}

func TestCancelExchangeOrder(t *testing.T) {
	t.Parallel()

	if areTestAPIKeysSet() && !canManipulateRealOrders && !mockTests {
		t.Skip("API keys set, canManipulateRealOrders false, skipping test")
	}
	var orderCancellation = &order.Cancel{
		ID:            "1",
		WalletAddress: core.BitcoinDonationAddress,
		AccountID:     "1",
		Pair:          currency.NewPair(currency.LTC, currency.BTC),
	}

	err := b.CancelOrder(orderCancellation)
	switch {
	case areTestAPIKeysSet() && err != nil:
		t.Error("CancelExchangeOrder() error", err)
	case !areTestAPIKeysSet() && err == nil && !mockTests:
		t.Error("CancelExchangeOrder() expecting an error when no keys are set")
	case mockTests && err != nil:
		t.Error("Mock CancelExchangeOrder() error", err)
	}
}

func TestCancelAllExchangeOrders(t *testing.T) {
	t.Parallel()

	if areTestAPIKeysSet() && !canManipulateRealOrders && !mockTests {
		t.Skip("API keys set, canManipulateRealOrders false, skipping test")
	}
	var orderCancellation = &order.Cancel{
		ID:            "1",
		WalletAddress: core.BitcoinDonationAddress,
		AccountID:     "1",
		Pair:          currency.NewPair(currency.LTC, currency.BTC),
	}

	_, err := b.CancelAllOrders(orderCancellation)
	switch {
	case areTestAPIKeysSet() && err != nil:
		t.Error("CancelAllExchangeOrders() error", err)
	case !areTestAPIKeysSet() && err == nil && !mockTests:
		t.Error("CancelAllExchangeOrders() expecting an error when no keys are set")
	case mockTests && err != nil:
		t.Error("Mock CancelAllExchangeOrders() error", err)
	}
}

func TestGetAccountInfo(t *testing.T) {
	t.Parallel()

	_, err := b.UpdateAccountInfo()
	switch {
	case areTestAPIKeysSet() && err != nil:
		t.Error("GetAccountInfo() error", err)
	case !areTestAPIKeysSet() && err == nil && !mockTests:
		t.Error("GetAccountInfo() expecting an error when no keys are set")
	case mockTests && err != nil:
		t.Error("Mock GetAccountInfo() error", err)
	}
}

func TestModifyOrder(t *testing.T) {
	t.Parallel()

	_, err := b.ModifyOrder(&order.Modify{})
	if err == nil {
		t.Error("ModifyOrder() error cannot be nil")
	}
}

func TestWithdraw(t *testing.T) {
	t.Parallel()
	if areTestAPIKeysSet() && !canManipulateRealOrders && !mockTests {
		t.Skip("API keys set, canManipulateRealOrders false, skipping test")
	}

	withdrawCryptoRequest := withdraw.Request{
		Amount:      0,
		Currency:    currency.BTC,
		Description: "WITHDRAW IT ALL",
		Crypto: &withdraw.CryptoRequest{
			Address: core.BitcoinDonationAddress,
		},
	}

	_, err := b.WithdrawCryptocurrencyFunds(&withdrawCryptoRequest)
	switch {
	case areTestAPIKeysSet() && err != nil:
		t.Error("Withdraw() error", err)
	case !areTestAPIKeysSet() && err == nil && !mockTests:
		t.Error("Withdraw() expecting an error when no keys are set")
	case mockTests && err != nil:
		t.Error("Mock Withdraw() error", err)
	}
}

func TestWithdrawFiat(t *testing.T) {
	t.Parallel()

	_, err := b.WithdrawFiatFunds(&withdraw.Request{})
	if err != common.ErrFunctionNotSupported {
		t.Errorf("Expected '%v', received: '%v'", common.ErrFunctionNotSupported, err)
	}
}

func TestWithdrawInternationalBank(t *testing.T) {
	t.Parallel()

	_, err := b.WithdrawFiatFundsToInternationalBank(&withdraw.Request{})
	if err != common.ErrFunctionNotSupported {
		t.Errorf("Expected '%v', received: '%v'", common.ErrFunctionNotSupported, err)
	}
}

func TestGetDepositAddress(t *testing.T) {
	t.Parallel()

	_, err := b.GetDepositAddress(currency.BTC, "")
	switch {
	case areTestAPIKeysSet() && err != nil:
		t.Error("GetDepositAddress() error", err)
	case !areTestAPIKeysSet() && err == nil && !mockTests:
		t.Error("GetDepositAddress() error cannot be nil")
	case mockTests && err != nil:
		t.Error("Mock GetDepositAddress() error", err)
	}
}

func TestWSSubscriptionHandling(t *testing.T) {
	t.Parallel()
	pressXToJSON := []byte(`{
  "method": "SUBSCRIBE",
  "params": [
    "btcusdt@aggTrade",
    "btcusdt@depth"
  ],
  "id": 1
}`)
	err := b.wsHandleData(pressXToJSON)
	if err != nil {
		t.Error(err)
	}
}

func TestWSUnsubscriptionHandling(t *testing.T) {
	pressXToJSON := []byte(`{
  "method": "UNSUBSCRIBE",
  "params": [
    "btcusdt@depth"
  ],
  "id": 312
}`)
	err := b.wsHandleData(pressXToJSON)
	if err != nil {
		t.Error(err)
	}
}

func TestWsOrderUpdateHandling(t *testing.T) {
	t.Parallel()
	pressXToJSON := []byte(`{
	  "e": "executionReport",        
	  "E": 1499405658658,            
	  "s": "BTCUSDT",                 
	  "c": "mUvoqJxFIILMdfAW5iGSOW", 
	  "S": "BUY",                    
	  "o": "LIMIT",                  
	  "f": "GTC",                    
	  "q": "1.00000000",             
	  "p": "0.10264410",             
	  "P": "0.00000000",             
	  "F": "0.00000000",             
	  "g": -1,                       
	  "C": null,                     
	  "x": "NEW",                    
	  "X": "NEW",                    
	  "r": "NONE",                   
	  "i": 4293153,                  
	  "l": "0.00000000",             
	  "z": "0.00000000",             
	  "L": "0.00000000",             
	  "n": "0",                      
	  "N": null,                     
	  "T": 1499405658657,            
	  "t": -1,                       
	  "I": 8641984,                  
	  "w": true,                     
	  "m": false,                    
	  "M": false,                    
	  "O": 1499405658657,            
	  "Z": "0.00000000",             
	  "Y": "0.00000000",              
	  "Q": "0.00000000"              
	}`)
	err := b.wsHandleData(pressXToJSON)
	if err != nil {
		t.Error(err)
	}
}

func TestWsOutboundAccountPosition(t *testing.T) {
	t.Parallel()
	pressXToJSON := []byte(`{
  "e": "outboundAccountPosition", 
  "E": 1564034571105,             
  "u": 1564034571073,             
  "B": [                          
    {
      "a": "ETH",                 
      "f": "10000.000000",        
      "l": "0.000000"             
    }
  ]
}`)
	err := b.wsHandleData(pressXToJSON)
	if err != nil {
		t.Error(err)
	}
}

func TestWsTickerUpdate(t *testing.T) {
	t.Parallel()
	pressXToJSON := []byte(`{"stream":"btcusdt@ticker","data":{"e":"24hrTicker","E":1580254809477,"s":"BTCUSDT","p":"420.97000000","P":"4.720","w":"9058.27981278","x":"8917.98000000","c":"9338.96000000","Q":"0.17246300","b":"9338.03000000","B":"0.18234600","a":"9339.70000000","A":"0.14097600","o":"8917.99000000","h":"9373.19000000","l":"8862.40000000","v":"72229.53692000","q":"654275356.16896672","O":1580168409456,"C":1580254809456,"F":235294268,"L":235894703,"n":600436}}`)
	err := b.wsHandleData(pressXToJSON)
	if err != nil {
		t.Error(err)
	}
}

func TestWsKlineUpdate(t *testing.T) {
	t.Parallel()
	pressXToJSON := []byte(`{"stream":"btcusdt@kline_1m","data":{
	  "e": "kline",     
	  "E": 123456789,   
	  "s": "BNBBTC",    
	  "k": {
		"t": 123400000, 
		"T": 123460000, 
		"s": "BNBBTC",  
		"i": "1m",      
		"f": 100,       
		"L": 200,       
		"o": "0.0010",  
		"c": "0.0020",  
		"h": "0.0025",  
		"l": "0.0015",  
		"v": "1000",    
		"n": 100,       
		"x": false,     
		"q": "1.0000",  
		"V": "500",     
		"Q": "0.500",   
		"B": "123456"   
	  }
	}}`)
	err := b.wsHandleData(pressXToJSON)
	if err != nil {
		t.Error(err)
	}
}

func TestWsTradeUpdate(t *testing.T) {
	t.Parallel()
	pressXToJSON := []byte(`{"stream":"btcusdt@trade","data":{
	  "e": "trade",     
	  "E": 123456789,   
	  "s": "BNBBTC",    
	  "t": 12345,       
	  "p": "0.001",     
	  "q": "100",       
	  "b": 88,          
	  "a": 50,          
	  "T": 123456785,   
	  "m": true,        
	  "M": true         
	}}`)
	err := b.wsHandleData(pressXToJSON)
	if err != nil {
		t.Error(err)
	}
}

func TestWsDepthUpdate(t *testing.T) {
	t.Parallel()
	pressXToJSON := []byte(`{"stream":"btcusdt@depth","data":{
	  "e": "depthUpdate", 
	  "E": 123456789,     
	  "s": "BTCUSDT",      
	  "U": 157,           
	  "u": 160,           
	  "b": [              
		[
		  "0.0024",       
		  "10"            
		]
	  ],
	  "a": [              
		[
		  "0.0026",       
		  "100"           
		]
	  ]
	}}`)

	err := b.wsHandleData(pressXToJSON)
	if err.Error() != "Binance - UpdateLocalCache error: ob.Base could not be found for Exchange Binance CurrencyPair: BTC-USDT AssetType: spot" {
		t.Error(err)
	}
}

func TestWsBalanceUpdate(t *testing.T) {
	t.Parallel()
	pressXToJSON := []byte(`{
  "e": "balanceUpdate",         
  "E": 1573200697110,           
  "a": "BTC",                   
  "d": "100.00000000",          
  "T": 1573200697068            
}`)
	err := b.wsHandleData(pressXToJSON)
	if err != nil {
		t.Error(err)
	}
}

func TestWsOCO(t *testing.T) {
	t.Parallel()
	pressXToJSON := []byte(`{
  "e": "listStatus",                
  "E": 1564035303637,               
  "s": "ETHBTC",                    
  "g": 2,                           
  "c": "OCO",                       
  "l": "EXEC_STARTED",              
  "L": "EXECUTING",                 
  "r": "NONE",                      
  "C": "F4QN4G8DlFATFlIUQ0cjdD",    
  "T": 1564035303625,               
  "O": [                            
    {
      "s": "ETHBTC",                
      "i": 17,                      
      "c": "AJYsMjErWJesZvqlJCTUgL" 
    },
    {
      "s": "ETHBTC",
      "i": 18,
      "c": "bfYPSQdLoqAJeNrOr9adzq"
    }
  ]
}`)
	err := b.wsHandleData(pressXToJSON)
	if err != nil {
		t.Error(err)
	}
}

func TestGetWsAuthStreamKey(t *testing.T) {
	key, err := b.GetWsAuthStreamKey()
	switch {
	case mockTests && err != nil,
		!mockTests && areTestAPIKeysSet() && err != nil:
		t.Fatal(err)
	case !mockTests && !areTestAPIKeysSet() && err == nil:
		t.Fatal("Expected error")
	}

	if key == "" {
		t.Error("Expected key")
	}
}

func TestMaintainWsAuthStreamKey(t *testing.T) {
	err := b.MaintainWsAuthStreamKey()
	switch {
	case mockTests && err != nil,
		!mockTests && areTestAPIKeysSet() && err != nil:
		t.Fatal(err)
	case !mockTests && !areTestAPIKeysSet() && err == nil:
		t.Fatal("Expected error")
	}
}

func TestExecutionTypeToOrderStatus(t *testing.T) {
	type TestCases struct {
		Case   string
		Result order.Status
	}
	testCases := []TestCases{
		{Case: "NEW", Result: order.New},
		{Case: "CANCELLED", Result: order.Cancelled},
		{Case: "REJECTED", Result: order.Rejected},
		{Case: "TRADE", Result: order.PartiallyFilled},
		{Case: "EXPIRED", Result: order.Expired},
		{Case: "LOL", Result: order.UnknownStatus},
	}
	for i := range testCases {
		result, _ := stringToOrderStatus(testCases[i].Case)
		if result != testCases[i].Result {
			t.Errorf("Exepcted: %v, received: %v", testCases[i].Result, result)
		}
	}
}

func TestGetHistoricCandles(t *testing.T) {
	currencyPair := currency.NewPairFromString("BTCUSDT")
	startTime := time.Unix(1546300800, 0)
	end := time.Unix(1577836799, 0)
	_, err := b.GetHistoricCandles(currencyPair, asset.Spot, startTime, end, kline.OneDay)
	if err != nil {
		t.Fatal(err)
	}

	_, err = b.GetHistoricCandles(currencyPair, asset.Spot, startTime, end, kline.Interval(time.Hour*7))
	if err == nil {
		t.Fatal("unexpected result")
	}
}

func TestGetHistoricCandlesEx(t *testing.T) {
	currencyPair := currency.NewPairFromString("BTCUSDT")
	startTime := time.Unix(1546300800, 0)
	end := time.Unix(1577836799, 0)
	_, err := b.GetHistoricCandlesEx(currencyPair, asset.Spot, startTime, end, kline.OneDay)
	if err != nil {
		t.Fatal(err)
	}
	_, err = b.GetHistoricCandlesEx(currencyPair, asset.Spot, startTime, end, kline.Interval(time.Hour*7))
	if err == nil {
		t.Fatal("unexpected result")
	}
}
