package passwordGenerator

type Options struct {
	Len                     int
	IncludeUpperCaseLetters bool
	IncludeLowerCaseLetters bool
	IncludeDigits           bool
	RemoveAmbiguousRunes    bool
	IncludeRunesList        []rune
	RemoveRunesList         []rune
}
