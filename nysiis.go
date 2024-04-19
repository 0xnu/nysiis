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
		if char == 'E' && i+1 < len(name) && name[i+1] == 'V' {
			char = 'A'
			i++
		} else if char == 'Q' {
			char = 'G'
		} else if char == 'Z' {
			char = 'S'
		} else if char == 'M' {
			char = 'N'
		} else if char == 'K' {
			if i+1 < len(name) && name[i+1] == 'N' {
				continue
			} else {
				char = 'C'
			}
		} else if char == 'S' && i+2 < len(name) && name[i:i+3] == "SCH" {
			char = 'S'
			i += 2
		} else if char == 'P' && i+1 < len(name) && name[i+1] == 'H' {
			char = 'F'
			i++
		} else if char == 'H' && (prevChar != 'A' && prevChar != 'E' && prevChar != 'I' && prevChar != 'O' && prevChar != 'U' || (i+1 < len(name) && rune(name[i+1]) != 'A' && rune(name[i+1]) != 'E' && rune(name[i+1]) != 'I' && rune(name[i+1]) != 'O' && rune(name[i+1]) != 'U')) {
			char = prevChar
		} else if char == 'W' && (prevChar == 'A' || prevChar == 'E' || prevChar == 'I' || prevChar == 'O' || prevChar == 'U') {
			char = prevChar
		}

		if char != prevChar {
			key += string(char)
		}

		prevChar = char
	}

	if len(key) > 1 && strings.HasSuffix(key, "S") {
		key = key[:len(key)-1]
	}

	if strings.HasSuffix(key, "AY") {
		key = key[:len(key)-2] + "Y"
	}

	if len(key) > 1 && strings.HasSuffix(key, "A") {
		key = key[:len(key)-1]
	}

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
