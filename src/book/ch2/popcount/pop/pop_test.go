package pop

import (
	
	"testing"
)

func BenchmarkPopCount(b *testing.B){
	var x uint64 = 0xFFFFFFFFFFFFFFFF
	for i := 0; i < b.N; i++ {
		PopCount(x)
	}
}