package nysiis_test

import (
	"fmt"
	"testing"

	"github.com/0xnu/nysiis"
)

func TestNysiis_Encode(t *testing.T) {
	name1 := "Watkins"
	encodedName1 := nysiis.NewNysiis().Encode(name1)
	fmt.Printf("Encoded name for %q: %s\n", name1, encodedName1)

	name2 := "Robert Johnson"
	encodedName2 := nysiis.NewNysiis().Encode(name2)
	fmt.Printf("Encoded name for %q: %s\n", name2, encodedName2)

	name3 := "Samantha Williams"
	encodedName3 := nysiis.NewNysiis().Encode(name3)
	fmt.Printf("Encoded name for %q: %s\n", name3, encodedName3)

	name4 := "Olanrewaju Akinyele"
	encodedName4 := nysiis.NewNysiis().Encode(name4)
	fmt.Printf("Encoded name for %q: %s\n", name4, encodedName4)

	name5 := "Obinwanne Obiora"
	encodedName5 := nysiis.NewNysiis().Encode(name5)
	fmt.Printf("Encoded name for %q: %s\n", name5, encodedName5)

	name6 := "Abdussalamu Abubakar"
	encodedName6 := nysiis.NewNysiis().Encode(name6)
	fmt.Printf("Encoded name for %q: %s\n", name6, encodedName6)

	// Output:
	// Encoded name for "Watkins": WATCAN
	// Encoded name for "Robert Johnson": RABART
	// Encoded name for "Samantha Williams": SANANT
	// Encoded name for "Olanrewaju Akinyele": OLANRA
	// Encoded name for "Obinwanne Obiora": OBAWAN
	// Encoded name for "Abdussalamu Abubakar": ABDASA
}
