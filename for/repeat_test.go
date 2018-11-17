package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, excepted, repeated string) {
		t.Helper()
		if excepted != repeated {
			t.Errorf("excepted '%s' but repeated '%s'", repeated, excepted)
		}
	}

	t.Run("repeat", func(t *testing.T) {
		repeated := Repeat("a", 5)
		excepted := "aaaaa"
		assertCorrectMessage(t, repeated, excepted)
	})

	t.Run("specefy repeated time", func(t *testing.T) {
		repeated := Repeat("a", 6)
		excepted := "aaaaaa"
		assertCorrectMessage(t, repeated, excepted)
	})
}

func ExampleRepeat() {
	repeated := Repeat("b", 7)
	fmt.Println(repeated)
	// Output: bbbbbbb
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}