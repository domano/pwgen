package password

import (
	"testing"
)

func BenchmarkCryptoSource_Int63(b *testing.B) {
	for i := 0; i < b.N; i++ {
		c := cryptoSource{}
		c.Int63()
	}
}
