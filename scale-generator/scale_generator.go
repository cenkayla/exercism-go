package scale

import "strings"

func Scale(tonic string, interval string) (res []string) {

	var scale []string

	switch tonic {
	case "C", "a", "G", "D", "A", "E", "B", "F#", "e", "b", "f#", "c#", "g#", "d#":
		scale = []string{"A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#"}
	case "F", "Bb", "Eb", "Ab", "Db", "Gb", "d", "g", "c", "f", "bb", "eb":
		scale = []string{"A", "Bb", "B", "C", "Db", "D", "Eb", "E", "F", "Gb", "G", "Ab"}
	}
	for i, s := range scale {
		if s == strings.Title(tonic) {
			scale = append(scale[i:], scale[:i]...)
		}
	}

	if interval == "" {
		interval = "mmmmmmmmmmmm"
	}
	for _, s := range interval {
		res = append(res, scale[0])

		switch s {
		case rune('m'):
			scale = scale[1:]
		case rune('M'):
			scale = scale[2:]
		case rune('A'):
			scale = scale[3:]
		}
	}

	return res
}