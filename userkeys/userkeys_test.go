// Package userkeys manages acces to users' private keys
package userkeys

import (
	"testing"

	"github.com/kappapay/backend/lib/golang/src/kappa/common/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_entityToKeyRelPath(t *testing.T) {
	type args struct {
		entity *models.Entity
	}
	tests := map[string]struct {
		args    args
		want    string
		wantErr bool
	}{
		"invalid_uuid": {
			args:    args{entity: &models.Entity{Id: "not a valid uuid"}},
			wantErr: true,
		},
		"valid_uuid": {
			args:    args{entity: &models.Entity{Id: "00000001-0000-0001-0000-000100000001"}},
			want:    "1/1/1/1",
			wantErr: false,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := entityToKeyRelPath(tt.args.entity)
			if (err != nil) != tt.wantErr {
				t.Errorf("entityToKeyRelPath() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("entityToKeyRelPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_GetEntityAddressAsString(t *testing.T) {
	seed := []byte("12345678901234567890123456789012")
	entityID := "00000001-0000-0001-0000-000100000001"
	entity := &models.Entity{Id: entityID}
	expectedAddress := "0x0Ba99F6F261f26Aa587112676fdA499E0E8bF759"
	expectedPrivateKey := "4a6bb377c1eb2eba8bd8660f9138f3b3b8c1bf87f53c10a313ff5a9d5ede2b9e"

	wallet, err := NewHDWallet(seed)
	require.Nilf(t, err, "couldn't set up hd wallet for test")
	require.NotNilf(t, wallet, "wallet must not be nil for test")

	address, err := GetEntityAddressAsString(wallet, entity)
	assert.Nilf(t, err, "not expecting error when generating address")
	assert.Equalf(t, expectedAddress, address, "derived address has changed")

	privateKey, err := GetEntityPrivateKeyAsString(wallet, entity)
	assert.Nilf(t, err, "not expecting error when generating private key")
	assert.Equalf(t, expectedPrivateKey, privateKey, "derived private key has changed")
}

func Benchmark_GetEntityPrivateKeyAsString(b *testing.B) {
	seed := []byte("12345678901234567890123456789012")
	entityID := "00000001-0000-0001-0000-000100000001"
	entity := &models.Entity{Id: entityID}

	wallet, err := NewHDWallet(seed)
	require.Nilf(b, err, "couldn't set up hd wallet for test")
	require.NotNilf(b, wallet, "wallet must not be nil for test")

	b.StartTimer()

	for i := 0; i < b.N; i++ {
		_, err := GetEntityPrivateKeyAsString(wallet, entity)
		require.Nil(b, err)
	}
}
