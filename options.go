package passwordGenerator

type Options struct {
	Len                     int
	IncludeUpperCaseLetters bool
	IncludeLowerCaseLetters bool
	IncludeDigits           bool
	ExcludeAmbiguousRunes   bool
	IncludeRunesList        []rune
	ExcludeRunesList        []rune
}
