package iteration

import (
	"fmt"
	"testing"
)

func ExampleRepeat() {
	fmt.Println(Repeat("a", 5))
	// Output: aaaaa
}

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 5)
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func TestRepeatWithCount(t *testing.T) {
	repeated := Repeat("a", 3)
	expected := "aaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 0)
	}
}