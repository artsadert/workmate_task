package mapper

import (
	trade_request "gitlab.com/digineat/go-broker-test/internal/interface/api/rest/dto/request"
	"testing"
)

func TestTradeValidate(t *testing.T) {
	tests := []struct {
		name    string
		trade   trade_request.TradeRequest
		wantErr bool
	}{
		{
			name: "valid buy trade",
			trade: trade_request.TradeRequest{
				Account: "acc1",
				Symbol:  "EURUSD",
				Volume:  1.0,
				Open:    1.1,
				Close:   1.2,
				Side:    "buy",
			},
			wantErr: false,
		},
		{
			name: "valid sell trade",
			trade: trade_request.TradeRequest{
				Account: "acc1",
				Symbol:  "EURUSD",
				Volume:  1.0,
				Open:    1.1,
				Close:   1.2,
				Side:    "sell",
			},
			wantErr: false,
		},
		{
			name: "empty account",
			trade: trade_request.TradeRequest{
				Account: "",
				Symbol:  "EURUSD",
				Volume:  1.0,
				Open:    1.1,
				Close:   1.2,
				Side:    "buy",
			},
			wantErr: true,
		},
		{
			name: "invalid symbol",
			trade: trade_request.TradeRequest{
				Account: "acc1",
				Symbol:  "eurusd",
				Volume:  1.0,
				Open:    1.1,
				Close:   1.2,
				Side:    "buy",
			},
			wantErr: true,
		},
		{
			name: "zero volume",
			trade: trade_request.TradeRequest{
				Account: "acc1",
				Symbol:  "EURUSD",
				Volume:  0,
				Open:    1.1,
				Close:   1.2,
				Side:    "buy",
			},
			wantErr: true,
		},
		{
			name: "zero open",
			trade: trade_request.TradeRequest{
				Account: "acc1",
				Symbol:  "EURUSD",
				Volume:  1.0,
				Open:    0,
				Close:   1.2,
				Side:    "buy",
			},
			wantErr: true,
		},
		{
			name: "zero close",
			trade: trade_request.TradeRequest{
				Account: "acc1",
				Symbol:  "EURUSD",
				Volume:  1.0,
				Open:    1.1,
				Close:   0,
				Side:    "buy",
			},
			wantErr: true,
		},
		{
			name: "invalid side",
			trade: trade_request.TradeRequest{
				Account: "acc1",
				Symbol:  "EURUSD",
				Volume:  1.0,
				Open:    1.1,
				Close:   1.2,
				Side:    "hold",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateTrade(tt.trade)
			if (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
