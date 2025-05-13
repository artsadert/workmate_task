package response

type StatsResponse struct {
	Account string  `json:"account"`
	Trades  int     `json:"trades"`
	Profit  float64 `json:"profit"`
}
