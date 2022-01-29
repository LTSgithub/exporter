package model

const (
	Up   = "up"
	Down = "down"
)

type AlertConfigDeleteRequest struct {
	Id string
}

type AlertConfigDeleteResponse struct {
}

type AlertConfigUpdateRequest struct {
	Id         string  `json:"id"`
	NotifyType string  `json:"notify_type"`
	Deadline   int64   `json:"deadline"`
	Price      float64 `json:"price"`
	Desc       string  `json:"desc"`
}
type AlertConfigUpdateResponse struct {
}

type AlertConfigCreateRequest struct {
	UserId     string
	StockCode  string
	NotifyType string
	Deadline   int64
	Price      float64
	Desc       string
}

type AlertConfigCreateResponse struct {
}

type AlertConfigListRequest struct {
	UserId   string
	Page     int
	PageSize int
}

type AlertConfigListResponse struct {
	Total int                `json:"total"`
	Items []*AlertConfigInfo `json:"items"`
}

type AlertConfigInfo struct {
	Id         string  `json:"id"`
	UserId     string  `json:"user_id"`
	StockCode  string  `json:"stock_code"`
	NotifyType string  `json:"type"`
	Price      float64 `json:"price"`
	Desc       string  `json:"desc"`
	Deadline   int64   `json:"deadline"`
	CreateTime int64   `json:"create_time"`
}
