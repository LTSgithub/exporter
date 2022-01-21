package typing

type GetStockRequest struct {
	Code string `json:"code"`
}
type GetStockResponse struct {
	Stock
}

type Stock struct {
	Code string `json:"code"`
	Name string `json:"name"`
	Price float64 `json:"price"`
}
