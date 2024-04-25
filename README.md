## nysiis

[![Release](https://img.shields.io/github/release/0xnu/nysiis.svg)](https://github.com/0xnu/nysiis/releases/latest)
[![Go Report Card](https://goreportcard.com/badge/github.com/0xnu/nysiis)](https://goreportcard.com/report/github.com/0xnu/nysiis)
[![Go Reference](https://pkg.go.dev/badge/github.com/0xnu/nysiis.svg)](https://pkg.go.dev/github.com/0xnu/nysiis)
[![PyPI version](https://img.shields.io/pypi/v/pynysiis.svg)](https://pypi.org/project/pynysiis/)
[![npm version](https://img.shields.io/npm/v/nysiis.svg)](https://www.npmjs.com/package/nysiis)
[![LuaRocks](https://img.shields.io/luarocks/v/0xnu/nysiis.svg)](https://luarocks.org/modules/0xnu/nysiis)
[![License](https://img.shields.io/github/license/0xnu/nysiis)](/LICENSE)

The `nysiis` package provides a Golang implementation of the [New York State Identification and Intelligence System](https://en.wikipedia.org/wiki/New_York_State_Identification_and_Intelligence_System) (NYSIIS) phonetic encoding algorithm. NYSIIS encodes names based on pronunciation, which is helpful in name-matching and searching applications.

### Installation

```sh
go install github.com/0xnu/nysiis
```

### Usage

```go
package main

import (
	"fmt"

	"github.com/0xnu/nysiis"
)

func main() {
	// Create a new instance of the NYSIIS encoder
	encoder := nysiis.NewNysiis()

	// Example names to encode
	names := []string{
		"Watkins",
		"Robert Johnson",
		"Samantha Williams",
		"Olanrewaju Akinyele", // Yoruba
		"Obinwanne Obiora", // Igbo
		"Abdussalamu Abubakar", // Hausa
		"Virat Kohli", // Hindi
		"Usman Shah", // Urdu
	}

	// Encode each name using NYSIIS
	for _, name := range names {
		encodedName := encoder.Encode(name)
		fmt.Printf("Encoded name for %q: %s\n", name, encodedName)
	}
}
```

### Test

To test `nysiis`, ensure that you have Go installed on your system. Then, follow these steps:

```sh
go test ./...
```

### Reference

```tex
@inproceedings{Rajkovic2007,
  author    = {Petar Rajkovic and Dragan Jankovic},
  title     = {Adaptation and Application of Daitch-Mokotoff Soundex Algorithm on Serbian Names},
  booktitle = {XVII Conference on Applied Mathematics},
  editors   = {D. Herceg and H. Zarin},
  pages     = {193--204},
  year      = {2007},
  publisher = {Department of Mathematics and Informatics, Novi Sad},
  url       = {https://jmp.sh/hukNujCG}
}
```

### Additional References

+ [Commission Implementing Regulation (EU) 2016/480](https://www.legislation.gov.uk/eur/2016/480/contents)
+ [Commission Implementing Regulation (EU) 2023/2381](https://eur-lex.europa.eu/eli/reg_impl/2023/2381/oj)

### License

This project is licensed under the [MIT License](./LICENSE).

### Copyright

(c) 2024 [Finbarrs Oketunji](https://finbarrs.eu).