package reverse

import "unicode/utf8"


func Reverse(s string) string {

	buf := make([]byte, len(s))

	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])
		i += size
		utf8.EncodeRune(buf[len(s)-i:], r)
	}
	return string(buf)

}
