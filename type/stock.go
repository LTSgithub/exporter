package typing

type GetStockRequest struct {
	Code string `json:"code"`
}
type GetStockResponse struct {
	Stock
}

type TV struct {
	Time  int64
	Price float64
}

type Stock struct {
	Code  string  `json:"code"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Time  int64   `json:"time"`
}
