package cache

import (
	"sync"

	typing "github.com/lits01/xiaozhan/type"
	"github.com/pkg/errors"
)

type Cache struct {
	status    map[string]*Status
	realTimes map[string]*typing.TV
	days      map[string][]*typing.TV
	weeks     map[string][]*typing.TV
	month     map[string][]*typing.TV
	lock      sync.RWMutex
}

func NewCache() *Cache {
	return &Cache{
		status:    map[string]*Status{},
		realTimes: map[string]*typing.TV{},
		days:      map[string][]*typing.TV{},
		weeks:     map[string][]*typing.TV{},
		month:     map[string][]*typing.TV{},
	}
}

func (m *Cache) GetCodeListByTime(Time int64, count int) []string {
	m.lock.Lock()
	defer m.lock.Unlock()
	var resp []string

	if count <= 0 {
		count = 10
	}

	for k, v := range m.status {
		if len(resp) >= count {
			return resp
		}
		if v.UpdateTime != Time {
			resp = append(resp, k)
		}
	}

	return resp
}

func (m *Cache) UpdateUpdateTime(codes []string, Time int64) {
	m.lock.Lock()
	defer m.lock.Unlock()

	for _, v := range codes {
		status, ok := m.status[v]
		if ok {
			status.UpdateTime = Time
		}
	}
}

func (m *Cache) GetRealTime(code string) (*typing.TV, error) {
	m.lock.RLock()
	defer m.lock.RUnlock()

	data, ok := m.realTimes[code]
	if !ok {
		return nil, errors.New("数据不存在")
	}

	return data, nil
}

func (m *Cache) SetRealTime(code string, data *typing.TV) {
	m.lock.Lock()
	defer m.lock.Unlock()

	m.realTimes[code] = data
}
