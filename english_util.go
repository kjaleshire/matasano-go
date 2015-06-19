package matasano

func stringScore(s string) (score uint) {
	for _, c := range s {
		switch {
		case c >= 'A' && c <= 'Z':
			score++
		case c >= 'a' && c <= 'z':
			score++
		case c >= '0' && c <= '9':
			score++
		}
		switch c {
		case ' ', '-', '\'', '\n', '/', ',', '.', '?', '!':
			score++
		}
	}
	return
}
