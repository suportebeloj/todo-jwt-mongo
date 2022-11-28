package validator

import (
	"github.com/dlclark/regexp2"
)

func PasswordValidator(plainPassword string) bool {
	validationPattern := "(?=.*[a-z])(?=.*[A-Z])(?=.*[0-9])(?=.*[\\!-\\/\\:-\\@\\[-\\`\\{-\\~])(?:.+)"
	r := regexp2.MustCompile(validationPattern, regexp2.RE2)

	if m, _ := r.FindStringMatch(plainPassword); m != nil {
		return true
	} else {
		return false
	}
}
