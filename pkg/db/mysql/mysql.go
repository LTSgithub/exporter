package mysql

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/lits01/xiaozhan/pkg/configs"

	_ "github.com/go-sql-driver/mysql"
	"github.com/lits01/xiaozhan/pkg/db/mysqx"
)

func readDBConfig(cfg configs.Configuration) (options mysqx.Options, err error) {
	port, err := cfg.GetInt("DB.port", 3306)
	if err != nil {
		err = fmt.Errorf("invalid port: %d, error: %w", port, err)
		return
	}

	return mysqx.Options{
		User:     cfg.GetString("DB.user", "root"),
		Password: cfg.GetString("DB.password", "dev"),
		Host:     cfg.GetString("DB.host", "localhost"),
		Port:     port,
		Database: cfg.GetString("DB.database", ""),
	}, nil
}

func NewDBGetter(config configs.Configuration) (mysqx.DBGetter, error) {
	dbGetter, _, err := OpenMysqlConnection(config)
	if err != nil {
		return nil, err
	}

	return dbGetter, nil
}

func NewSqlDb(config configs.Configuration) (*sql.DB, error) {

	_, mysqlConn, err := OpenMysqlConnection(config)
	if err != nil {
		return nil, err
	}

	return mysqlConn.GetDB(), nil
}

func OpenMysqlConnection(cfg configs.Configuration) (mysqx.DBGetter, *mysqx.Connection, error) {
	dbOptions, err := readDBConfig(cfg)
	if err != nil {
		return nil, nil, err
	}
	databases := strings.Split(dbOptions.Database, ",")
	connMap := mysqx.NewMultiConnectionMap()
	var defaultConn *mysqx.Connection
	for _, database := range databases {
		dbOptions.Database = database
		conn, err := mysqx.NewConnectionWithOptions(dbOptions)
		if err != nil {
			return nil, nil, err
		}
		if err = conn.Ping(5 * time.Second); err != nil {
			return nil, nil, fmt.Errorf("failed to connect mysql, %s:%d",
				dbOptions.Host,
				dbOptions.Port)
		}
		if defaultConn == nil {
			defaultConn = conn
		}
		connMap.Add(database, conn)
	}
	return connMap, defaultConn, nil
}

func OpenMysqlConnection1(dbOptions mysqx.Options) (*mysqx.Connection, error) {

	conn, err := mysqx.NewConnectionWithOptions(dbOptions)
	if err != nil {
		return nil, err
	}

	if err = conn.Ping(5 * time.Second); err != nil {
		return nil, fmt.Errorf("failed to connect mysql, %s:%d", dbOptions.Host, dbOptions.Port)
	}

	return conn, nil
}
