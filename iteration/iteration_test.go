package iteration

import "testing"


func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 6)
	expected := "aaaaaa"

	if repeated != expected {
		t.Errorf("Expected %v, got %v instead\n", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 6)
	}
}