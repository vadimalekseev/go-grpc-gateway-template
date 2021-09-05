package bijection

import (
	"strings"

	aux "github.com/go-sink/sink/internal/pkg"
)

type EncodingAlgorithm interface {
	Encode(rune) string
	Decode(string) rune
}

type numberSystemConverter struct {

}

func NewNumberSystemConverter() *numberSystemConverter {
	return &numberSystemConverter{}
}

const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNIPQRSTUVWXYZ0123456789"
const base = rune(len(alphabet))

func (n *numberSystemConverter) Encode(toEncode rune) (encoded string) {
	if toEncode == 0 {
		return string(alphabet[0])
	}

	for toEncode > 0 {
		encoded = encoded + string(alphabet[toEncode%base])
		toEncode = toEncode / base
	}

	return aux.Reverse(encoded)
}

func (n *numberSystemConverter) Decode(encoded string) (decoded rune) {
	for _, char := range encoded {
		decoded = (decoded * base) + rune(strings.Index(alphabet, string(char)))
	}
	return
}
