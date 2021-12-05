package passwordGenerator

type Data struct {
	availableRunes []rune
	options        *Options
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
