package bijection_test

import (
	"github.com/go-sink/sink/internal/pkg/bijection"
	"testing"
)

func TestFunction(t *testing.T) {
	t.Run("it encodes an int", func(t *testing.T) {
		stringToEncode := 2
		want := "c"
		got := bijection.Encode(stringToEncode)

		if got != want {
			t.Errorf("doesn't encode correctly: want: `%v`, got: `%v`", want, got)
		}
	})

	t.Run("it decodes a string", func(t *testing.T) {
		intToEncode := "a"
		want := 0
		got := bijection.Decode(intToEncode)

			if got != want {
				t.Errorf("doesn't decode correctly: want: `%v`, got: `%v`", want, got)
			}
	})
}
