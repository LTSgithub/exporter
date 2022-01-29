package engin

import (
	"context"

	model "github.com/lits01/xiaozhan/pkg/engin/model"
)

func (m *Engin) GetRealTimeList(ctx context.Context, req *model.GetRealTimeListRequest) (*model.GetRealTimeListResponse, error) {
	resp := &model.GetRealTimeListResponse{}

	return resp, nil
}

func (m *Engin) GetStockDetail(ctx context.Context, code string) (*model.StockDetail, error) {
	resp := &model.StockDetail{
		Status: &model.StockStatus{},
	}

	data, err := m.cache.GetRealTime(code)
	if err != nil {
		return nil, err
	}
	resp.Code = code
	resp.Status.DaysPrice

	return resp, nil
}

func (m *Engin) GetStockList(ctx context.Context) {

}
