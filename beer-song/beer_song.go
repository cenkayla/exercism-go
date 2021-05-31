package beer

import (
	"errors"
	"fmt"
	"strings"
)

func Verse(num int) (string, error) {
 	switch {
	case num > 100 || num < 0:
		return "", errors.New("invalid verse")
	case num == 0:
		return "No more bottles of beer on the wall, no more bottles of beer.\nGo to the store and buy some more, 99 bottles of beer on the wall.\n", nil
	case num == 2:
		return fmt.Sprintf("%d bottles of beer on the wall, %d bottles of beer.\nTake one down and pass it around, %d bottle of beer on the wall.\n", num, num, num-1), nil
	case num == 1:
		return fmt.Sprintf("%d bottle of beer on the wall, %d bottle of beer.\nTake it down and pass it around, no more bottles of beer on the wall.\n", num, num), nil
	default:
		return fmt.Sprintf("%d bottles of beer on the wall, %d bottles of beer.\nTake one down and pass it around, %d bottles of beer on the wall.\n", num, num, num-1), nil
	} 
}

func Verses(upper, lower int) (string, error) {
	if upper <= lower {
		return "", errors.New("invalid stop")
	}
	if upper > 100 || lower < 0 {
		return "", errors.New("invalid start of invalid stop")
	}
	var sb strings.Builder

	for i := upper; i >= lower; i-- {
		v, _ := Verse(i)
		sb.WriteString(v + "\n")
	}
	return sb.String(), nil
}

func Song() string {
	song, _ := Verses(99, 0)
	return song
}
