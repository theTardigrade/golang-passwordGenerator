package passwordGenerator

// Options is passed to the New function to determine
// exactly how passwords should be generated.
// The field Len determines the length of any passwords
// that are generated, while the other fields determine
// which runes are available to be chosen at random
// when generating passwords.
type Options struct {
	Len                     int
	IncludeUpperCaseLetters bool
	IncludeLowerCaseLetters bool
	IncludeDigits           bool
	ExcludeAmbiguousRunes   bool
	IncludeRunesList        []rune
	ExcludeRunesList        []rune
}
