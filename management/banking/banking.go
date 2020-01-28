package banking

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
	"github.com/thrasher-corp/gocryptotrader/currency"
)

// GetBankAccountByID Returns a bank account based on its ID
func GetBankAccountByID(id string) (*Account, error) {
	m.Lock()
	defer m.Unlock()

	for x := range Accounts {
		if strings.EqualFold(Accounts[x].ID, id) {
			return &Accounts[x], nil
		}
	}
	return nil, fmt.Errorf(ErrBankAccountNotFound, id)
}

func (b *Account) ExchangeSupported(exchange string) bool {
	exchList := strings.Split(b.SupportedExchanges, ",")
	for x := range exchList {
		if exchList[x] == exchange {
			return true
		}
	}
	return false
}

// Validate validates bank account settings
func (b *Account) Validate() error {
	if b == nil {
		return errors.New("nope")
	}

	if b.BankName == "" ||
		b.BankAddress == "" ||
		b.BankPostalCode == "" ||
		b.BankPostalCity == "" ||
		b.BankCountry == "" ||
		b.AccountName == "" ||
		b.SupportedCurrencies == "" {
		return fmt.Errorf(
			"banking details for %s is enabled but variables not set correctly",
			b.BankName)
	}

	if b.SupportedExchanges == "" {
		b.SupportedExchanges = "ALL"
	}

	if strings.Contains(strings.ToUpper(
		b.SupportedCurrencies),
		currency.AUD.String()) {
		if b.BSBNumber == "" ||
			b.SWIFTCode == "" {
			return fmt.Errorf(
				"banking details for %s is enabled but BSB/SWIFT values not set",
				b.BankName)
		}
	} else {
		// Either IBAN or SWIFT code is OK
		if b.IBAN == "" && b.SWIFTCode == "" {
			return fmt.Errorf(
				"banking details for %s is enabled but SWIFT/IBAN values not set",
				b.BankName)
		}
	}
	return nil
}

func (b *Account) ValidateForWithdrawal(cur currency.Code) (err []string) {
	if b == nil {
		return []string{"Account cannot not be nil"}
	}

	if b.AccountNumber == "" {
		err = append(err, "Bank Account Number cannot be empty")
	}

	if cur == currency.AUD {
		if b.BSBNumber == "" {
			err = append(err, "BSB must be set for AUD values not set")
		}
	} else {
		if b.IBAN == "" && b.SWIFTCode == "" {
			err = append(err, "IBAN/SWIFT values not set")
		}
	}

	return
}