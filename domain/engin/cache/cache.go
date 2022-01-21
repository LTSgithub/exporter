package cache

import (
	_type "github.com/lits01/xiaozhan/domain/engin/type"
	"github.com/pkg/errors"
	"sync"
)

type Cache struct {
	RealTimes map[string]*_type.TV
	Days      map[string][]*_type.TV
	lock      sync.RWMutex
}

func NewCache() *Cache {
	return &Cache{
		RealTimes: map[string]*_type.TV{},
		Days:      map[string][]*_type.TV{},
	}
}

func (m *Cache) GetRealTime(code string)( *_type.TV,error) {
	m.lock.RLock()
	defer m.lock.RUnlock()

	data ,ok := m.RealTimes[code]
	if !ok {
		return nil,errors.New("数据不存在")
	}

	return data,nil
}

func (m *Cache) SetRealTime(code string, data *_type.TV) {
	m.lock.Lock()
	defer m.lock.Unlock()

	m.RealTimes[code] = data
}
