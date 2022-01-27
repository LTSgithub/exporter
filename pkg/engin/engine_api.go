package engin

import (
	"context"

	typing "github.com/lits01/xiaozhan/type"
)

func (m *Engin) GetStock(ctx context.Context, code string) (*typing.Stock, error) {
	resp := &typing.Stock{}

	data, err := m.cache.GetRealTime(code)
	if err != nil {
		return nil, err
	}
	resp.Code = code
	resp.Time = data.Time
	resp.Price = data.Price

	return resp, nil
}
