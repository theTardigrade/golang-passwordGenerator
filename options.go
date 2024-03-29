package passwordGenerator

// Options is passed to the New function to determine
// exactly how any passwords should be generated by the
// Data struct.
// The Len field determines the length of any passwords
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
