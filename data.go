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
		'o', 'O', 'D', '0',
		'I', '1',
		'K', 'X',
		's', 'S', '5',
		'u', 'v', 'U', 'V',
		'Z', '2',
		'g', '9',
	}
)
