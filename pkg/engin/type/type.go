package _type

type StockInfo struct {
	Code  string
	Name  string
	Price float64
	Err   error
}

type SinaStockListRequest struct {
	P   int `json:"p"`
	SrP int `json:"sr_p"`
}

type TV struct {
	Time  int64
	Value float64
}