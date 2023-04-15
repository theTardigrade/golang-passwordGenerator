package passwordGenerator

import "math/big"

// Data is used to generate passwords, by calling the appropriate
// method on it. The New function is used to create a pointer
// to a Data struct.
type Data struct {
	availableRunes       []rune
	availableRunesLenBig *big.Int
	options              *Options
}

var (
	dataAmbiguousRunes = [...]rune{
		'A', '4',
		'B', '8',
		'g', '9',
		'I', '1',
		'K', 'X',
		'o', 'O', 'D', '0',
		's', 'S', '5',
		'T', '7',
		'u', 'v', 'U', 'V',
		'Z', '2',
	}
)
