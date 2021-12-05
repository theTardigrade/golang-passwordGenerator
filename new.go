package passwordGenerator

func New(options Options) *Data {
	d := &Data{
		options: &options,
	}

	d.initAvailableRunes()

	return d
}
