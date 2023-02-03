// Package kms for all the functionality related to kms
package kms

//go:generate mockgen -source kms.go -destination mock/mock.go -package mock

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kms"
	"golang.org/x/crypto/bcrypt"
)

// Config ...
type Config struct {
	KeyAliasName string
}

// Client ...
type Client interface {
	Encrypt(plaintext string) (*string, error)
	Decrypt(ciphertext string) ([]byte, error)
	GetHashedData(data string) string
	GetEncryptedHashedData(data string) (*string, error)
	CompareHashedString(original string, compare string) bool
}

// ClientImpl client Implementation
type ClientImpl struct {
	Client       *kms.KMS
	KeyAliasName string
}

var _ Client = &ClientImpl{}

// NewKMSClient Client for KMS
func NewKMSClient(config Config) (*ClientImpl, error) {
	keyAlias := config.KeyAliasName
	ok := strings.Contains(config.KeyAliasName, "alias/")
	if !ok {
		keyAlias = "alias/" + config.KeyAliasName
	}

	// Initialize a session that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials.
	sess, err := session.NewSession()
	if err != nil {
		return nil, err
	}
	kmsClient := kms.New(sess)
	return &ClientImpl{
		Client:       kmsClient,
		KeyAliasName: keyAlias,
	}, nil
}

// Encrypt encrypts the plaintext
func (k *ClientImpl) Encrypt(plaintext string) (*string, error) {

	// Encrypt the plaintext using the Key Alias
	result, err := k.Client.Encrypt(&kms.EncryptInput{
		KeyId:     aws.String(k.KeyAliasName),
		Plaintext: []byte(plaintext),
	})
	if err != nil {
		return nil, fmt.Errorf("error encrypting plaintext:: %w", err)
	}

	// encode to base64
	encodedResult := k.encodeToBase64String(result.CiphertextBlob)

	return &encodedResult, nil
}

// Decrypt decrypts the ciphertext
func (k *ClientImpl) Decrypt(ciphertext string) ([]byte, error) {

	// decode to base64
	decodeCipherText, err := k.decodeBase64StringToBytes(ciphertext)
	if err != nil {
		return nil, err
	}

	decodedString, err := k.Client.Decrypt(&kms.DecryptInput{
		CiphertextBlob: decodeCipherText,
	})
	if err != nil {
		return nil, fmt.Errorf("unable to decrypt data key: %w", err)
	}

	return decodedString.Plaintext, nil
}

// encodeBase64String function to encode string to base64
func (k *ClientImpl) encodeToBase64String(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}

// decodeBase64String function to decode base64 string to byte array
func (k *ClientImpl) decodeBase64StringToBytes(data string) ([]byte, error) {
	res, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// GetHashedData function to hash given data
func (k *ClientImpl) GetHashedData(data string) string {
	hash := sha256.New()
	// write the plaintext to the hash
	hash.Write([]byte(data))
	// get the resulting hash
	hashEmailBytes := hash.Sum(nil)

	return fmt.Sprintf("%x", hashEmailBytes)
}

// GetEncryptedHashedData function to hash given data using becrypt, used for hashing transaction pin
func (k *ClientImpl) GetEncryptedHashedData(data string) (*string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	hashedPasswordStr := string(hashedPassword)
	return &hashedPasswordStr, nil
}

// CompareHashedString function to compare the plain data with orginal string (Encrypted)
func (k *ClientImpl) CompareHashedString(original string, compare string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(original), []byte(compare))
	return err == nil
}
