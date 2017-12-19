package syncmap

import "sync"

// SyncMap 是并发安全字典的接口。
type SyncMap interface {
	Load(key interface{}) (value interface{}, ok bool)
	Store(key, value interface{})
	LoadOrStore(key, value interface{}) (actual interface{}, loaded bool)
	Delete(key interface{})
	Range(f func(key, value interface{}) bool)
}

// simpleSyncMap 为一个基于读写锁的并发安全字典的实现。
type simpleSyncMap struct {
	m map[interface{}]interface{}
	mu sync.RWMutex
}

func (m *simpleSyncMap) Load(key interface{}) (value interface{}, ok bool) {
	m.mu.RLock()
	value, ok = m.m[key]
	m.mu.RUnlock()
	return
}

func (m *simpleSyncMap) Store(key, value interface{}) {
	m.mu.Lock()
	m.m[key] = value
	m.mu.Unlock()
	return
}

func (m *simpleSyncMap) LoadOrStore(key, value interface{}) (actual interface{}, loaded bool) {
	m.mu.RLock()
	actual, loaded = m.m[key]
	m.mu.RUnlock()
	if loaded {
		return
	}
	m.mu.Lock()
	m.m[key] = value
	m.mu.Unlock()
	return value, false
}

func (m *simpleSyncMap) Delete(key interface{}) {
	m.mu.Lock()
	delete(m.m, key)
	m.mu.Unlock()
	return
}

func (m *simpleSyncMap) Range(f func(key, value interface{}) bool) {
	m.mu.RLock()
	for key, value := range m.m {
		if !f(key, value) {
			break
		}
	}
	m.mu.Unlock()
}


// NewSimpleSyncMap 用于创建并返回一个基于读写锁的并发安全字典。
func NewSimpleSyncMap() SyncMap {
	return &simpleSyncMap{
		m: make(map[interface{}]interface{}),
	}
}

// NewOfficialSyncMap 用于创建并返回一个官方的并发安全字典。
func NewOfficialSyncMap() SyncMap {
	return &sync.Map{}
}
