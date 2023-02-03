package shared_test

import (
	"fmt"
	"testing"

	"github.com/kappapay/backend/lib/golang/src/kappa/shared"
	"github.com/stretchr/testify/assert"
)

func TestGenerateRandString(t *testing.T) {
	tests := []struct {
		name   string
		length int
		err    error
	}{
		{
			name:   "Negative length",
			length: -1,
			err:    fmt.Errorf("length should be greater than 0"),
		},
		{
			name:   "Zero length",
			length: 0,
			err:    fmt.Errorf("length should be greater than 0"),
		},
		{
			name:   "Postive length",
			length: 10,
			err:    nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output, err := shared.GenerateRandString(test.length)
			if test.err != nil {
				assert.NotNil(t, err)
				assert.Equal(t, test.err.Error(), err.Error())
			} else {
				assert.Nil(t, err)
				assert.Equal(t, len(output), test.length)
			}
		})
	}
}
