package ptr

import (
	"testing"
)

func TestPtr(t *testing.T) {
	t.Parallel()

	tests := map[string]any{
		"string": "anything",
		"int":    200,
		"bool":   false,
	}

	for name, input := range tests {
		input := input
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			// Act
			result := Ptr(input)

			// Assert
			if input != *result {
				t.Error("not equal")
			}
		})
	}
}

func BenchmarkPtr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Ptr(true)
	}
}

func BenchmarkAmpersand(b *testing.B) {
	for i := 0; i < b.N; i++ {
		variable := true
		_ = &variable
	}
}
