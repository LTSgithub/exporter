//go:build wireinject
// +build wireinject

package server

import (
	"github.com/google/wire"
	"github.com/lits01/xiaozhan/domain/engin"
	"github.com/lits01/xiaozhan/pkg/db/mysql"
	logging2 "github.com/lits01/xiaozhan/pkg/logging"
)

func InitServer() (*Server, error) {
	wire.Build(
		NewServer,
		NewConfig,
		mysql.NewSqlDb,
		logging2.NewLogger,
		engin.NewEngin,
	)

	return &Server{}, nil
}
