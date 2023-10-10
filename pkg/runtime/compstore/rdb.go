package compstore

import (
	"github.com/dapr/components-contrib/rdb"
)

func (c *ComponentStore) AddRDB(name string, rdbClient rdb.RDB) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.rdbs[name] = rdbClient
}

func (c *ComponentStore) GetRDB(name string) (rdb.RDB, bool) {
	c.lock.RLock()
	defer c.lock.RUnlock()
	rdbClient, ok := c.rdbs[name]
	return rdbClient, ok
}

func (c *ComponentStore) ListRDBs() map[string]rdb.RDB {
	c.lock.RLock()
	defer c.lock.RUnlock()
	return c.rdbs
}

func (c *ComponentStore) DeleteRDB(name string) {
	c.lock.Lock()
	defer c.lock.Unlock()
	delete(c.rdbs, name)
}

func (c *ComponentStore) RDBsLen() int {
	c.lock.RLock()
	defer c.lock.RUnlock()
	return len(c.rdbs)
}
