// Package constants defines quasi-status variables and values for operational reasons
package constants

const (
	// Country codes

	// ISOCountryCodeCMR CMR Country Code
	ISOCountryCodeCMR = "CMR"
	// ISOCountryCodeUSA USA Country Code
	ISOCountryCodeUSA = "USA"

	// Currency Codes

	// ISOCurrencyCodeXAF XAF Currency Code
	ISOCurrencyCodeXAF = "XAF"
	// ISOCurrencyCodeUSD USD Currency Code
	ISOCurrencyCodeUSD = "USD"

	// Kappa Entity Official Names

	// BusinessEntityOfficialNameKappaPayCMR name
	BusinessEntityOfficialNameKappaPayCMR = "Kappa Pay Cameroon PLC"
	// BusinessEntityOfficialNameKappaPayUSA name
	BusinessEntityOfficialNameKappaPayUSA = "Kappa Pay Inc"

	// CMR Partner Entity Official Names

	// BusinessEntityOfficialNameEcoBankCMR ...
	BusinessEntityOfficialNameEcoBankCMR = "Ecobank Cameroon PLC"
	// BusinessEntityOfficialNameMTNCMR ...
	BusinessEntityOfficialNameMTNCMR = "MTN Cameroon"
	// BusinessEntityOfficialNameOrangeCMR ...
	BusinessEntityOfficialNameOrangeCMR = "Orange Cameroon"

	// USA Partner Entity Official Names

	// BusinessEntityOfficialNameSila ...
	BusinessEntityOfficialNameSila = "Sila Inc"
	// BusinessEntityOfficialNameEvolve ...
	BusinessEntityOfficialNameEvolve = "Evolve Bank & Trust"
	// BusinessEntityOfficialNameMercury ...
	BusinessEntityOfficialNameMercury = "Mercury Technologies Inc"

	// Ledger Constants

	// Assets ...
	Assets = "Assets"
	// Liabilities ...
	Liabilities = "Liabilities"
	// Wallets ...
	Wallets = "Wallets"
	// Customers ...
	Customers = "Customers"
	// Operations ...
	Operations = "Operations"
	// Equity ...
	Equity = "Equity"
	// Revenue ...
	Revenue = "Revenue"
	// Expenses ...
	Expenses = "Expenses"
	// Trading ...
	Trading = "Trading"
	// NostroAccounts ...
	NostroAccounts = "Nostro Accounts"
	// VostroAccounts ...
	VostroAccounts = "Vostro Accounts"
	// Current ...
	Current = "Current"
	// NonCurrent ...
	NonCurrent = "Non-current"
	// BankAccounts ...
	BankAccounts = "Bank Accounts"
	// BaaSProviderAccounts ...
	BaaSProviderAccounts = "BaaS Provider Accounts"
	// PreFundedFBOAccount ...
	PreFundedFBOAccount = "Pre-funded FBO Account"
	// GeneralCorporateAccount ...
	GeneralCorporateAccount = "General Corporate Account"
	// GeneralFBOAccount ...
	GeneralFBOAccount = "General FBO Account"
	// OperatingAccount ...
	OperatingAccount = "Operating Account"
	// CurrencyTradingAccounts ...
	CurrencyTradingAccounts = "Currency Trading Accounts"
	// CommonStock ...
	CommonStock = "Common Stock"
	// RetainedEarnings ...
	RetainedEarnings = "Retained Earnings"
	// OperatingWallet ...
	OperatingWallet = "Operating Wallet"
	// CustomerAndRelatedAccounts ...
	CustomerAndRelatedAccounts = "Customer and Related Accounts"
	// AccountsPayable ...
	AccountsPayable = "Accounts Payable"
	// WalletsOutgoingTransactions ...
	WalletsOutgoingTransactions = "Wallets - Outgoing Transactions"
	// RelatedPartyAccounts ...
	RelatedPartyAccounts = "Related Party Accounts"
	// ServicesInProgress ...
	ServicesInProgress = "Services in Progress"
	// TransactionServicesInProgress ...
	TransactionServicesInProgress = "Transaction Services in Progress"
	// TransactionFeeRevenue ...
	TransactionFeeRevenue = "Transaction Fee Revenue"
	// TransactionFeeExpenses ...
	TransactionFeeExpenses = "Transaction Fee Expenses"
	// InterestExpenses ...
	InterestExpenses = "Interest Expenses"
	// InterestBuoyRevolverLoan ...
	InterestBuoyRevolverLoan = "Interest Buoy Revolver Loan"
	// FXGains ...
	FXGains = "FX Gains"
	// FXLosses ...
	FXLosses = "FX Losses"
	// CostOfGoodsSold ...
	CostOfGoodsSold = "Cost of Goods Sold"

	// PaymentGatewaySilaStandardACH ...
	PaymentGatewaySilaStandardACH = "SILA_STANDARD_ACH"
	// PaymentGatewaySilaSameDayACH ...
	PaymentGatewaySilaSameDayACH = "SILA_SAME_DAY_ACH"
	// PaymentGatewaySilaWalletTransfer ...
	PaymentGatewaySilaWalletTransfer = "SILA_WALLET_TRANSFER"
	// PaymentGatewayEcobankPaymentsAPI ...
	PaymentGatewayEcobankPaymentsAPI = "ECOBANK_PAYMENTS_API"
	// PaymentGatewayManualBankTransferCMR ...
	PaymentGatewayManualBankTransferCMR = "MANUAL_BANK_TRANSFER_CMR"
	// PaymentGatewayKPCWalletTransfer ...
	PaymentGatewayKPCWalletTransfer = "KPC_WALLET_TRANSFER"

	// Sila Constants

	// SilaTransactionTypeTransfer ...
	SilaTransactionTypeTransfer = "transfer"
	// SilaTransactionTypeRedeem is the status associated with withdrawal transactions
	SilaTransactionTypeRedeem = "redeem"
	// SilaTransactionTypeIssue is the status associated with deposit transactions
	SilaTransactionTypeIssue = "issue"
	// SilaTransactionStatusSuccess ...
	SilaTransactionStatusSuccess = "success"
	// SilaTransactionStatusFailed ...
	SilaTransactionStatusFailed = "failed"
	// SilaTransactionStatusReview ...
	SilaTransactionStatusReview = "review"
	// SilaTransactionStatusQueued ...
	SilaTransactionStatusQueued = "queued"
)

// ISOCountryCodes ...
type ISOCountryCodes struct {
	CMR string
	USA string
}

// ISOCurrencyCodes ...
type ISOCurrencyCodes struct {
	XAF string
	USD string
}

// BankPartners ...
type BankPartners struct {
	EcoBankCMR string
	Mercury    string
	Evolve     string
}

// BaaSProviders ...
type BaaSProviders struct {
	Sila string
}

// MNOPartners ...
type MNOPartners struct {
	MTNCMR    string
	OrangeCMR string
}

// KappaEntities ...
type KappaEntities struct {
	KappaPayCMR string
	KappaPayUSA string
}

// PaymentGateways ...
type PaymentGateways struct {
	SilaStandardACH       string
	SilaSameDayACH        string
	EcobankPaymentsAPI    string
	ManualBankTransferCMR string
}

// GetISOCountryCodes ...
func GetISOCountryCodes() ISOCountryCodes {
	return ISOCountryCodes{
		CMR: ISOCountryCodeCMR,
		USA: ISOCountryCodeUSA,
	}
}

// GetISOCurrencyCodes ...
func GetISOCurrencyCodes() ISOCurrencyCodes {
	return ISOCurrencyCodes{
		XAF: ISOCurrencyCodeXAF,
		USD: ISOCurrencyCodeUSD,
	}
}

// GetBankPartners ...
func GetBankPartners() BankPartners {
	return BankPartners{
		EcoBankCMR: BusinessEntityOfficialNameEcoBankCMR,
		Mercury:    BusinessEntityOfficialNameMercury,
		Evolve:     BusinessEntityOfficialNameEvolve,
	}
}

// GetBaaSProviders ...
func GetBaaSProviders() BaaSProviders {
	return BaaSProviders{
		Sila: BusinessEntityOfficialNameSila,
	}
}

// GetMNOPartners ...
func GetMNOPartners() MNOPartners {
	return MNOPartners{
		MTNCMR:    BusinessEntityOfficialNameMTNCMR,
		OrangeCMR: BusinessEntityOfficialNameOrangeCMR,
	}
}

// GetKappaEntities ...
func GetKappaEntities() KappaEntities {
	return KappaEntities{
		KappaPayCMR: BusinessEntityOfficialNameKappaPayCMR,
		KappaPayUSA: BusinessEntityOfficialNameKappaPayUSA,
	}
}

// GetBusinessEntityOfficialNames ...
func GetBusinessEntityOfficialNames() []string {
	return []string{
		BusinessEntityOfficialNameMTNCMR,
		BusinessEntityOfficialNameOrangeCMR,
		BusinessEntityOfficialNameEcoBankCMR,
		BusinessEntityOfficialNameKappaPayCMR,
		BusinessEntityOfficialNameKappaPayUSA,
		BusinessEntityOfficialNameSila,
		BusinessEntityOfficialNameEvolve,
		BusinessEntityOfficialNameMercury,
	}
}

// GetPaymentGateways ...
func GetPaymentGateways() PaymentGateways {
	return PaymentGateways{
		PaymentGatewaySilaStandardACH,
		PaymentGatewaySilaSameDayACH,
		PaymentGatewayEcobankPaymentsAPI,
		PaymentGatewayManualBankTransferCMR,
	}
}
