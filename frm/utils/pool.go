package utils

import "sync"

type pool[T any] struct {
	pool sync.Pool
}

func NewPool[T any]() *pool[T] {
	return &pool[T]{
		pool: sync.Pool{
			New: func() interface{} {
				return new(T)
			},
		},
	}
}

func (this_ *pool[T]) Get() *T {
	return this_.pool.Get().(*T)
}

func (this_ *pool[T]) Put(v *T) {
	this_.pool.Put(v)
}
