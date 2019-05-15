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

// The sets of letters used to generate our passwords
const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const numbers = "0123456789"
const specialChars = " !\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~"

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

// Password generates a password with the generators' configuration
func (g Generator) Password() string {
	var passwordBytes []byte

	// Create numbers, special chars and letters for the password randomly
	passwordBytes = append(passwordBytes, randomBytes(numbers, g.nums)...)
	passwordBytes = append(passwordBytes, randomBytes(specialChars, g.specialChars)...)
	if g.minLength > len(passwordBytes) {
		passwordBytes = append(passwordBytes, randomBytes(letters, g.minLength-len(passwordBytes))...)
	}

	// Shuffle the password
	var password = make([]byte, len(passwordBytes))
	for i, v := range random.Perm(len(passwordBytes)) {
		password[i] = passwordBytes[v]
	}

	return string(password)
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
