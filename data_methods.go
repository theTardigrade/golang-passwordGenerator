package passwordGenerator

import (
	"crypto/rand"
	"math/big"
	"strings"
	"sync"
	"unicode"
)

func (d *Data) initAvailableRunes() {
	if d.options.IncludeUpperCaseLetters {
		for r := 'A'; r <= 'Z'; r++ {
			d.availableRunes = append(d.availableRunes, r)
		}
	}

	if d.options.IncludeLowerCaseLetters {
		for r := 'a'; r <= 'z'; r++ {
			d.availableRunes = append(d.availableRunes, unicode.ToLower(r))
		}
	}

	if d.options.IncludeDigits {
		for r := '0'; r <= '9'; r++ {
			d.availableRunes = append(d.availableRunes, r)
		}
	}

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

	if d.options.RemoveAmbiguousRunes {
		for _, r := range dataAmbiguousRunes {
			for i, r2 := range d.availableRunes {
				if r2 == r {
					d.availableRunes = append(d.availableRunes[:i], d.availableRunes[i+1:]...)
					break
				}
			}
		}
	}

	for _, r := range d.options.RemoveRunesList {
		for i, r2 := range d.availableRunes {
			if r2 == r {
				d.availableRunes = append(d.availableRunes[:i], d.availableRunes[i+1:]...)
				break
			}
		}
	}
}

func (d *Data) generateRune() (r rune, err error) {
	len64 := int64(len(d.availableRunes))
	lenBig := big.NewInt(len64)

	iBig, err := rand.Int(rand.Reader, lenBig)
	if err != nil {
		return
	}

	i := iBig.Int64()
	r = d.availableRunes[i]

	return
}

func (d *Data) Generate() (password string, err error) {
	var builder strings.Builder
	var r rune

	for i := d.options.Len - 1; i >= 0; i-- {
		r, err = d.generateRune()
		if err != nil {
			return
		}

		builder.WriteRune(r)
	}

	password = builder.String()

	return
}

const (
	dataGenerateManyBatchSize = 128
)

func (d Data) generateManyBatch(passwords []string, start int, end int) (err error) {
	for i := start; i < end; i++ {
		var p string

		p, err = d.Generate()
		if err != nil {
			return
		}

		passwords[i] = p
	}

	return
}
func (d Data) GenerateMany(n int) (passwords []string, err error) {
	passwords = make([]string, n)

	if n <= dataGenerateManyBatchSize {
		d.generateManyBatch(passwords, 0, n)
	} else {
		var wg sync.WaitGroup

		wholeBatches := n / dataGenerateManyBatchSize

		wg.Add(wholeBatches)

		for i := 0; i < wholeBatches; i++ {
			go func(i int) {
				defer wg.Done()

				start := i * dataGenerateManyBatchSize
				end := start + dataGenerateManyBatchSize

				d.generateManyBatch(passwords, start, end)
			}(i)
		}

		if start := wholeBatches * dataGenerateManyBatchSize; wholeBatches < start {
			wg.Add(1)

			go func() {
				defer wg.Done()

				d.generateManyBatch(passwords, start, n)
			}()
		}

		wg.Wait()
	}

	return
}
