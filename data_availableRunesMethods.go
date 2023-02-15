package passwordGenerator

// AvailableRunesCount returns the number of runes that can be randomly selected
// when generating passwords.
// These runes are determined by the options contained in the Data struct.
// If the number of runes is zero, no passwords can be successfully generated.
func (d *Data) AvailableRunesCount() int {
	return len(d.availableRunes)
}

// AvailableRunes returns a freshly allocated slice containing all of the
// runes that can be randomly selected when generating passwords.
// These runes are determined by the options contained in the Data struct.
func (d *Data) AvailableRunes() []rune {
	clonedAvailableRunes := make([]rune, len(d.availableRunes))

	copy(clonedAvailableRunes, d.availableRunes)

	return clonedAvailableRunes
}
