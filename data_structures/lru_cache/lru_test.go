package lru_test

import (
	"testing"

	lru "github.com/kissoonjoshua/Learning-Go/data_structures/lru_cache"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type pair struct{ key, val int }

func TestLRU(t *testing.T) {
	t.Run("TestPut", func(t *testing.T) {
		tests := []struct {
			name          string
			capacity      int
			initial       []pair
			expectedEntry pair
			expectedOk    bool
		}{
			{"Test Put Unique Key", 10, []pair{{1, 2}}, pair{1, 2}, true},
			{"Test Put Duplicate Key", 10, []pair{{1, 2}, {1, 3}}, pair{1, 3}, true},
			{"Test Put Evict", 2, []pair{{1, 2}, {2, 3}, {3, 4}}, pair{1, 0}, false},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				cache := lru.New[int, int](tt.capacity)
				for _, item := range tt.initial {
					cache.Put(item.key, item.val)
				}
				actVal, ok := cache.Get(tt.expectedEntry.key)
				require.Equal(t, tt.expectedOk, ok)
				assert.Equal(t, tt.expectedEntry.val, actVal)
			})
		}
	})

	t.Run("TestGet", func(t *testing.T) {
		tests := []struct {
			name          string
			capacity      int
			initial       []pair
			expectedEntry pair
			expectedOk    bool
		}{
			{"Test Get Existing Key", 10, []pair{{1, 2}}, pair{1, 2}, true},
			{"Test Get Non-Existing Key", 10, []pair{{1, 2}}, pair{2, 0}, false},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				cache := lru.New[int, int](tt.capacity)
				for _, item := range tt.initial {
					cache.Put(item.key, item.val)
				}
				actVal, ok := cache.Get(tt.expectedEntry.key)
				require.Equal(t, tt.expectedOk, ok)
				assert.Equal(t, tt.expectedEntry.val, actVal)
			})
		}
	})

	t.Run("TestRemove", func(t *testing.T) {
		tests := []struct {
			name         string
			capacity     int
			initial      []pair
			removeKey    int
			expectedSize int
		}{
			{"Test Remove Existing Key", 10, []pair{{1, 2}}, 1, 0},
			{"Test Remove Non-Existing Key", 10, []pair{{1, 2}}, 2, 1},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				cache := lru.New[int, int](tt.capacity)
				for _, item := range tt.initial {
					cache.Put(item.key, item.val)
				}
				cache.Remove(tt.removeKey)
				assert.Equal(t, tt.expectedSize, cache.Size())
			})
		}
	})

	t.Run("TestSize", func(t *testing.T) {
		tests := []struct {
			name         string
			capacity     int
			initial      []pair
			expectedSize int
		}{
			{"Test Size Empty", 10, []pair{}, 0},
			{"Test Size Non-Empty", 10, []pair{{1, 2}}, 1},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				cache := lru.New[int, int](tt.capacity)
				for _, item := range tt.initial {
					cache.Put(item.key, item.val)
				}
				assert.Equal(t, tt.expectedSize, cache.Size())
			})
		}
	})

	t.Run("TestClear", func(t *testing.T) {
		cache := lru.New[int, int](10)
		cache.Put(1, 2)
		cache.Clear()
		assert.Equal(t, 0, cache.Size())
	})
}
