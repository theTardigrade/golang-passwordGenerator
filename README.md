# golang-passwordGenerator

This package allows cryptographically random passwords to be generated.

## Example

```golang
package main

import (
	"fmt"

	passwordGenerator "github.com/theTardigrade/golang-passwordGenerator"
)

func main() {
	pg := passwordGenerator.New(
		passwordGenerator.Options{
			Len:                     128,
			IncludeUpperCaseLetters: true,
			IncludeLowerCaseLetters: true,
			IncludeDigits:           true,
			IncludeRunesList: []rune{
				'!', '?', '-', '_', '=', '@', '$',
				'#', '(', ')', '[', ']', '{', '}',
			},
			RemoveAmbiguousRunes: true,
			RemoveRunesList:      []rune{'X', 'x'},
		},
	)

	passwords, err := pg.GenerateMany(5)
	if err != nil {
		panic(err)
	}

	for _, p := range passwords {
		fmt.Println(p)
	}
}
```