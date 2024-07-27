package iterations

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func TestRepeat(t *testing.T) {
	char := "a"
	t.Run("Repeat should return the given string", func(t *testing.T) {
		repeated := Repeat(char, 2)
		expected := "aa"
		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	})
	t.Run("it should return empty string when negative repeat count is given", func(t *testing.T) {
		repeated := Repeat(char, -1)
		expected := ""
		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	})
}

func TestRepeatCompare(t *testing.T) {
	char := "a"
	repeated := Repeat(char, 2)
	expected := "aa"
	compared := strings.Compare(repeated, expected)
	if compared != 0 {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func TestCompareBuiltin(t *testing.T) {
	char := "a"
	ownrepeated := Repeat(char, 2)
	builtinrepeated := strings.Repeat(char, 2)
	if builtinrepeated != ownrepeated {
		t.Errorf("expected %q but got %q", builtinrepeated, ownrepeated)
	}
}

func TestRepeatSpaced(t *testing.T) {
	char := "a"
	repeated := Repeat("a ", 2)
	spaced := Spaced(repeated)
	expected := []string{char, char}
	if !reflect.DeepEqual(spaced, expected) {
		t.Errorf("expected %q but got %q", expected, spaced)
	}
}

func Spaced(s string) []string {
	return strings.Fields(s)
}

// faster repeat with buffer
// instead of string concatenation
func Repeat(char string, iter int) string {
	var buf bytes.Buffer
	for i := 0; i < iter; i++ {
		buf.WriteString(char)
	}
	return buf.String()
}

// faster repeat with stringbuilder
// instead of string concatenation
func RepeatStringBuiler(char string, iter int) string {
	var sb strings.Builder

	for i := 0; i < iter; i++ {
		sb.WriteString(char)
	}
	return sb.String()
}

// faster repeat with stringbuilder
// instead of string concatenation
func RepeatBuitlinStringRepeater(char string, iter int) string {
	return strings.Repeat(char, iter)
}

func RepeatRange(char string, iter int) string {
	var repeated string
	for range iter {
		repeated += char
	}
	return repeated
}

// func BenchmarkRepeatRange(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		RepeatRange("a", benchIter)
// 	}
// }

const benchIter = 50000

func BenchmarkRepeatBuffer(b *testing.B) {
	char := "a"
	for i := 0; i < b.N; i++ {
		Repeat(char, benchIter)
	}
}

func BenchmarkBuiltinStringRepeatBuilder(b *testing.B) {
	char := "a"
	for i := 0; i < b.N; i++ {
		RepeatStringBuiler(char, benchIter)
	}
}

func BenchmarkBuiltinStringRepeat(b *testing.B) {
	char := "a"
	for i := 0; i < b.N; i++ {
		RepeatBuitlinStringRepeater(char, benchIter)
	}
}

func ExampleRepeat() {
	char := "a"
	result := Repeat(char, 2)
	fmt.Println(result)
	// output: aa
}
