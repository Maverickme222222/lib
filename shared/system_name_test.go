package shared_test

import (
	"database/sql/driver"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/kappapay/backend/lib/golang/src/kappa/shared"
)

var (
	invalidSystemName = shared.SystemName("NOT A SYSTEMNAME")

	systemNameList = []shared.SystemName{
		shared.SystemWallets,
		shared.SystemTransfers,
		shared.SystemSila,
		shared.SystemQuotes,
		shared.SystemPhones,
		shared.SystemNotifications,
	}
)

func TestSystemName_Scan(t *testing.T) {
	typeErr := invalidSystemName.Scan(math.MaxInt8)
	require.Error(t, typeErr)
	assert.Contains(t, typeErr.Error(), "is not a string")

	invalidErr := invalidSystemName.Scan(invalidSystemName.String())
	assert.Equal(t, invalidSystemName.Validate(), invalidErr)

	for _, expected := range systemNameList {
		actual := new(shared.SystemName)
		require.NoError(t, actual.Scan(expected.String()))
		assert.Equal(t, expected, *actual)
	}
}

func TestSystemName_Validate(t *testing.T) {
	for _, systemName := range systemNameList {
		assert.NoError(t, systemName.Validate())
	}

	assert.Error(t, invalidSystemName.Validate())
}

func TestSystemName_Value(t *testing.T) {
	invalidValue, invalidErr := invalidSystemName.Value()
	assert.Nil(t, invalidValue)
	assert.Equal(t, invalidSystemName.Validate(), invalidErr)

	for _, systemName := range systemNameList {
		actual, actualErr := systemName.Value()
		assert.Nil(t, actualErr)

		expected := driver.Value(string(systemName))
		assert.Equal(t, expected, actual)
	}
}
