package strand

import "strings"

func ToRNA(dna string) string {
	d := make([]string,len(dna))
 	for i := range dna {
		switch dna[i] {
		case 'G':
			d[i] = "C"
		case 'C':
			d[i] = "G"
		case 'T':
			d[i] = "A"
		case 'A':
			d[i] = "U"
		}
	} 
	return strings.Join(d,"")
}
 