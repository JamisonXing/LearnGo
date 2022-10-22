package factory

import (
	"bookstore/store"
	"fmt"
	"sync"
)

// 设计模式多种创建模式，创建一个满足Store接口的实例
var (
	providersMu sync.RWMutex
	providers   = make(map[string]store.Store)
)

func Register(name string, p store.Store) {
	providersMu.Lock()
	defer providersMu.Unlock()
	if p == nil {
		panic("store: Register provider is null" + name)
	}

	if _, dup := providers[name]; dup {
		panic("store: Register called twice for provider" + name)
	}
	providers[name] = p
}

func New(providerName string) (store.Store, error) {
	providersMu.RLock()
	p, ok := providers[providerName]
	providersMu.RUnlock()
	if !ok {
		return nil, fmt.Errorf("store: unknown provider %s", providerName)
	}

	return p, nil
}