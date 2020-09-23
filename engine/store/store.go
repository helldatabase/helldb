// package store provides APIs for low level
// interaction with the key-value data store.
package store

import (
	"encoding/json"
	"sync"

	"helldb/engine/types"
)

// Store is a singleton struct (per circle) used for
// keeping track of pairs (circle specific) with thread
// safe operations on members.
type Store struct {
	clust  uint64
	length uint64
	mutex  sync.Mutex
	pairs  map[string]types.BaseType
}

// Init returns a reference to a single Store struct.
// It should only be called once per initialization of circle
// to maintain only one single store per circle.
func Init() *Store {
	s := Store{}
	s.pairs = make(map[string]types.BaseType)
	return &s
}

// Del performs a thread safe delete over keys provided and
// returns a slice of Booleans (BaseType implement) signifying
// successful deletion if key was found.
func (s *Store) Del(keys []string) []types.Boolean {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	var statuses []types.Boolean
	for _, key := range keys {
		if _, ok := s.pairs[key]; ok {
			delete(s.pairs, key)
			statuses = append(statuses, *types.NewBoolean(true))
		} else {
			statuses = append(statuses, *types.NewBoolean(false))
		}
	}
	return statuses
}

// Get performs a thread safe read over the keys provided
// and returns a list of values that implement BaseType. A
// native nil is returned for keys not found.
func (s *Store) Get(keys []string) []types.BaseType {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	var values = make([]types.BaseType, len(keys), len(keys))
	for i, key := range keys {
		if val, ok := s.pairs[key]; ok {
			values[i] = val
		} else {
			values[i] = nil
		}
	}
	return values
}

// JSON returns a utf-8 string as JSON with native conversions
// between types that implement BaseTypes including Collections.
func (s *Store) JSON() string {
	var object = make(map[string]interface{})
	for key, val := range s.pairs {
		object[key] = val.Native()
	}
	j, _ := json.Marshal(object)
	return string(j)
}

// Len returns a uint64 for the number of keys in the circle
// specific Store.
func (s *Store) Len() uint64 {
	return s.length
}

// Put inserts (or updates) a string into a circle specific
// store instance with the provided val implementing BaseType.
func (s *Store) Put(key string, val types.BaseType) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.length++
	s.pairs[key] = val
}
