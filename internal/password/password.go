// Package password provides types and functions for password generation.
package password

import (
	"bytes"
	"math"
)

// Generator can generate passwords with a given configuration
// passed via functional Options in its constructor.
type Generator struct {
	minLength, specialChars, nums int
	swap bool
}

// Option is the functional option type to allow variadic and
// generic configration of generators.
type Option func(*Generator)

// The sets of letters used to generate our passwords
const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const numbers = "0123456789"
const specialChars = " !\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~"
const vowels, vowelNums = "aAeEiIoO", "4310"


// NewGenerator will create a Generator which can generate passwords.
// A number of Options can be passed to configure the resulting Generator.
func NewGenerator(options ...Option) Generator {
	g := Generator{}
	for i := range options {
		options[i](&g)
	}
	return g
}

// MinLength configures a minimum length for generated passwords.
func MinLength(length int) Option {
	return func(g *Generator) {
		g.minLength = length
	}
}

// SpecialChars configures the exact amount of special characters in generated passwords.
func SpecialChars(amount int) Option {
	return func(g *Generator) {
		g.specialChars = amount
	}
}

// Numbers configures the exact amount of numbers in generated passwords.
func Numbers(amount int) Option {
	return func(g *Generator) {
		g.nums = amount
	}
}

// Swap configures if vowels should be swapped with numbers
func Swap(shouldSwap bool) Option {
	return func(g *Generator) {
		g.swap = shouldSwap
	}
}

// Password generates a password with the generators' configuration
func (g Generator) Password() string {
	var passwordBytes []byte

	// Create numbers, special chars and letters for the password randomly
	passwordBytes = g.generate(passwordBytes)

	// Shuffle the password
	password := g.shuffle(passwordBytes)

	return string(password)
}

func (g Generator) generate(pw []byte) []byte {
	pw = append(pw, randomBytes(numbers, g.nums)...)
	pw = append(pw, randomBytes(specialChars, g.specialChars)...)
	if g.minLength > len(pw) {
		pw = append(pw, randomBytes(letters, g.minLength-len(pw))...)
	}
	return pw
}

func (g Generator) shuffle(passwordBytes []byte) []byte {
	var password = make([]byte, len(passwordBytes))
	for i, v := range random.Perm(len(passwordBytes)) {
		if g.swap {
			g.swapVowel(passwordBytes[v])
		}
		password[i] = passwordBytes[v]
	}
	return password
}

func (g Generator) swapVowel(char byte) byte{
		index := bytes.IndexByte([]byte(vowels), char)
		if index > 0 && random.Intn(2) == 1 {
			vowelNumIndex := int(math.Ceil(float64(index) / 2.0))
			return vowelNums[vowelNumIndex]
		}
		return char
}


func randomBytes(from string, length int) []byte {
	str := make([]byte, length)
	for i := 0; i < length; i++ {
		str[i] = randomChar(from)
	}
	return str
}

func randomChar(from string) byte {
	i := random.Intn(len(from))
	return from[i]
}
