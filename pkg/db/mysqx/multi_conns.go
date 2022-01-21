package mysqx

import (
	"database/sql"
	"log"
	"sync"
)

type MysqlMultiConnMap struct {
	dbConnectionMap map[string]*Connection
	sync.RWMutex
}

func NewMultiConnectionMap() *MysqlMultiConnMap {
	return &MysqlMultiConnMap{dbConnectionMap: make(map[string]*Connection)}
}

func (m *MysqlMultiConnMap) Get(database string) *sql.DB {
	m.RLock()
	defer m.RUnlock()
	if conn, ok := m.dbConnectionMap[database]; ok {
		return conn.GetDB()
	}
	log.Fatalf("db %s not found", database)
	return nil
}

func (m *MysqlMultiConnMap) Add(database string, connection *Connection) {
	m.Lock()
	defer m.Unlock()
	m.dbConnectionMap[database] = connection
}

var _ DBGetter = (*MysqlMultiConnMap)(nil)
