package engin

import (
	"context"
	"database/sql"
	"github.com/go-logr/logr"
	"github.com/lits01/xiaozhan/domain/common"
	"github.com/lits01/xiaozhan/domain/engin/cache"
	"github.com/lits01/xiaozhan/domain/engin/sina"
	_type "github.com/lits01/xiaozhan/domain/engin/type"
	"github.com/lits01/xiaozhan/pkg/configs"
	time2 "github.com/lits01/xiaozhan/pkg/time"
	"github.com/lits01/xiaozhan/repositories/generated"
	"github.com/pkg/errors"
	"time"
)

type Engin struct {
	config configs.Configuration
	log    *logr.Logger
	query  *generated.Queries
	sina   *sina.Sina
	cache  *cache.Cache
}

func NewEngin(conf configs.Configuration, log *logr.Logger, db *sql.DB) *Engin {
	return &Engin{
		config: conf,
		log:    log,
		query:  generated.New(db),
		sina:   sina.NewSina(conf, log),
		cache: cache.NewCache(),
	}
}

func (m *Engin) Run() error {
	//m.syncStockList()
	m.syncStockPrice()
	return nil
}

func (m *Engin) syncStockPrice() error {
	t := time.Now()

	for {
		if err := func() error {
			ctx, cancal := common.NewDefaultContext()
			defer cancal()
			notUpdateStockList, err := m.query.GetNotUpdateStockList(ctx, time2.GetDateTimeString(t))
			if err != nil {
				return errors.Wrap(err, "查询数据失败")
			}
			if len(notUpdateStockList) == 0 {
				t = time.Now()
				m.log.Info("更新完一轮数据")
				return nil
			}

			stockInfo, err := m.sina.GetStocksPrice(ctx, notUpdateStockList)
			if err != nil {
				return errors.Wrap(err, "获取新浪数据失败")
			}

			for _, v := range stockInfo {
				if err := m.updateStockInfo(v.Code,"",v.Sprice,t);err != nil {
					m.log.Error(err, "")
				}
			}

			return nil
		}(); err != nil {
			m.log.Error(err, "更新数据库失败")
		}else {
			m.log.Info("更新完99条数据")
		}

		time.Sleep(5 * time.Second)
	}

	return nil
}

func (m * Engin)updateStockInfo(code string,name string,price float64,t time.Time) error  {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	arg := generated.UpdateStockStatusParams{
		Sprice:     price,
		Name:       name,
		UpdateTime: time2.GetDateTimeString(t),
		Code:       code,
	}

	if err := m.query.UpdateStockStatus(ctx, arg); err != nil {
		return errors.Wrap(err, "更新数据库失败")
	}

	data := _type.TV{
		Time :t.Unix(),
		Value :price,
	}
	m.cache.SetRealTime(code,&data)

	return nil
}

func (m *Engin) syncStockList() error {
	codeList := m.sina.GetStockList()

	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Second)
	defer cancel()
	dbStockList, err := m.query.GetStockStatusList(ctx)
	if err != nil {
		return errors.Wrap(err, "查询数据失败")
	}
	dbStockMap := map[string]*generated.StockStatus{}
	for _, stock := range dbStockList {
		stock := stock
		dbStockMap[stock.Code] = &stock
	}

	t := time.Now()
	for _, code := range codeList {
		if _, ok := dbStockMap[code]; ok {
			continue
		}

		arg := generated.InsertStockStatusParams{
			Code:       code,
			CreateTime: time2.GetDateTimeString(t),
		}
		if err := m.query.InsertStockStatus(ctx, arg); err != nil {
			return nil
		}
	}

	return nil
}
