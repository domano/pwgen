package password

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewGenerator(t *testing.T) {
	testCases := []struct {
		desc     string
		options  []Option
		expected Generator
	}{
		{
			desc:     "Generator without configuration",
			options:  []Option{},
			expected: Generator{minLength: 16},
		},
		{
			desc:     "Generator with min length 5",
			options:  []Option{MinLength(5)},
			expected: Generator{minLength: 5},
		},
		{
			desc:     "Generator with special characters 4",
			options:  []Option{SpecialChars(4)},
			expected: Generator{minLength: 16, specialChars: 4},
		},
		{
			desc:     "Generator with nums 3",
			options:  []Option{Nums(3)},
			expected: Generator{minLength: 16, nums: 3},
		},
		{
			desc:     "Generator with nums 3, special chars 4 & min length 5",
			options:  []Option{Nums(3), SpecialChars(4), MinLength(5)},
			expected: Generator{minLength: 5, nums: 3, specialChars: 4},
		},
		{
			desc:     "Generator with special characters 4 & special characters 5. 5 should overwrite 4",
			options:  []Option{SpecialChars(4), SpecialChars(5)},
			expected: Generator{minLength: 16, specialChars: 5},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			generator := NewGenerator(tC.options...)
			assert.Equal(t, tC.expected, generator)
		})
	}
}
