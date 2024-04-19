package example

import (
	"fmt"

	"github.com/0xnu/nysiis"
)

func EncodeNames() {
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
