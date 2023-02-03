// Package money ...
package money

import (
	"fmt"
	"math"
	"math/big"
	"strconv"

	"github.com/kappapay/backend/lib/golang/src/kappa/common/models"
	"github.com/kappapay/backend/lib/golang/src/kappa/common/types"
)

var (
	// ErrInvalidAmount this error is returned when the amount given is invalid. An empty Currency for example.
	ErrInvalidAmount = fmt.Errorf("invalid Amount")

	// ErrInvalidConversion the exchange rate conversion could not be completed because the given inputs were not valid.
	// For example, if the currency of the amount being converted is neither the "to" nor "from" currency in the exchange rate, this would return in an invalid conversion error.
	ErrInvalidConversion = fmt.Errorf("invalid currency conversion")

	// ErrInvalidRate the rate was not valid. A rate less than or equal to zero is invalid.
	ErrInvalidRate = fmt.Errorf("invalid rate")
)

// SumAmounts should be used to get the sum for n amounts
func SumAmounts(amounts ...*models.Amount) (*models.Amount, error) {
	if len(amounts) < 1 {
		return nil, fmt.Errorf("cannot sum zero values")
	}

	total := big.NewFloat(float64(0))
	baseCurrency := amounts[0].GetCurrency()
	for _, am := range amounts {
		if !CurrenciesEqual(am.GetCurrency(), baseCurrency) {
			return nil, fmt.Errorf("cannot sum amounts in different currencies")
		}
		parsedAmount, err := ParseAmountToFloat(am)
		if err != nil {
			return nil, err
		}
		parsedAmountF := big.NewFloat(float64(*parsedAmount))
		// Add sets total to the (possibly rounded) sum x+y and returns total
		total = total.Add(total, parsedAmountF)
	}

	totalAmount := &models.Amount{
		Value: &types.Decimal{
			Value: total.String(),
		},
		Currency: baseCurrency,
	}

	// format value if possible, minor unit zero padding
	if totalAmount.GetCurrency().GetSubunit() != nil {
		floatValue, _ := total.Float64()
		totalAmount.Value.Value = strconv.FormatFloat(floatValue, 'f', int(totalAmount.Currency.Subunit.Precision), 64)
	}

	return totalAmount, nil
}

// SubAmounts should be used to get the difference of two amounts
func SubAmounts(leftAmount *models.Amount, rightAmount *models.Amount) (*models.Amount, error) {
	if !CurrenciesEqual(leftAmount.GetCurrency(), rightAmount.GetCurrency()) {
		return nil, fmt.Errorf("cannot sub amounts in different currencies")
	}

	diff := big.NewFloat(float64(0))
	parsedLeftAmount, err := ParseAmountToFloat(leftAmount)
	if err != nil {
		return nil, err
	}
	parsedRightAmount, err := ParseAmountToFloat(rightAmount)
	if err != nil {
		return nil, err
	}
	// Sub sets diff to the possibly rounded difference left-right and returns diff.
	diff = diff.Sub(big.NewFloat(float64(*parsedLeftAmount)), big.NewFloat(float64(*parsedRightAmount)))

	diffAmount := &models.Amount{
		Value: &types.Decimal{
			Value: diff.String(),
		},
		Currency: leftAmount.GetCurrency(),
	}

	// format value if possible, minor unit zero padding
	if diffAmount.GetCurrency().GetSubunit() != nil {
		floatValue, _ := diff.Float64()
		diffAmount.Value.Value = strconv.FormatFloat(floatValue, 'f', int(diffAmount.Currency.Subunit.Precision), 64)
	}

	return diffAmount, nil
}

// ParseAmountToFloat should be used to obtain float values for amounts
func ParseAmountToFloat(amount *models.Amount) (*float64, error) {

	val := big.NewFloat(float64(0))
	val, _, err := val.Parse(amount.Value.Value, 10)
	if err != nil {
		return nil, err
	}
	floatAmount, _ := val.Float64()

	return &floatAmount, nil
}

// RoundingMethod is used for rounding amounts
type RoundingMethod int

const (
	// RoundingMethodCeil ceil to the smallest number i such that i >= x
	RoundingMethodCeil RoundingMethod = iota
	// RoundingMethodFloor floor to the largest number i such that i <= x
	RoundingMethodFloor
	// RoundingMethodTruncate discard fractional part of x
	RoundingMethodTruncate
	// RoundingMethodNearest round to the nearest number, rounding half up
	RoundingMethodNearest
	// RoundingMethodHalfToEven round to the nearest number, rounding half to even
	RoundingMethodHalfToEven
)

// RoundAmount will return rounded amount based on rounding method and currency subunit precision
// optional override precision to skip using currency subunit
func RoundAmount(amount *models.Amount, roundingMethod RoundingMethod, overridePrecision *uint) (*models.Amount, error) {

	parsedAmount, err := ParseAmountToFloat(amount)
	if err != nil {

		return nil, err
	}

	var precisionMultiplier float64
	var precision int
	if overridePrecision != nil {
		precisionMultiplier = math.Pow(10, float64(*overridePrecision))
		precision = int(*overridePrecision)
	} else if amount.Currency.Subunit != nil {
		precisionMultiplier = math.Pow(10, float64(amount.Currency.Subunit.Precision))
		precision = int(amount.Currency.Subunit.Precision)
	} else {
		return nil, fmt.Errorf("both override and subunit precision cannot be nil")
	}

	tempAmount := *parsedAmount * precisionMultiplier
	switch roundingMethod {
	case RoundingMethodCeil:
		tempAmount = math.Ceil(tempAmount)
	case RoundingMethodFloor, RoundingMethodTruncate:
		tempAmount = math.Floor(tempAmount)
	case RoundingMethodNearest:
		tempAmount = math.Round(tempAmount)
	case RoundingMethodHalfToEven:
		tempAmount = math.RoundToEven(tempAmount)
	default:
		return nil, fmt.Errorf("rounding not supported: %v", roundingMethod)
	}

	tempAmount = tempAmount / precisionMultiplier

	return &models.Amount{
		Value: &types.Decimal{
			Value: strconv.FormatFloat(tempAmount, 'f', precision, 64),
		},
		Currency: amount.Currency,
	}, nil
}

// validateCurrency should be used to compare currencies of two amounts
func validateCurrency(leftAmount *models.Amount, rightAmount *models.Amount) error {

	if leftAmount.Currency == nil || leftAmount.Currency.Id == "" {
		return fmt.Errorf("missing currency: leftAmount")
	}

	if rightAmount.Currency == nil || rightAmount.Currency.Id == "" {
		return fmt.Errorf("missing currency: rightAmount")
	}

	if leftAmount.Currency.Id != rightAmount.Currency.Id {
		return fmt.Errorf("cannot compare amounts with different currencies, got left currency ID '%s' and right currency ID '%s'", leftAmount.Currency.Id, rightAmount.Currency.Id)
	}

	return nil
}

// EqualTo performs an equality comparison on two provided amounts
func EqualTo(leftAmount *models.Amount, rightAmount *models.Amount) (*bool, error) {

	err := validateCurrency(leftAmount, rightAmount)
	if err != nil {
		return nil, err
	}

	leftAmountF, err := ParseAmountToFloat(leftAmount)
	if err != nil {
		return nil, err
	}
	leftAmountFloat := big.NewFloat(float64(*leftAmountF))

	rightAmountF, err := ParseAmountToFloat(rightAmount)
	if err != nil {
		return nil, err
	}
	rightAmountFloat := big.NewFloat(float64(*rightAmountF))

	result := compare(leftAmountFloat, rightAmountFloat)
	equalTo := result == 0
	return &equalTo, nil
}

// LessThan performs a less than comparison on provided amounts
func LessThan(leftAmount *models.Amount, rightAmount *models.Amount) (*bool, error) {

	err := validateCurrency(leftAmount, rightAmount)
	if err != nil {
		return nil, err
	}

	leftAmountF, err := ParseAmountToFloat(leftAmount)
	if err != nil {
		return nil, err
	}
	leftAmountFloat := big.NewFloat(float64(*leftAmountF))

	rightAmountF, err := ParseAmountToFloat(rightAmount)
	if err != nil {
		return nil, err
	}
	rightAmountFloat := big.NewFloat(float64(*rightAmountF))

	result := compare(leftAmountFloat, rightAmountFloat)
	lessThan := result < 0
	return &lessThan, nil
}

// GreaterThan performs a greater than comparison on provided amounts
func GreaterThan(leftAmount *models.Amount, rightAmount *models.Amount) (*bool, error) {

	err := validateCurrency(leftAmount, rightAmount)
	if err != nil {
		return nil, err
	}

	leftAmountF, err := ParseAmountToFloat(leftAmount)
	if err != nil {
		return nil, err
	}
	leftAmountFloat := big.NewFloat(float64(*leftAmountF))

	rightAmountF, err := ParseAmountToFloat(rightAmount)
	if err != nil {
		return nil, err
	}
	rightAmountFloat := big.NewFloat(float64(*rightAmountF))

	result := compare(leftAmountFloat, rightAmountFloat)
	greaterThan := result > 0
	return &greaterThan, nil
}

// compare compares x and y and returns:
//
//   -1 if x <  y
//    0 if x == y (incl. -0 == 0, -Inf == -Inf, and +Inf == +Inf)
//   +1 if x >  y

// compare ...
func compare(x *big.Float, y *big.Float) int {
	return x.Cmp(y)
}

// MulAmount amount multiple an amount by a float
func MulAmount(amount *models.Amount, x float64) (*models.Amount, error) {
	parsedAmount, err := ParseAmountToFloat(amount)
	if err != nil {
		return nil, fmt.Errorf("could not parse value as float: value=%s %w", amount.Value.Value, ErrInvalidAmount)
	}
	result := big.NewFloat(float64(8)).Mul(big.NewFloat(*parsedAmount), big.NewFloat(x))

	return &models.Amount{
		Value: &types.Decimal{
			Value: result.String(),
		},
		Currency: amount.Currency,
	}, nil

}

// ConvertCurrency convert one currency to another given an exchange rate
func ConvertCurrency(amount *models.Amount, rate *models.ExchangeRate) (*models.Amount, error) {
	if !CurrenciesEqual(amount.GetCurrency(), rate.GetCurrencyFrom()) &&
		!CurrenciesEqual(amount.GetCurrency(), rate.GetCurrencyTo()) {
		return nil, ErrInvalidConversion
	}

	if rate.GetValue() <= 0 {
		return nil, fmt.Errorf("rate must be greater than zero: rate=%f: %w", rate.GetValue(), ErrInvalidRate)
	}

	// First assume we are going from the "from currency" to the "to currency"
	factor := rate.Value
	resultingCurrency := rate.GetCurrencyTo()

	// If the amount currency is actually the "to currency" we will invert the rate
	if CurrenciesEqual(amount.GetCurrency(), rate.GetCurrencyTo()) {
		factor = 1 / rate.Value
		resultingCurrency = rate.GetCurrencyFrom()
	}

	// note converted amount will still be in the original currency, we update it below
	convertedAmount, err := MulAmount(amount, factor)
	if err != nil {
		// The source amount must contain an invalid string, thus it is an invalid quote error.
		return nil, fmt.Errorf("couldn't calculate the converted amount: %w", err)
	}

	convertedAmount.Currency = resultingCurrency
	return convertedAmount, nil
}

// CurrenciesEqual compares two currencies to see if they are equal.
// This method will first compare the country ID, followed by ISO country code and finally name.
func CurrenciesEqual(a, b *models.Currency) bool {
	if a.GetId() != "" && b.GetId() != "" {
		return a.GetId() == b.GetId()
	}

	if a.GetIsoCurrencyCode() != "" && b.GetIsoCurrencyCode() != "" {
		return a.GetIsoCurrencyCode() == b.GetIsoCurrencyCode()
	}

	if a.GetName() != "" && b.GetName() != "" {
		return a.GetName() == b.GetName()
	}

	// no overlapping fields. If all fields are empty then we can say the currencies match.
	return a.GetId()+a.GetIsoCurrencyCode()+a.GetName()+b.GetId()+b.GetIsoCurrencyCode()+b.GetName() == ""
}

// TransactionFee is the satisfying interface for any proto model with a GetAmount func
type TransactionFee interface {
	GetAmount() *models.Amount
}

// SumFeeAmounts ...
func SumFeeAmounts[T TransactionFee](fees []T, fallbackCurrency *models.Currency) (*models.Amount, error) {
	total := &models.Amount{
		Currency: fallbackCurrency,
		Value:    &types.Decimal{Value: "0"},
	}

	if len(fees) > 0 {
		total.Currency = fees[0].GetAmount().GetCurrency() // Set Base Currency from [0] with the expectation that all fees are in the same currency
	}
	for _, fee := range fees {
		sum, err := SumAmounts(total, fee.GetAmount())
		if err != nil {
			return &models.Amount{}, err
		}
		total = sum
	}
	return total, nil
}
