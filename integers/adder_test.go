package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Add(1, 2)
	expected := 3
	if sum != expected {
		t.Errorf("expected %d, but got %d", expected, sum)
	}
}

// Add takes two integers and returns the sum of them
func Add(x, y int) int {
	return x + y
}

func ExampleAdd() {
	sum := Add(1, 1)
	fmt.Println(sum)
	// Output: 2
}
