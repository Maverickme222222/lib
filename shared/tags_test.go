package shared

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateAlphaNumeric(t *testing.T) {
	tests := map[string]struct {
		Input string
		Want  bool
	}{
		"Valid Input Underscore": {
			Input: "_hungai",
			Want:  true,
		},
		"Valid Input Dash": {
			Input: "hungai-",
			Want:  true,
		},
		"Invalid Input": {
			Input: "&hungai",
			Want:  false,
		},
		"Invalid Input Number": {
			Input: "123456",
			Want:  false,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			valid := ValidAlphaNumeric(test.Input)
			assert.Equal(t, valid, test.Want)
		})
	}
}
