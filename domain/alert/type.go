package alert_config

const (
	Up   = "up"
	Down = "down"
)

type AlertDetail struct {
	Id         string  `json:"id"`
	UserId     string  `json:"user_id"`
	StockCode  string  `json:"stock_code"`
	NotifyType string  `json:"type"`
	Price      float64 `json:"price"`
	Desc       string  `json:"desc"`
}
