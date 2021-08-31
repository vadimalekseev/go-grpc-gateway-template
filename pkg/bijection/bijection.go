package bijection

const alphabet = "abcdefghijklmnopqrstuvwxyz0123456789"
const base = len(alphabet)

func Encode(numberToEncode int) (encodedString string) {
	if numberToEncode == 0 {
		return string(alphabet[0])
	}

	for numberToEncode > 0 {
		encodedString = encodedString + string(alphabet[numberToEncode%base])
		numberToEncode = numberToEncode / base
	}

	return Reverse(encodedString)
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
