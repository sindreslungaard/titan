package hh

import "sync"

type Dictionary[K comparable, T any] struct {
	sync.RWMutex
	src   []*T
	store map[K]*T
}

func store[K comparable, T any]() Dictionary[K, T] {
	return Dictionary[K, T]{
		src:   []*T{},
		store: make(map[K]*T),
	}
}

func (d *Dictionary[K, T]) iter() []*T {
	d.RLock()
	defer d.RUnlock()

	tmp := make([]*T, len(d.src))

	copy(tmp, d.src)

	return tmp
}

func (d *Dictionary[K, T]) add(key K, value *T) {
	d.Lock()
	defer d.Unlock()

	_, ok := d.store[key]

	if ok {
		return
	}

	d.src = append(d.src, value)
	d.store[key] = value
}

func (d *Dictionary[K, T]) remove(key K) {
	d.Lock()
	defer d.Unlock()

	value, ok := d.store[key]

	if !ok {
		return
	}

	toRemove := -1
	for i, v := range d.src {
		if v == value {
			toRemove = i
			break
		}
	}

	if toRemove < 0 {
		return
	}

	d.src = append(d.src[:toRemove], d.src[toRemove+1:]...)
	delete(d.store, key)
}

func (d *Dictionary[K, T]) find(key K) (*T, bool) {
	d.RLock()
	defer d.RUnlock()

	value, ok := d.store[key]

	return value, ok
}
