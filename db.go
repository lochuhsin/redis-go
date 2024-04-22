package rredis

import (
	"fmt"
	"sync"
)

var (
	b     *bucket
	bOnce *sync.Once = &sync.Once{}
)

type database struct {
	collections map[collectionType]collection
	mu          sync.RWMutex
}

func (d *database) get(key string) (string, error) {
	d.mu.RLock()
	defer d.mu.RUnlock()
	ctype, exist := d.exist(key)
	if !exist {
		return *new(string), nil
	}
	db := d.collections[ctype]
	entry, _ := db.get(key)

	return entry.v, nil
}

func (d *database) add(e entry) error {
	d.mu.Lock()
	defer d.mu.Unlock()
	ctype, exist := d.exist(e.k)
	if exist && ctype != HASHMAP {
		return DatabaseKeyError{key: e.k, detail: fmt.Sprintf("key exists in different collection type: %v", ctype)}
	}
	if exist && ctype == HASHMAP {
		d.collections[HASHMAP].set(e)
	}

	if !exist {
		if collection, ok := d.collections[HASHMAP]; ok {
			collection.set(e)
		} else {
			collection = newHashmap()
			collection.set(e)
			d.collections[HASHMAP] = collection
		}
	}
	return nil
}

func (d *database) delete(key string) error {
	return nil
}

func (d *database) exist(key string) (collectionType, bool) {
	for ctype, collection := range d.collections {
		if collection.exist(key) {
			return ctype, true
		}
	}
	return "", false
}

func newDatabase() *database {
	return &database{map[collectionType]collection{}, sync.RWMutex{}}
}

type bucket struct {
	dbs []*database // convert -> array of database
	sync.RWMutex
}

func newBucket(dbCount int) *bucket {
	arr := make([]*database, dbCount)
	return &bucket{
		dbs: arr,
	}
}

func (c *bucket) create(db int) error {
	c.Lock()
	defer c.Unlock()
	if db >= len(c.dbs) {
		return DatabaseError{"Exceed maximum number of database"}
	}
	if c.dbs[db] != nil {
		return DatabaseExistsError{db}
	}
	c.dbs[db] = newDatabase()
	return nil
}

func (c *bucket) flush(db int) error {
	c.Lock()
	defer c.Unlock()
	if c.dbs[db] != nil {
		c.dbs[db] = nil
		return nil
	}
	return DatabaseNotFoundError{db}
}

func (c *bucket) get(db int) (*database, error) {
	c.RLock()
	defer c.RUnlock()
	if db >= len(c.dbs) {
		return nil, DatabaseError{"Exceed maximum number of database"}
	} else if db < 0 {
		return nil, DatabaseError{"Database namespace should be greater than 0"}
	}
	if db := c.dbs[db]; db != nil {
		return db, nil
	}
	return nil, DatabaseNotFoundError{db}
}

func (c *bucket) exist(db int) bool {
	if db >= len(c.dbs) {
		return false
	}
	return c.dbs[db] != nil
}

func initBucket() {
	bOnce.Do(func() {
		b = newBucket(DEFAULT_DATABASE_COUNT)
		b.create(0) // creating database 0 as default, just like redis
	})
}
