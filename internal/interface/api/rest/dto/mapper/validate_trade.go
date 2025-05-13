package mapper

import (
	"errors"
	"regexp"
	"strings"

	trade_request "gitlab.com/digineat/go-broker-test/internal/interface/api/rest/dto/request"
)

var correct_symbol = regexp.MustCompile(`^[A-Z]{6}$`)

func ValidateTrade(trade trade_request.TradeRequest) error {
	if strings.TrimSpace(trade.Account) == "" {
		return errors.New("Account field must not be empty")
	}
	if !correct_symbol.MatchString(trade.Symbol) {
		return errors.New("Symbol must be 6 uppercase letters")
	}
	if trade.Volume <= 0 {
		return errors.New("Volume must greater than 0")
	}
	if trade.Open <= 0 {
		return errors.New("Open price must be greater than 0")
	}
	if trade.Close <= 0 {
		return errors.New("Close price must be greater than 0")
	}
	side := strings.ToLower(trade.Side)
	if side != "buy" && side != "sell" {
		return errors.New("Side must be 'buy' or 'sell'")
	}
	trade.Side = side
	return nil
}
