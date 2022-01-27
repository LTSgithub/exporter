package engin

import (
	"context"
	"database/sql"
	"time"

	"github.com/go-logr/logr"
	"github.com/lits01/xiaozhan/pkg/configs"
	"github.com/lits01/xiaozhan/pkg/engin/cache"
	"github.com/lits01/xiaozhan/pkg/engin/webdata"
	time2 "github.com/lits01/xiaozhan/pkg/time"
	"github.com/lits01/xiaozhan/repositories/generated"
	typing "github.com/lits01/xiaozhan/type"
	"github.com/pkg/errors"
)

type Engin struct {
	config  configs.Configuration
	log     *logr.Logger
	query   *generated.Queries
	webdata *webdata.Webdata
	cache   *cache.Cache
}

func NewEngin(conf configs.Configuration, log *logr.Logger, db *sql.DB) *Engin {
	return &Engin{
		config:  conf,
		log:     log,
		query:   generated.New(db),
		webdata: webdata.NewWebdata(conf, log),
		cache:   cache.NewCache(),
	}
}

func (m *Engin) Run() error {
	go m.syncStockList()

	go m.syncRealtimeInfo()

	go m.syncRealtimeInfo()

	return nil
}

func (m *Engin) syncHistoryInfo() error {

	return nil
}

func (m *Engin) syncRealtimeInfo() error {
	t := time.Now()

	for {
		if err := func() error {
			ctx, cancal := NewDefaultContext()
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

			stockInfo, err := m.webdata.GetStocksPrice(ctx, notUpdateStockList)
			if err != nil {
				return errors.Wrap(err, "获取新浪数据失败")
			}

			for _, v := range stockInfo {
				if err := m.updateStockInfo(v.Code, "", v.Sprice, t); err != nil {
					m.log.Error(err, "")
				}
			}

			return nil
		}(); err != nil {
			m.log.Error(err, "更新数据库失败")
		} else {
			m.log.Info("更新完99条数据")
		}

		time.Sleep(5 * time.Second)
	}

	return nil
}

func (m *Engin) updateStockInfo(code string, name string, price float64, t time.Time) error {
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

	data := typing.TV{
		Time:  t.Unix(),
		Price: price,
	}
	m.cache.SetRealTime(code, &data)

	return nil
}

func (m *Engin) syncStockList() error {
	codeList := m.webdata.GetStockList()

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
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
