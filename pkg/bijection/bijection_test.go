package bijection_test

import (
	"testing"

	"github.com/go-sink/sink/pkg/bijection"
)

func TestFunction(t *testing.T) {
	t.Run("it encodes a char", func(t *testing.T) {
		stringToEncode := 2
		want := "c"
		got := bijection.Encode(stringToEncode)

		if got != want {
			t.Errorf("doesn't encode correctly: want: `%v`, got: `%v`", want, got)
		}
	})
}
