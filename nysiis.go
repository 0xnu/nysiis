package nysiis

import (
	"regexp"
	"strings"
)

type Nysiis struct {
	vowels map[rune]bool
}

func NewNysiis() *Nysiis {
	return &Nysiis{
		vowels: map[rune]bool{
			'A': true,
			'E': true,
			'I': true,
			'O': true,
			'U': true,
		},
	}
}

func (n *Nysiis) preprocessName(name string) string {
	name = strings.ToUpper(name)
	name = regexp.MustCompile(`[^A-Z]`).ReplaceAllString(name, "")
	return name
}

func (n *Nysiis) translateFirstCharacters(name string) string {
	switch {
	case strings.HasPrefix(name, "MAC"):
		name = "MCC" + name[3:]
	case strings.HasPrefix(name, "KN"):
		name = "NN" + name[2:]
	case strings.HasPrefix(name, "K"):
		name = "C" + name[1:]
	case strings.HasPrefix(name, "PH"):
		name = "FF" + name[2:]
	case strings.HasPrefix(name, "PF"):
		name = "FF" + name[2:]
	case strings.HasPrefix(name, "SCH"):
		name = "SSS" + name[3:]
	case strings.HasPrefix(name, "GB"):
		name = "J" + name[2:] // Igbo: 'Gb' -> 'J'
	case strings.HasPrefix(name, "KP"):
		name = "P" + name[2:] // Igbo: 'Kp' -> 'P'
	case strings.HasPrefix(name, "NW"):
		name = "W" + name[2:] // Igbo: 'Nw' -> 'W'
	case strings.HasPrefix(name, "TS"):
		name = "S" + name[2:] // Yoruba: 'Ts' -> 'S'
	case strings.HasPrefix(name, "SH"):
		name = "S" + name[2:] // Hausa: 'Sh' -> 'S'
	case strings.HasPrefix(name, "BH"):
		name = "B" + name[2:] // Hindi: 'Bh' -> 'B'
	case strings.HasPrefix(name, "DH"):
		name = "D" + name[2:] // Hindi: 'Dh' -> 'D'
	case strings.HasPrefix(name, "GH"):
		name = "G" + name[2:] // Hindi: 'Gh' -> 'G'
	case strings.HasPrefix(name, "JH"):
		name = "J" + name[2:] // Hindi: 'Jh' -> 'J'
	case strings.HasPrefix(name, "KH"):
		name = "K" + name[2:] // Hindi: 'Kh' -> 'K'
	case strings.HasPrefix(name, "PH"):
		name = "F" + name[2:] // Hindi: 'Ph' -> 'F'
	case strings.HasPrefix(name, "TH"):
		name = "T" + name[2:] // Hindi: 'Th' -> 'T'
	}
	return name
}

func (n *Nysiis) translateLastCharacters(name string) string {
	switch {
	case strings.HasSuffix(name, "EE"), strings.HasSuffix(name, "IE"):
		name = name[:len(name)-2] + "Y"
	case strings.HasSuffix(name, "DT"), strings.HasSuffix(name, "RT"), strings.HasSuffix(name, "RD"), strings.HasSuffix(name, "NT"), strings.HasSuffix(name, "ND"):
		name = name[:len(name)-2] + "D"
	}
	return name
}

func (n *Nysiis) generateKey(name string) string {
	key := string(name[0])
	var prevChar rune = rune(name[0])

	for i := 1; i < len(name); i++ {
		char := rune(name[i])
		if n.vowels[char] {
			char = 'A'
		}

		char = n.translateChar(char, name, i)
		char = n.handleVowelHarmony(char, prevChar)
		char = n.ignoreTonalDifferences(char)

		if char != prevChar {
			key += string(char)
		}

		prevChar = char
	}

	key = n.removeTrailingS(key)
	key = n.translateAY(key)
	key = n.removeTrailingA(key)
	key = n.truncateKey(key)

	return key
}

func (n *Nysiis) translateChar(char rune, name string, i int) rune {
	if char == 'E' && i+1 < len(name) && name[i+1] == 'V' {
		char = 'A'
	} else if char == 'Q' {
		char = 'G'
	} else if char == 'Z' {
		char = 'S'
	} else if char == 'M' {
		char = 'N'
	} else if char == 'K' {
		if i+1 < len(name) && name[i+1] == 'N' {
			char = rune(name[i])
		} else {
			char = 'C'
		}
	} else if char == 'S' && i+2 < len(name) && name[i:i+3] == "SCH" {
		char = 'S'
	} else if char == 'P' && i+1 < len(name) && name[i+1] == 'H' {
		char = 'F'
	} else if char == 'H' && (i == 0 || i+1 == len(name) || !n.vowels[rune(name[i-1])] || !n.vowels[rune(name[i+1])]) {
		char = rune(name[i-1])
	} else if char == 'W' && i > 0 && n.vowels[rune(name[i-1])] {
		char = rune(name[i-1])
	} else if char == 'G' && i+1 < len(name) && name[i+1] == 'B' {
		char = 'J' // Igbo: 'Gb' -> 'J'
	} else if char == 'K' && i+1 < len(name) && name[i+1] == 'P' {
		char = 'P' // Igbo: 'Kp' -> 'P'
	} else if char == 'N' && i+1 < len(name) && name[i+1] == 'W' {
		char = 'W' // Igbo: 'Nw' -> 'W'
	} else if char == 'T' && i+1 < len(name) && name[i+1] == 'S' {
		char = 'S' // Yoruba: 'Ts' -> 'S'
	} else if char == 'S' && i+1 < len(name) && name[i+1] == 'H' {
		char = 'S' // Hausa: 'Sh' -> 'S'
	} else if char == 'B' && i+1 < len(name) && name[i+1] == 'H' {
		char = 'B' // Hindi: 'Bh' -> 'B'
	} else if char == 'D' && i+1 < len(name) && name[i+1] == 'H' {
		char = 'D' // Hindi: 'Dh' -> 'D'
	} else if char == 'G' && i+1 < len(name) && name[i+1] == 'H' {
		char = 'G' // Hindi: 'Gh' -> 'G'
	} else if char == 'J' && i+1 < len(name) && name[i+1] == 'H' {
		char = 'J' // Hindi: 'Jh' -> 'J'
	} else if char == 'K' && i+1 < len(name) && name[i+1] == 'H' {
		char = 'K' // Hindi: 'Kh' -> 'K'
	} else if char == 'P' && i+1 < len(name) && name[i+1] == 'H' {
		char = 'F' // Hindi: 'Ph' -> 'F'
	} else if char == 'T' && i+1 < len(name) && name[i+1] == 'H' {
		char = 'T' // Hindi: 'Th' -> 'T'
	}

	return char
}

func (n *Nysiis) handleVowelHarmony(char, prevChar rune) rune {
	if n.vowels[char] && n.vowels[prevChar] {
		if prevChar == 'A' || prevChar == 'O' || prevChar == 'U' {
			if char == 'E' || char == 'I' {
				char = 'A'
			}
		} else if prevChar == 'E' || prevChar == 'I' {
			if char == 'A' || char == 'O' || char == 'U' {
				char = 'E'
			}
		}
	}
	return char
}

func (n *Nysiis) ignoreTonalDifferences(char rune) rune {
	if char >= 'A' && char <= 'Z' {
		char = rune(strings.ToUpper(string(char))[0])
	}
	return char
}

func (n *Nysiis) removeTrailingS(key string) string {
	if len(key) > 1 && strings.HasSuffix(key, "S") {
		key = key[:len(key)-1]
	}
	return key
}

func (n *Nysiis) translateAY(key string) string {
	if strings.HasSuffix(key, "AY") {
		key = key[:len(key)-2] + "Y"
	}
	return key
}

func (n *Nysiis) removeTrailingA(key string) string {
	if len(key) > 1 && strings.HasSuffix(key, "A") {
		key = key[:len(key)-1]
	}
	return key
}

func (n *Nysiis) truncateKey(key string) string {
	if len(key) > 6 {
		key = key[:6]
	}
	return key
}

func (n *Nysiis) Encode(name string) string {
	if name == "" {
		return ""
	}

	name = n.preprocessName(name)

	if len(name) < 2 {
		return name
	}

	name = n.translateFirstCharacters(name)
	name = n.translateLastCharacters(name)
	key := n.generateKey(name)

	return key
}
