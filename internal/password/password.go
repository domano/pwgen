// Package password provides types and functions for password generation.
package password

// Generator can generate passwords with a given configuration
// passed via functional Options in its constructor.
type Generator struct {
	minLength, specialChars, nums int
}

// Option is the functional option type to allow variadic and
// generic configration of generators.
type Option func(*Generator)

// NewGenerator will create a Generator which can generate passwords.
// A number of Options can be passed to configure the resulting Generator.
// If no MinLength is given, then 16 is the default.
// If no SpecialChars or Numbers is given, then no limitation will be assumed.
func NewGenerator(options ...Option) Generator {
	g := Generator{minLength: 16}
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

// Nums configures the exact amount of numbers in generated passwords.
func Nums(amount int) Option {
	return func(g *Generator) {
		g.nums = amount
	}
}
