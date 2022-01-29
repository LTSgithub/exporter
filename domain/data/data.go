package data

import (
	"database/sql"
	"github.com/go-logr/logr"
	"github.com/lits01/xiaozhan/pkg/configs"
	"github.com/lits01/xiaozhan/repositories/generated"
)

type Data struct {
	config configs.Configuration
	log *logr.Logger
	query   *generated.Queries
	AlertConfig *AlertConfig
}

func NewData(config configs.Configuration,log *logr.Logger, db *sql.DB ) *Data {
	return &Data{
		log: log,
		config: config,
		query: generated.New(db),
		AlertConfig: NewAlertConfig(config,log,db),
	}
}

