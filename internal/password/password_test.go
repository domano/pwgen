package password

import (
	"strings"
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
			expected: Generator{},
		},
		{
			desc:     "Generator with min length 5",
			options:  []Option{MinLength(5)},
			expected: Generator{minLength: 5},
		},
		{
			desc:     "Generator with special characters 4",
			options:  []Option{SpecialChars(4)},
			expected: Generator{specialChars: 4},
		},
		{
			desc:     "Generator with nums 3",
			options:  []Option{Nums(3)},
			expected: Generator{nums: 3},
		},
		{
			desc:     "Generator with nums 3, special chars 4 & min length 5",
			options:  []Option{Nums(3), SpecialChars(4), MinLength(5)},
			expected: Generator{minLength: 5, nums: 3, specialChars: 4},
		},
		{
			desc:     "Generator with special characters 4 & special characters 5. 5 should overwrite 4",
			options:  []Option{SpecialChars(4), SpecialChars(5)},
			expected: Generator{specialChars: 5},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			// when
			generator := NewGenerator(tC.options...)

			// then
			assert.Equal(t, tC.expected, generator)
		})
	}
}

func TestPassword(t *testing.T) {
	testCases := []struct {
		desc                          string
		minLength, specialChars, nums int
	}{
		{
			desc:      "3 minimum length, 1 special char, 1 number",
			minLength: 3, specialChars: 1, nums: 1,
		},
		{
			desc:      "3 minimum length, 5 special char, 5 number",
			minLength: 3, specialChars: 5, nums: 5,
		},
		{
			desc:      "10 minimum length, 0 special char, 0 number",
			minLength: 10, specialChars: 0, nums: 0,
		},
		{
			desc:      "10 minimum length, 0 special char, 10 number",
			minLength: 10, specialChars: 0, nums: 10,
		},
		{
			desc:      "10 minimum length, 10 special char, 0 number",
			minLength: 10, specialChars: 10, nums: 0,
		},
		{
			desc:      "0 minimum length, 10 special char, 0 number",
			minLength: 0, specialChars: 10, nums: 0,
		},
		{
			desc:      "0 minimum length, 0 special char, 10 number",
			minLength: 0, specialChars: 0, nums: 10,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			//given
			generator := Generator{minLength: tC.minLength, specialChars: tC.specialChars, nums: tC.nums}

			//when
			pw := generator.Password()

			//then
			assert.True(t, len(pw) >= tC.minLength, "password was below min length")
			assert.True(t, countAny(pw, specialChars) >= tC.specialChars, "password did not have enough special chars")
			assert.True(t, countAny(pw, numbers) >= tC.nums, "password did not have enough numbers")
		})
	}
}

// Counts occurences of any char in chars in s
func countAny(s, chars string) int {
	var count int
	for _, c := range chars {
		count += strings.Count(s, string(c))
	}
	return count
}
