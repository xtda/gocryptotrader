package engine

import (
	"errors"
	"time"

	"github.com/gofrs/uuid"
	"github.com/thrasher-corp/gocryptotrader/withdraw"
	withdrawal "github.com/thrasher-corp/gocryptotrader/database/repository/withdraw"
)

func SubmitWithdrawal(exchName string, req *withdraw.Request) (*withdraw.Response, error) {
	if req == nil {
		return nil, errors.New("crypto withdraw request param is nil")
	}

	err := withdraw.Valid(req)
	if err != nil {
		return nil, err
	}
	exch := GetExchangeByName(exchName)
	if exch == nil {
		return nil, ErrExchangeNotFound
	}

	var exchID string
	if req.Type == withdraw.Crypto {
		exchID, err = exch.WithdrawCryptocurrencyFunds(req)
		if err != nil {
			return nil, err
		}
	}

	if req.Type == withdraw.Fiat {
		exchID, err = exch.WithdrawFiatFunds(req)
		if err != nil {
			return nil, err
		}
	}
	id, _ := uuid.NewV4()
	resp := &withdraw.Response{
		ID:             id,
		ExchangeID:    	exchID,
		RequestDetails: req,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	withdrawal.Event(resp)

	return resp, nil
}



// RequestByID returns a withdrawal request by ID
func RequestByID(id uuid.UUID) (*withdraw.Response, error) {
	return nil, nil
}

func RequestsByExchange(exchange string, limit int) ([]withdraw.Response, error) {
	return nil, nil
}

func RequestsByDate(start, end time.Time) ([]withdraw.Response, error) {
	return nil, nil
}
