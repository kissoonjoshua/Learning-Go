package stack

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStack(t *testing.T) {
	t.Run("Push", func(t *testing.T) {
		tests := []struct {
			input int
		}{ {123}, {456}, }

		for _, test := range tests {
			t.Run(fmt.Sprintf("val_%d", test.input), func(t *testing.T) {
				s := Stack[int]{} 
				s.Push(test.input)
				val, err := s.Peek()
				require.NoError(t, err)
				assert.Equal(t, test.input, val) 
			})
		}
	})

	t.Run("Pop", func(t *testing.T) {
		tests := []struct {
			name     string
			initial  []int
			expected int
		}{
			{"pop from single", []int{1}, 1},
			{"pop from multiple", []int{5, 3, 1}, 1},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				s := Stack[int]{items: tt.initial}
				
				act, err := s.Pop()
				
				require.NoError(t, err)
				assert.Equal(t, tt.expected, act)
			})
		}
	})
}