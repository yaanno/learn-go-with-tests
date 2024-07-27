package iterations

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("Repeat should return the given string", func(t *testing.T) {
		repeated := Repeat("a", 2)
		expected := "aa"
		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	})
	t.Run("it should return empty string when negative repeat count is given", func(t *testing.T) {
		repeated := Repeat("a", -1)
		expected := ""
		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	})
}

func TestRepeatCompare(t *testing.T) {
	repeated := Repeat("a", 2)
	expected := "aa"
	compared := strings.Compare(repeated, expected)
	if compared != 0 {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func TestCompareBuiltin(t *testing.T) {
	ownrepeated := Repeat("a", 2)
	builtinrepeated := strings.Repeat("a", 2)
	if builtinrepeated != ownrepeated {
		t.Errorf("expected %q but got %q", builtinrepeated, ownrepeated)
	}
}

func TestRepeatSpaced(t *testing.T) {
	repeated := Repeat("a ", 2)
	spaced := Spaced(repeated)
	expected := []string{"a", "a"}
	if !reflect.DeepEqual(spaced, expected) {
		t.Errorf("expected %q but got %q", expected, spaced)
	}
}

func Spaced(s string) []string {
	return strings.Fields(s)
}

func Repeat(char string, iter int) string {
	var repeated string
	for i := 0; i < iter; i++ {
		repeated += char
	}
	return repeated
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	result := Repeat("a", 2)
	fmt.Println(result)
	// output: aa
}
