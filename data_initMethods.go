package passwordGenerator

import "math/big"

func (d *Data) initAvailableRunes() {
	d.initAvailableRunesByIncludingUpperCaseLetters()
	d.initAvailableRunesByIncludingLowerCaseLetters()
	d.initAvailableRunesByIncludingDigits()
	d.initAvailableRunesByIncludingRunesList()
	d.initAvailableRunesByExcludingAmbiguousRunes()
	d.initAvailableRunesByExcludingRunesList()

	d.availableRunesLenBig = big.NewInt(int64(len(d.availableRunes)))
}

func (d *Data) initAvailableRunesByIncludingUpperCaseLetters() {
	if d.options.IncludeUpperCaseLetters {
		for r := 'A'; r <= 'Z'; r++ {
			d.availableRunes = append(d.availableRunes, r)
		}
	}
}

func (d *Data) initAvailableRunesByIncludingLowerCaseLetters() {
	if d.options.IncludeLowerCaseLetters {
		for r := 'a'; r <= 'z'; r++ {
			d.availableRunes = append(d.availableRunes, r)
		}
	}
}

func (d *Data) initAvailableRunesByIncludingDigits() {
	if d.options.IncludeDigits {
		for r := '0'; r <= '9'; r++ {
			d.availableRunes = append(d.availableRunes, r)
		}
	}
}

func (d *Data) initAvailableRunesByIncludingRunesList() {
	for _, r := range d.options.IncludeRunesList {
		var alreadyFound bool

		for _, r2 := range d.availableRunes {
			if r2 == r {
				alreadyFound = true
				break
			}
		}

		if !alreadyFound {
			d.availableRunes = append(d.availableRunes, r)
		}
	}
}

func (d *Data) initAvailableRunesByExcludingAmbiguousRunes() {
	if d.options.ExcludeAmbiguousRunes {
		for _, r := range dataAmbiguousRunes {
			for i, r2 := range d.availableRunes {
				if r2 == r {
					d.availableRunes = append(d.availableRunes[:i], d.availableRunes[i+1:]...)
					break
				}
			}
		}
	}
}

func (d *Data) initAvailableRunesByExcludingRunesList() {
	for _, r := range d.options.ExcludeRunesList {
		for i, r2 := range d.availableRunes {
			if r2 == r {
				d.availableRunes = append(d.availableRunes[:i], d.availableRunes[i+1:]...)
				break
			}
		}
	}
}
