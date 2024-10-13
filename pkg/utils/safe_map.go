package utils

import "sync"

type Safe[K, K1 comparable, V any] struct {
	sync.RWMutex
	m map[K]map[K1]V
}

func NewSafe[K, K1 comparable, V any]() *Safe[K, K1, V] {
	return &Safe[K, K1, V]{
		m: make(map[K]map[K1]V),
	}
}

func (s *Safe[K, K1, V]) Set(key K, value map[K1]V) {
	s.Lock()
	defer s.Unlock()

	s.m[key] = value
}

func (s *Safe[K, K1, V]) SetValue(key K, key1 K1, value V) {
	s.Lock()
	defer s.Unlock()

	_, ok := s.m[key]

	if !ok {
		s.m[key] = make(map[K1]V)
	}

	s.m[key][key1] = value
}

func (s *Safe[K, K1, V]) Get(k K) (value map[K1]V, ok bool) {
	s.RLock()
	defer s.RUnlock()

	value, ok = s.m[k]

	return value, ok
}

func (s *Safe[K, K1, V]) GetValue(key K, key1 K1) (value V, ok bool) {
	s.RLock()
	defer s.RUnlock()

	result, ok := s.m[key]
	if !ok {
		return value, ok
	}

	return result[key1], ok
}

func (s *Safe[K, K1, V]) GetValueByIndex(key K, index int) (value V, ok bool) {
	s.RLock()
	defer s.RUnlock()

	result, ok := s.m[key]
	if !ok {
		return value, ok
	}

	i := 0

	for keyM, _ := range result {
		if index == i {
			return result[keyM], true
		}

		i++
	}

	return value, ok
}

func (s *Safe[K, K1, V]) Del(k K) {
	s.Lock()
	defer s.Unlock()

	delete(s.m, k)
}

func (s *Safe[K, K1, V]) Keys() []K {
	s.RLock()
	defer s.RUnlock()

	keys := make([]K, 0, len(s.m))
	for k := range s.m {
		keys = append(keys, k)
	}

	return keys
}

func (s *Safe[K, K1, V]) SetAll(newMap map[K]map[K1]V) {
	s.Lock()
	defer s.Unlock()

	s.m = newMap
}

func (s *Safe[K, K1, V]) LenValueByKey(k K) int {
	s.RLock()
	defer s.RUnlock()

	return len(s.m[k])
}
