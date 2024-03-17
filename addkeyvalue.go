package main

import (
	"errors"
	"fmt"
	"sync"
)

// Capacity
// Linked List

// RedisClone represents a Redis-like object with a shared hashmap
type RedisClone struct {
	data     map[string]string
	capacity int32
	size     int32
	mu       sync.Mutex
}

// NewRedisClone creates a new instance of RedisClone
func NewRedisClone() *RedisClone {
	return &RedisClone{
		data: make(map[string]string),
	}
}

// AddKeyValue adds a key-value pair to the hashmap
func (r *RedisClone) AddKeyValue(key string, val string) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if r.size+1 > r.capacity {
		// Boot values out
	} else {
		r.data[key] = val
	}
}

func (r *RedisClone) GetValue(key string) (string, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if val, ok := r.data[key]; ok {
		return val, nil
	}

	errMsg := fmt.Sprintf("Key: [%s] does not exist.", key)

	return "", errors.New(errMsg)
}
