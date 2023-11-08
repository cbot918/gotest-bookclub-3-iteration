package strings

import (
	"fmt"
	"testing"
)

func TestContains(t *testing.T) {

	isContain := Contains("apple", "app")
	expected := true

	if isContain != expected {
		t.Errorf("expected %v but got %v", expected, isContain)
	}
}

func ExampleContains() {
	isContain := Contains("apple", "app")
	fmt.Println(isContain)
	// Output: true
}

func BenchmarkContains(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Contains("apple", "app")
	}
}
