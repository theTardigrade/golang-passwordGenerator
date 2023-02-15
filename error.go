package passwordGenerator

import "errors"

// ErrNoAvailableRunes is used if the number of runes that can be randomly selected
// when generating passwords is zero.
var ErrNoAvailableRunes = errors.New("no runes are available for password-generation")
