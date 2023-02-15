# golang-passwordGenerator

This package allows cryptographically random passwords to be generated.

[![Go Reference](https://pkg.go.dev/badge/github.com/theTardigrade/golang-passwordGenerator.svg)](https://pkg.go.dev/github.com/theTardigrade/golang-passwordGenerator) [![Go Report Card](https://goreportcard.com/badge/github.com/theTardigrade/golang-passwordGenerator)](https://goreportcard.com/report/github.com/theTardigrade/golang-passwordGenerator)

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
				'<', '>', '+', '/', '*', '\\', '/',
				':', ';', '&', '\'', '"', '%', '^',
				'🙂', '🙃',
			},
			ExcludeAmbiguousRunes: true,
			ExcludeRunesList:      []rune{'X', 'x'},
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

## Support

If you use this package, or find any value in it, please consider donating:

[![ko-fi](https://ko-fi.com/img/githubbutton_sm.svg)](https://ko-fi.com/S6S2EIRL0)