package queue_test

import (
	"fmt"
	"testing"

	"github.com/kissoonjoshua/Learning-Go/data_structures/queue"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestQueue(t *testing.T) {
	t.Run("Enqueue", func(t *testing.T) {
		tests := []struct {
			input int
		}{{123}, {456}}

		for _, test := range tests {
			t.Run(fmt.Sprintf("val_%d", test.input), func(t *testing.T) {
				q := queue.Queue[int]{}
				q.Enqueue(test.input)
				val, ok := q.Peek()
				require.True(t, ok)
				assert.Equal(t, test.input, val)
			})
		}
	})

	t.Run("Dequeue", func(t *testing.T) {
		tests := []struct {
			name        string
			initial     []int
			expectedVal int
			expectedOk  bool
		}{
			{"dequeue from empty", []int{}, 0, false},
			{"dequeue from single", []int{1}, 1, true},
			{"dequeue from multiple", []int{1, 3, 5}, 1, true},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				q := queue.New(tt.initial...)
				act, ok := q.Dequeue()
				assert.Equal(t, tt.expectedOk, ok)
				assert.Equal(t, tt.expectedVal, act)
			})
		}
	})

	t.Run("Peek", func(t *testing.T) {
		tests := []struct {
			name        string
			initial     []int
			expectedVal int
			expectedOk  bool
		}{
			{"peek from empty", []int{}, 0, false},
			{"peek from single", []int{3}, 3, true},
			{"peek from multiple", []int{1, 3, 5}, 1, true},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				q := queue.New(tt.initial...)
				act, ok := q.Peek()
				assert.Equal(t, tt.expectedOk, ok)
				assert.Equal(t, tt.expectedVal, act)
			})
		}
	})

	t.Run("Size", func(t *testing.T) {
		tests := []struct {
			name     string
			initial  []int
			expected int
		}{
			{"size of empty", []int{}, 0},
			{"size of single", []int{3}, 1},
			{"size of multiple", []int{1, 3, 5}, 3},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				q := queue.New(tt.initial...)
				act := q.Size()
				assert.Equal(t, tt.expected, act)
			})
		}
	})

	t.Run("Empty", func(t *testing.T) {
		tests := []struct {
			name     string
			initial  []int
			expected bool
		}{
			{"empty queue", []int{}, true},
			{"non-empty queue", []int{1, 2, 3}, false},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				q := queue.New(tt.initial...)
				assert.Equal(t, tt.expected, q.Empty())
			})
		}
	})

	t.Run("Clear", func(t *testing.T) {
		q := queue.New(1, 2, 3)
		q.Clear()
		assert.True(t, q.Empty())
	})
}
