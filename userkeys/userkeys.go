// Package userkeys manages acces to users' private keys
package userkeys

import (
	"crypto/ecdsa"
	"encoding/binary"
	"encoding/hex"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/google/uuid"

	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"

	"github.com/kappapay/backend/lib/golang/src/kappa/common/models"
)

// HDWallet returns public and private keys for an entity
type HDWallet struct {
	wallet *hdwallet.Wallet
}

var _ KeysManager = &HDWallet{}

// KeysManager Fetch users' private keys
type KeysManager interface {
	// GetEntityPrivateKey Get an entity's private key
	// Note: no contex is provided because the key should be gnerated in memory
	GetEntityPrivateKey(entity *models.Entity) (*ecdsa.PrivateKey, error)

	// GetEntityAddress - Get an entity's public address
	GetEntityAddress(entity *models.Entity) (*common.Address, error)
}

// GetEntityPrivateKeyAsString Helper function to get a private key and encode as a hex string
func GetEntityPrivateKeyAsString(keysMgr KeysManager, entity *models.Entity) (string, error) {
	k, err := keysMgr.GetEntityPrivateKey(entity)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(k.D.Bytes()), nil
}

// GetEntityAddressAsString is just a helper function to get an entity's address
// and convert it to a string at the same time.
func GetEntityAddressAsString(keysMgr KeysManager, entity *models.Entity) (string, error) {
	address, err := keysMgr.GetEntityAddress(entity)
	if err != nil {
		return "", err
	}
	return address.Hex(), nil
}

// NewHDWallet creates a new HDWallet from a seed. The seed must be between 128 and 512 bits
func NewHDWallet(seed []byte) (*HDWallet, error) {
	wallet, err := hdwallet.NewFromSeed(seed)
	if err != nil {
		return nil, err
	}

	return &HDWallet{
		wallet: wallet,
	}, nil
}

// GetEntityPrivateKey implements KeysManager
func (w *HDWallet) GetEntityPrivateKey(entity *models.Entity) (*ecdsa.PrivateKey, error) {
	account, err := getDerivedAccount(w.wallet, entity)
	if err != nil {
		return nil, err
	}

	return w.wallet.PrivateKey(*account)
}

// GetEntityAddress - return an entity's address
func (w *HDWallet) GetEntityAddress(entity *models.Entity) (*common.Address, error) {
	account, err := getDerivedAccount(w.wallet, entity)
	if err != nil {
		return nil, err
	}
	return &account.Address, nil
}

// getDerivedAccount - helper function for mapping Enitity's to ethereum accounts
func getDerivedAccount(wallet *hdwallet.Wallet, entity *models.Entity) (*accounts.Account, error) {
	derivedKeyRelPath, err := entityToKeyRelPath(entity)
	if err != nil {
		return nil, err
	}
	entityPath, err := hdwallet.ParseDerivationPath(derivedKeyRelPath)
	if err != nil {
		return nil, err
	}
	account, err := wallet.Derive(entityPath, false)
	if err != nil {
		return nil, err
	}

	return &account, nil
}

func entityToKeyRelPath(entity *models.Entity) (string, error) {

	id, err := uuid.Parse(entity.Id)
	if err != nil {
		return "", err
	}

	idBytes, err := id.MarshalBinary()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%d/%d/%d/%d",
		binary.BigEndian.Uint32(idBytes[:4]),
		binary.BigEndian.Uint32(idBytes[4:8]),
		binary.BigEndian.Uint32(idBytes[8:12]),
		binary.BigEndian.Uint32(idBytes[12:])), nil

}
