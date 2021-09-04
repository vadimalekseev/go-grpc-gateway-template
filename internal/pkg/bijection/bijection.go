package bijection

import (
	"strings"

	aux "github.com/go-sink/sink/internal/pkg"
)

const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNIPQRSTUVWXYZ0123456789"
const base = len(alphabet)

func Encode(toEncode int) (encoded string) {
	if toEncode == 0 {
		return string(alphabet[0])
	}

	for toEncode > 0 {
		encoded = encoded + string(alphabet[toEncode%base])
		toEncode = toEncode / base
	}

	return aux.Reverse(encoded)
}

func Decode(encoded string) (decoded int) {
	for _, char := range encoded {
		decoded = (decoded * base) + strings.Index(alphabet, string(char))
	}
	return
}
