package main

import "sync"

type HashMap struct {
	lock sync.RWMutex
	m    map[string]interface{}
}

func NewHashMap() *HashMap {
	return &HashMap{m: make(map[string]interface{})}
}

func (this *HashMap) Get(key string) (interface{}, bool) {
	this.lock.RLock()
	defer this.lock.RUnlock()
	value, ok := this.m[key]
	return value, ok
}

func (this *HashMap) Put(key string, value interface{}) {
	this.lock.Lock()
	defer this.lock.Unlock()
	this.m[key] = value
}
