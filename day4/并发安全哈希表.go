package main

import (
	"fmt"
	"sync"
)

type ConcurrentMap struct {
	mu    sync.RWMutex
	table map[string]interface{}
}

func InitConcurrentMap() *ConcurrentMap {
	return &ConcurrentMap{
		table: make(map[string]interface{}),
	}
}

func (m *ConcurrentMap) Get(key string) (interface{}, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	v, ok := m.table[key]
	return v, ok
}

func (m *ConcurrentMap) Set(key string, value interface{}) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.table[key] = value
}

func (m *ConcurrentMap) Delete(key string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	delete(m.table, key)
}

func main() {
	cm := InitConcurrentMap()

	// 并发安全地设置键值对
	cm.Set("key1", "value1")
	cm.Set("key2", 123)

	// 并发安全地获取值
	if value, ok := cm.Get("key1"); ok {
		fmt.Println("key1:", value)
	} else {
		fmt.Println("key1 not found")
	}

	// 并发安全地删除键
	cm.Delete("key2")

	if _, ok := cm.Get("key2"); !ok {
		fmt.Println("key2 has been deleted")
	}
}
