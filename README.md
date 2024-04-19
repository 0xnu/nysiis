## nysiis

[![Release](https://img.shields.io/github/release/0xnu/nysiis.svg)](https://github.com/0xnu/nysiis/releases/latest)
[![Go Report Card](https://goreportcard.com/badge/github.com/0xnu/nysiis)](https://goreportcard.com/report/github.com/0xnu/nysiis)
[![Go Reference](https://pkg.go.dev/badge/github.com/0xnu/nysiis.svg)](https://pkg.go.dev/github.com/0xnu/nysiis)
[![License](https://img.shields.io/github/license/0xnu/nysiis)](/LICENSE)

The `nysiis` package provides a JavaScript implementation of the [New York State Identification and Intelligence System](https://en.wikipedia.org/wiki/New_York_State_Identification_and_Intelligence_System) (NYSIIS) phonetic encoding algorithm. NYSIIS encodes names based on pronunciation, which is helpful in name-matching and searching applications.

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

### License

This project is licensed under the [MIT License](./LICENSE).

### Copyright

(c) 2024 [Finbarrs Oketunji](https://finbarrs.eu).