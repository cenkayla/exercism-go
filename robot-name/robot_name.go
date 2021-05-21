package robotname

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type Robot struct {
	name string
}

const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const maxName = 26 * 26 * 10 * 10 * 10

var usedNames = map[string]bool{}

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func Random(min int, max int) int {
	return seededRand.Intn(max-min) + min

}

func StringWithCharset() string {
	b := make([]byte, 2)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b) + fmt.Sprint(Random(100, 999))
}

func (r *Robot) Name() (string, error) {
	if r.name != "" {
		return r.name, nil
	}

	if len(usedNames) == maxName {
		return "", errors.New("error")
	}

	r.name = StringWithCharset()
	//robot name tekrar ederse yenisi üretiliyor
	for usedNames[r.name] {
		r.name = StringWithCharset()
	}

	usedNames[r.name] = true
	return r.name, nil

}

func (r *Robot) Reset() {
	r.name = ""
}
