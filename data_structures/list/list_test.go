package list_test

import (
	"fmt"
	"testing"

	"github.com/kissoonjoshua/Learning-Go/data_structures/list"
	"github.com/stretchr/testify/assert"
)

func TestList(t *testing.T) {
	t.Run("PushFront", func(t *testing.T) {
		tests := []struct {
			key, value int
			initial    []struct{ key, value int }
		}{
			{1, 1, []struct{ key, value int }{{2, 3}, {4, 5}}},
			{4, 4, []struct{ key, value int }{{5, 6}, {7, 8}}},
		}

		for _, test := range tests {
			t.Run(fmt.Sprintf("val_%d", test.key), func(t *testing.T) {
				l := list.LinkedList[int, int]{}
				for _, item := range test.initial {
					l.PushFront(item.key, item.value)
				}
				l.PushFront(test.key, test.value)
				assert.Equal(t, test.key, l.Head.Value)
				assert.Equal(t, test.value, l.Head.Value)
			})
		}
	})

	t.Run("PushBack", func(t *testing.T) {
		tests := []struct {
			key, value int
			initial    []struct{ key, value int }
		}{
			{1, 1, []struct{ key, value int }{{2, 3}, {4, 5}}},
			{4, 4, []struct{ key, value int }{{5, 6}, {7, 8}}},
		}

		for _, test := range tests {
			t.Run(fmt.Sprintf("val_%d", test.key), func(t *testing.T) {
				l := list.LinkedList[int, int]{}
				for _, item := range test.initial {
					l.PushBack(item.key, item.value)
				}
				l.PushBack(test.key, test.value)
				assert.Equal(t, test.key, l.Tail.Key)
				assert.Equal(t, test.value, l.Tail.Value)
			})
		}
	})

	t.Run("PopFront", func(t *testing.T) {
		tests := []struct {
			name          string
			initial       []struct{ key, value int }
			expectedKey   int
			expectedValue int
			expectedOk    bool
		}{
			{"non-empty list", []struct{ key, value int }{{2, 3}, {4, 5}}, 2, 3, true},
			{"empty list", []struct{ key, value int }{}, 0, 0, false},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				l := list.LinkedList[int, int]{}
				for _, item := range test.initial {
					l.PushBack(item.key, item.value)
				}
				key, val, ok := l.PopFront()
				assert.Equal(t, test.expectedOk, ok)
				assert.Equal(t, test.expectedKey, key)
				assert.Equal(t, test.expectedValue, val)
			})
		}
	})

	t.Run("PopBack", func(t *testing.T) {
		tests := []struct {
			name          string
			initial       []struct{ key, value int }
			expectedKey   int
			expectedValue int
			expectedOk    bool
		}{
			{"non-empty list", []struct{ key, value int }{{2, 3}, {4, 5}}, 4, 5, true},
			{"empty list", []struct{ key, value int }{}, 0, 0, false},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				l := list.LinkedList[int, int]{}
				for _, item := range tt.initial {
					l.PushBack(item.key, item.value)
				}
				key, val, ok := l.PopBack()
				assert.Equal(t, tt.expectedOk, ok)
				assert.Equal(t, tt.expectedKey, key)
				assert.Equal(t, tt.expectedValue, val)
			})
		}
	})
}
