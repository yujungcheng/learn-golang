package concatenatingStrings

import "testing"

func BenchmarkStringFromAssignment(b *testing.B) {
	for n := 0; n < b.N; n++ {
		StringFromAssignment(100)
	}
}

func BenchmarkStringFromAppendJoin(b *testing.B) {
	for n := 0; n < b.N; n++ {
		StringFromAppendJoin(100)
	}
}

func BenchmarkStringFromBuffer(b *testing.B) {
	for n := 0; n < b.N; n++ {
		StringFromBuffer(100)
	}
}
