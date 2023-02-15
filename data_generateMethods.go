package passwordGenerator

import (
	"crypto/rand"
	"strings"
	"sync"
)

func (d *Data) generateRune() (r rune, err error) {
	iBig, err := rand.Int(rand.Reader, d.availableRunesLenBig)
	if err != nil {
		return
	}

	i := iBig.Int64()
	r = d.availableRunes[i]

	return
}

// Generate creates a single password made up of random runes.
// The password will conform to the options contained in the Data struct.
// An error is also returned from Generate, which will be non-nil if
// the random-number generator is not working, meaning that a password
// cannot be succesfully generated, or if the options in the Data struct
// provide no runes to be randomly selected when generating passwords.
func (d *Data) Generate() (password string, err error) {
	if len(d.availableRunes) == 0 {
		err = ErrNoAvailableRunes
		return
	}

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

func (d *Data) generateManyBatchWithoutConcurrency(
	passwords []string,
	start int,
	end int,
) (err error) {
	for i := start; i < end; i++ {
		var p string

		if p, err = d.Generate(); err != nil {
			return
		}

		passwords[i] = p
	}

	return
}

func (d *Data) generateManyBatchWithConcurrency(
	passwords []string,
	start int,
	end int,
	wg *sync.WaitGroup,
	errChan chan error,
) {
	defer wg.Done()

	for i := start; i < end; i++ {
		select {
		case err := <-errChan:
			errChan <- err
			return
		default:
		}

		p, err := d.Generate()
		if err != nil {
			select {
			case errChan <- err:
			default:
			}
			return
		}

		passwords[i] = p
	}
}

// GenerateMany creates many passwords made up of random runes.
// The function's single argument will determine exactly how many
// passwords are generated.
// Any passwords will conform to the options contained in the Data struct.
// An error is also returned from GenerateMany, which will be non-nil if
// the random-number generator is not working, meaning that one of the
// passwords cannot be succesfully generated, or if the options in the Data struct
// provide no runes to be randomly selected when generating passwords.
func (d *Data) GenerateMany(n int) (passwords []string, err error) {
	passwords = make([]string, n)

	if n <= dataGenerateManyBatchSize {
		err = d.generateManyBatchWithoutConcurrency(passwords, 0, n)
	} else {
		var wg sync.WaitGroup

		wholeBatches := n / dataGenerateManyBatchSize

		wg.Add(wholeBatches)

		errChan := make(chan error, 1)

		for i := 0; i < wholeBatches; i++ {
			start := i * dataGenerateManyBatchSize
			end := start + dataGenerateManyBatchSize

			go d.generateManyBatchWithConcurrency(passwords, start, end, &wg, errChan)
		}

		if start := wholeBatches * dataGenerateManyBatchSize; start < n {
			wg.Add(1)

			go d.generateManyBatchWithConcurrency(passwords, start, n, &wg, errChan)
		}

		wg.Wait()

		select {
		case err = <-errChan:
		default:
		}
	}

	return
}
