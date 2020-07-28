package core

import (
	"errors"
	"reflect"
	"sync"
)

var (
	poolManager PoolManager
)

type PoolManager struct {
	poolMap map[reflect.Type]sync.Pool
	mu      sync.RWMutex
}

func newPoolManager() PoolManager {
	poolManager := PoolManager{
		poolMap: make(map[reflect.Type]sync.Pool),
	}
	return poolManager
}

func (manager PoolManager) register(typ *Type, newFunc func() interface{}) error {
	if typ == nil {
		return errors.New("pool type cannot be null")
	}
	if newFunc == nil {
		return errors.New("new func type cannot be null")
	}
	poolType := typ.Typ
	if poolType.Kind() == reflect.Ptr {
		poolType = poolType.Elem()
	}
	manager.mu.Lock()
	if _, ok := manager.poolMap[poolType]; ok {
		manager.mu.Unlock()
		return errors.New("pool type already exists in pool map")
	}
	manager.mu.Unlock()
	newPool := sync.Pool{
		New: func() interface{} {
			return newFunc()
		},
	}
	manager.mu.Lock()
	manager.poolMap[poolType] = newPool
	manager.mu.Unlock()
	return nil
}

func (manager PoolManager) get(typ *Type) (interface{}, error) {
	if typ == nil {
		return nil, errors.New("pool type cannot be null")
	}
	poolType := typ.Typ
	if poolType.Kind() == reflect.Ptr {
		poolType = poolType.Elem()
	}
	manager.mu.Lock()
	if pool, ok := manager.poolMap[poolType]; ok {
		manager.mu.Unlock()
		return pool.Get(), nil
	}
	manager.mu.Unlock()
	return nil, nil
}

func (manager PoolManager) put(instance interface{}) {
	if instance == nil {
		return
	}
	typ := GetType(instance)
	poolType := typ.Typ
	if poolType.Kind() == reflect.Ptr {
		poolType = poolType.Elem()
	}
	manager.mu.Lock()
	if pool, ok := manager.poolMap[poolType]; ok {
		manager.mu.Unlock()
		pool.Put(instance)
		return
	}
	manager.mu.Unlock()
}

func RegisterPool(typ *Type, newFunc func() interface{}) {
	err := poolManager.register(typ, newFunc)
	if err != nil {
		panic(err)
	}
}

func GetFromPool(typ *Type) interface{} {
	instance, err := poolManager.get(typ)
	if err != nil {
		panic(err)
	}
	return instance
}

func PutToPool(instances ...interface{}) {
	for _, instance := range instances {
		poolManager.put(instance)
	}
}
