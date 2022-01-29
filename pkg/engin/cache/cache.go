package cache

import (
	"sync"

	"github.com/lits01/xiaozhan/pkg/engin/model"

	"github.com/pkg/errors"
)

type Cache struct {
	lock sync.RWMutex

	realTimes map[string]*model.Status
	days      map[string]*model.Status
	weeks     map[string]*model.Status
	months    map[string]*model.Status
}

func NewCache() *Cache {
	return &Cache{
		realTimes: map[string]*model.Status{},
		days:      map[string]*model.Status{},
		weeks:     map[string]*model.Status{},
		months:    map[string]*model.Status{},
	}
}

func (m *Cache) GetRealTimeCodeListByTime(Time int64, limit int) []string {
	m.lock.Lock()
	defer m.lock.Unlock()
	var resp []string

	if limit <= 0 {
		limit = 1
	}

	for k, v := range m.realTimes {
		if len(resp) >= limit {
			return resp
		}
		if v.UpdateTime != Time {
			resp = append(resp, k)
		}
	}

	return resp
}

func (m *Cache) GetRealTime(code string) (*model.Status, error) {
	m.lock.RLock()
	defer m.lock.RUnlock()

	data, ok := m.realTimes[code]
	if !ok {
		return nil, errors.New("数据不存在")
	}

	return data, nil
}

func (m *Cache) SetRealTime(code string, t int64, price float32) {
	m.lock.Lock()
	defer m.lock.Unlock()

	data := &model.Status{
		UpdateTime: t,
		Price:      price,
	}

	m.realTimes[code] = data
}
