// Package shared ...
package shared

import (
	"database/sql/driver"
	"fmt"
)

// SystemName represents the name of a downstream system on which this system relies
type SystemName string

const (
	//SystemNone represents an empty string
	SystemNone SystemName = ""
	// SystemAddress ...
	SystemAddress SystemName = "ADDRESS"
	// SystemAuthentication ...
	SystemAuthentication SystemName = "AUTHENTICATION"
	// SystemBankAccounts ...
	SystemBankAccounts SystemName = "BANK_ACCOUNTS"
	// SystemCompliance ...
	SystemCompliance SystemName = "COMPLIANCE"
	// SystemConfiguration ...
	SystemConfiguration SystemName = "CONFIGURATION"
	// SystemEmails ...
	SystemEmails SystemName = "EMAILS"
	// SystemEntities ...
	SystemEntities SystemName = "ENTITIES"
	// SystemExecution ...
	SystemExecution SystemName = "EXECUTION"
	// SystemFunding ...
	SystemFunding SystemName = "FUNDING"
	// SystemFundingMonitor ...
	SystemFundingMonitor SystemName = "FUNDING_MONITOR"
	// SystemLedger ...
	SystemLedger SystemName = "LEDGER"
	// SystemLimits ...
	SystemLimits SystemName = "LIMITS"
	// SystemLinkedAccounts ...
	SystemLinkedAccounts SystemName = "LINKED_ACCOUNTS"
	// SystemMobileMoneyAccounts ...
	SystemMobileMoneyAccounts SystemName = "MOBILE_MONEY_ACCOUNTS"
	// SystemNotifications ...
	SystemNotifications SystemName = "NOTIFICATIONS"
	// SystemPhones ...
	SystemPhones SystemName = "PHONES"
	// SystemQuotes ...
	SystemQuotes SystemName = "QUOTES"
	// SystemRecipients ...
	SystemRecipients SystemName = "RECIPIENTS"
	// SystemSettlements ...
	SystemSettlements SystemName = "SETTLEMENTS"
	// SystemSila ...
	SystemSila SystemName = "SILA"
	// SystemTransfers ...
	SystemTransfers SystemName = "TRANSFERS"
	// SystemsUser ...
	SystemsUser SystemName = "USERS"
	// SystemWallets ...
	SystemWallets SystemName = "WALLETS"
	// SystemRelay ...
	SystemRelay SystemName = "RELAY"
)

var systemSet = map[SystemName]struct{}{
	SystemNone:                {},
	SystemAddress:             {},
	SystemAuthentication:      {},
	SystemBankAccounts:        {},
	SystemCompliance:          {},
	SystemConfiguration:       {},
	SystemEmails:              {},
	SystemEntities:            {},
	SystemExecution:           {},
	SystemFunding:             {},
	SystemFundingMonitor:      {},
	SystemLedger:              {},
	SystemLimits:              {},
	SystemLinkedAccounts:      {},
	SystemMobileMoneyAccounts: {},
	SystemNotifications:       {},
	SystemPhones:              {},
	SystemQuotes:              {},
	SystemRecipients:          {},
	SystemSettlements:         {},
	SystemSila:                {},
	SystemTransfers:           {},
	SystemsUser:               {},
	SystemWallets:             {},
	SystemRelay:               {},
}

// Scan implements the Scanner interface, and only accepts a string type for src.
// If the string value doesn't exist, Scan() will return an error.
func (s *SystemName) Scan(src interface{}) error {
	strVal, isString := src.(string)
	if !isString {
		return fmt.Errorf("value scanned by SystemName is not a string %+v", src)
	}

	value := SystemName(strVal)
	if err := value.Validate(); err != nil {
		return err
	}

	*s = value

	return nil
}

// String implements the Stringer interface, converting the const type to its equivalent string value
// To confirm non-empty string values, users should call Validate() to confirm the mapping exists before calling this method.
func (s SystemName) String() string {
	return string(s)
}

// Validate ensures the SystemName const type is a previously defined value
func (s SystemName) Validate() error {
	if _, ok := systemSet[s]; !ok {
		return fmt.Errorf("the SystemName value %s is not valid", s)
	}

	return nil
}

// Value implements the Valuer interface, converting the const type to the application driver.Value value
func (s SystemName) Value() (driver.Value, error) {
	if err := s.Validate(); err != nil {
		return nil, err
	}

	return driver.Value(s.String()), nil
}
