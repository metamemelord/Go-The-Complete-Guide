package benchmarks

import (
	"fmt"
	"testing"
)

func TestGreet(t *testing.T) {
	s := Greet("Gaurav")
	if s != "Welcome my dear Gaurav" {
		t.Error("Expected 'Welcome my dear Gaurav', got '" + s + "'")
	}
}

func ExampleGreet() {
	fmt.Printf(Greet("Gaurav"))
	// Output:
	// Welcome my dear Gaurav
}

func BenchmarkGreet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Greet("Gaurav")
	}
}
