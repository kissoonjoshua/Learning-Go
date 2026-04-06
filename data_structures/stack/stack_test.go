package stack_test

import (
	"fmt"
	"testing"

	"github.com/kissoonjoshua/Learning-Go/data_structures/stack"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStack(t *testing.T) {
	t.Run("Push", func(t *testing.T) {
		tests := []struct {
			input int
		}{{123}, {456}}

		for _, tt := range tests {
			t.Run(fmt.Sprintf("val_%d", tt.input), func(t *testing.T) {
				s := stack.New(tt.input)
				val, ok := s.Peek()
				require.True(t, ok)
				assert.Equal(t, tt.input, val)
			})
		}
	})

	t.Run("Pop", func(t *testing.T) {
		tests := []struct {
			name        string
			initial     []int
			expectedVal int
			expectedOk  bool
		}{
			{"pop from empty", []int{}, 0, false},
			{"pop from single", []int{1}, 1, true},
			{"pop from multiple", []int{5, 3, 1}, 1, true},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				s := stack.New(tt.initial...)
				act, ok := s.Pop()
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
			{"peek from multiple", []int{5, 3, 9}, 9, true},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				s := stack.New(tt.initial...)
				act, ok := s.Peek()
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
			{"size of multiple", []int{5, 3, 9}, 3},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				s := stack.New(tt.initial...)
				act := s.Size()
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
			{"empty stack", []int{}, true},
			{"non-empty stack", []int{1, 2, 3}, false},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				s := stack.New(tt.initial...)
				assert.Equal(t, tt.expected, s.Empty())
			})
		}
	})

	t.Run("Clear", func(t *testing.T) {
		s := stack.New(1, 2, 3)
		s.Clear()
		assert.True(t, s.Empty())
	})
}
