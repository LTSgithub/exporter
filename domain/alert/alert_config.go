package alert_config

import (
	"context"
	"github.com/lits01/xiaozhan/repositories/generated"
)

type AlertConfig struct {
	query *generated.Queries
}

func NewAlertConfig() *AlertConfig {
	return &AlertConfig{}
}

func (m * AlertConfig)(ctx context.Context, alertConfig * generated.AlertConfig  ) error {




	return nil
}

func (m *AlertConfig) List() (map[string]*AlertDetail, error) {
	resp := map[string]*AlertDetail{}
	var list []*AlertDetail

	for _, v := range list {
		resp[v.Id] = v
	}

	return resp, nil
}
