package validator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGivenTrue_WhenTestAValidPassword(t *testing.T) {
	pwd := "}Y,y4wZEvLQgDT7"
	isValid := PasswordValidator(pwd)
	assert.True(t, isValid)
}

func TestGivenFalse_WhenTestInvalidPassword(t *testing.T) {
	pwd := "abc123"
	isValid := PasswordValidator(pwd)
	assert.False(t, isValid)
}
