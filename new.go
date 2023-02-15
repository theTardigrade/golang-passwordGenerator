package passwordGenerator

// New returns a pointer to a Data struct that is used
// to generate passwords. An options argument determines
// exactly how the passwords should be generated.
func New(options Options) *Data {
	d := &Data{
		options: &options,
	}

	d.initAvailableRunes()

	return d
}
