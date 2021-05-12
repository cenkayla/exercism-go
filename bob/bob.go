package bob

import (
	"strings"
	"unicode"
)

//1-Cümlenin sonunda (?) olursa "Sure"
//2-Hepsi büyük harf olursa "Whoa, chill out!"
//3-Hepsi büyük harf ve soru sorulursa (?) "Calm down, I know what I'm doing!"
//4-Boş tab vs olursa "Fine. Be that way!"
//5-Yukarıdaki durumların dışında olan her şey "Whatever"
func Hey(remark string) string {

	// Değişkenin başında ve sonunda bulunan boşlukları kaldırır.
	remark = strings.TrimSpace(remark)

	if remark == "" {
		return "Fine. Be that way!"
	}

	switch {
	case strings.HasSuffix(remark, "?") && allUpper(remark):
		return "Calm down, I know what I'm doing!"
	case strings.HasSuffix(remark, "?"):
		return "Sure."
	case allUpper(remark):
		return "Whoa, chill out!"
	default:
		return "Whatever."
	}
}

//Harflerin büyüklüğü kontrolu
func allUpper(s string) bool {
	if !hasLetter(s) {
		return false
	}

	for _, v := range s {
		if unicode.IsLower(v) {
			return false
		}
	}
	return true
}

//Değişkene gelen değerde harf var mı yok mu kontrolu
func hasLetter(s string) bool {
	for _, v := range s {
		if unicode.IsLetter(v) {
			return true
		}
	}
	return false
}
