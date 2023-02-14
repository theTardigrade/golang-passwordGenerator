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
