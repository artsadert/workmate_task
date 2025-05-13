package sqlite

type Trade struct {
	ID        int64   `db:"id"`
	Account   string  `db:"account"`
	Symbol    string  `db:"symbol"`
	Volume    float64 `db:"volume"`
	Open      float64 `db:"open"`
	Close     float64 `db:"close"`
	Side      string  `db:"side"`
	Processed bool    `db:"processed"`
}

type Stats struct {
	Account string  `json:"account"`
	Trades  int     `json:"trades"`
	Profit  float64 `json:"profit"`
}
