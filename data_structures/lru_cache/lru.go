package lru

import "github.com/kissoonjoshua/Learning-Go/data_structures/list"

type LRU[K comparable, V any] struct {
	cache    map[K]*list.Node[K, V]
	recent   *list.LinkedList[K, V]
	capacity int
}

func New[K comparable, V any](capacity int) *LRU[K, V] {
	return &LRU[K, V]{
		cache:    make(map[K]*list.Node[K, V], capacity),
		recent:   list.New[K, V](),
		capacity: capacity,
	}
}

func (l *LRU[K, V]) Get(key K) (V, bool) {
	if node, ok := l.cache[key]; ok {
		l.recent.MoveToFront(node)
		return node.Value, true
	}
	var zero V
	return zero, false
}

func (l *LRU[K, V]) Put(key K, value V) {
	if node, ok := l.cache[key]; ok {
		node.Value = value
		l.recent.MoveToFront(node)
		return
	}
	if len(l.cache) >= l.capacity {
		evictKey, _, ok := l.recent.PopBack()
		if ok {
			delete(l.cache, evictKey)
		}
	}
	l.recent.PushFront(key, value)
	l.cache[key] = l.recent.Head
}

func (l *LRU[K, V]) Remove(key K) {
	if node, ok := l.cache[key]; ok {
		l.recent.Remove(node)
		delete(l.cache, key)
	}
}

func (l *LRU[K, V]) Size() int {
	return len(l.cache)
}

func (l *LRU[K, V]) Clear() {
	clear(l.cache)
	l.recent = list.New[K, V]()
}
