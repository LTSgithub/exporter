package alert_config

import (
	"context"
	"database/sql"

	"github.com/lits01/xiaozhan/pkg/engin"

	"github.com/go-logr/logr"
	"github.com/lits01/xiaozhan/pkg/configs"
	"github.com/lits01/xiaozhan/repositories/generated"
)

type Alertd struct {
	config configs.Configuration
	log    *logr.Logger
	query  *generated.Queries
	engin  *engin.Engin
}

func NewAlertConfig(config configs.Configuration, log *logr.Logger, db *sql.DB) *Alertd {
	return &Alertd{
		config: config,
		log:    log,
		query:  generated.New(db),
	}
}

func (m *Alertd) Run() error {

	return nil
}

func (m *Alertd) getAlertConfigList(ctx context.Context, T int64) ([]*generated.AlertConfig, error) {
	resp := []*generated.AlertConfig{}

	return resp, nil
}
