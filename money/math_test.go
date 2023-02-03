package money

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/kappapay/backend/lib/golang/src/kappa/common/models"
	"github.com/kappapay/backend/lib/golang/src/kappa/common/types"
)

func TestSumAmounts(t *testing.T) {
	tests := map[string]struct {
		amounts   []string
		expected  string
		precision *uint
	}{
		"n_positive_amounts": {
			amounts:  []string{"100", "100"},
			expected: "200",
		},
		"n_negative_amounts": {
			amounts:  []string{"-100", "-100"},
			expected: "-200",
		},
		"n_mixed_sign_amounts": {
			amounts:  []string{"100", "-100"},
			expected: "0",
		},
		"test_with_precision_1": {
			amounts:   []string{"12", "15", "11"},
			expected:  "38.0",
			precision: uintPtr(1),
		},
		"test_with_precision_2": {
			amounts:   []string{"12", "15", "11.1"},
			expected:  "38.10",
			precision: uintPtr(2),
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			// set currency if we have precision
			var currency *models.Currency
			if test.precision != nil {
				currency = &models.Currency{
					Subunit: &models.CurrencySubunit{
						Precision: int32(*test.precision),
					},
				}
			}

			amounts := []*models.Amount{}
			for _, amount := range test.amounts {
				amounts = append(amounts, &models.Amount{
					Value:    &types.Decimal{Value: amount},
					Currency: currency,
				})
			}
			actualAmount, err := SumAmounts(amounts...)
			assert.NoError(t, err)

			assert.Equal(t, test.expected, actualAmount.Value.Value)
		})
	}
}

func TestSubAmounts(t *testing.T) {

	tests := map[string]struct {
		leftAmount  string
		rightAmount string
		expected    string
		precision   *uint // to test minor unit zero padding
	}{
		"two_positive_amounts": {
			leftAmount:  "100",
			rightAmount: "100",
			expected:    "0",
		},
		"two_negative_amounts": {
			leftAmount:  "-100",
			rightAmount: "-100",
			expected:    "0",
		},
		"two_mixed_sign_amounts_1": {
			leftAmount:  "100",
			rightAmount: "-100",
			expected:    "200",
		},
		"two_mixed_sign_amounts_2": {
			leftAmount:  "-100",
			rightAmount: "100",
			expected:    "-200",
		},
		"test_precision_amounts_1": {
			leftAmount:  "12.34",
			rightAmount: "1.04",
			expected:    "11.30",
			precision:   uintPtr(2),
		},
		"test_precision_amounts_2": {
			leftAmount:  "12",
			rightAmount: "1",
			expected:    "11.00",
			precision:   uintPtr(2),
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			// set currency if we have precision
			var currency *models.Currency
			if test.precision != nil {
				currency = &models.Currency{
					Subunit: &models.CurrencySubunit{
						Precision: int32(*test.precision),
					},
				}
			}

			leftAmount := &models.Amount{
				Value: &types.Decimal{
					Value: test.leftAmount,
				},
				Currency: currency,
			}
			rightAmount := &models.Amount{
				Value: &types.Decimal{
					Value: test.rightAmount,
				},
				Currency: currency,
			}
			actualAmount, err := SubAmounts(leftAmount, rightAmount)
			assert.NoError(t, err)

			assert.Equal(t, test.expected, actualAmount.Value.Value)
		})
	}
}

func TestRoundAmounts(t *testing.T) {
	testCases := map[string]struct {
		inputAmount              string
		expectedAmount           string
		currencySubunitPrecision uint
		overridePrecision        *uint
		roundType                RoundingMethod
	}{
		"precision 0 ceil": {
			inputAmount:              "2.3821",
			expectedAmount:           "3",
			currencySubunitPrecision: 100,
			overridePrecision:        uintPtr(0),
			roundType:                RoundingMethodCeil,
		},
		"precision 0 floor": {
			inputAmount:              "2.9838348329948821",
			expectedAmount:           "2",
			currencySubunitPrecision: 0,
			roundType:                RoundingMethodFloor,
		},
		"precision 2 ceil": {
			inputAmount:              "59.99999",
			expectedAmount:           "60.00",
			currencySubunitPrecision: 0,
			overridePrecision:        uintPtr(2),
			roundType:                RoundingMethodCeil,
		},
		"precision 2 floor": {
			inputAmount:              "59.99999",
			expectedAmount:           "59.99",
			currencySubunitPrecision: 2,
			roundType:                RoundingMethodFloor,
		},
		"precision 3 ceil": {
			inputAmount:              "123.1231",
			expectedAmount:           "123.124",
			currencySubunitPrecision: 3,
			roundType:                RoundingMethodCeil,
		},
		"precision 3 floor": {
			inputAmount:              "123.1239",
			expectedAmount:           "123.123",
			currencySubunitPrecision: 3,
			roundType:                RoundingMethodFloor,
		},
		"precision 2 truncate": {
			inputAmount:              "683.8999",
			expectedAmount:           "683.89",
			currencySubunitPrecision: 2,
			roundType:                RoundingMethodTruncate,
		},
		"precision 2 nearest": {
			inputAmount:              "888.133",
			expectedAmount:           "888.13",
			currencySubunitPrecision: 4,
			overridePrecision:        uintPtr(2),
			roundType:                RoundingMethodNearest,
		},
		"will break on float32": {
			inputAmount:              "1000000.011",
			expectedAmount:           "1000000.01",
			currencySubunitPrecision: 2,
			roundType:                RoundingMethodNearest,
		},
		"precision 2 nearest.": {
			inputAmount:              "888.135",
			expectedAmount:           "888.14",
			currencySubunitPrecision: 2,
			roundType:                RoundingMethodNearest,
		},
		"precision 2 nearest..": {
			inputAmount:              "888.136",
			expectedAmount:           "888.14",
			currencySubunitPrecision: 2,
			roundType:                RoundingMethodNearest,
		},
		"precision 2 half to even": {
			inputAmount:              "12.346",
			expectedAmount:           "12.35",
			currencySubunitPrecision: 2,
			roundType:                RoundingMethodHalfToEven,
		},
		"precision 2 half to even.": {
			inputAmount:              "12.3423",
			expectedAmount:           "12.34",
			currencySubunitPrecision: 2,
			roundType:                RoundingMethodHalfToEven,
		},
		"precision 3 half to even": {
			inputAmount:              "12.1235",
			expectedAmount:           "12.124",
			currencySubunitPrecision: 1,
			overridePrecision:        uintPtr(3),
			roundType:                RoundingMethodHalfToEven,
		},
		"precision 3 half to even.": {
			inputAmount:              "12.1245",
			expectedAmount:           "12.124",
			currencySubunitPrecision: 3,
			roundType:                RoundingMethodHalfToEven,
		},
		"precision 2 half to even..": {
			inputAmount:              "440.645",
			expectedAmount:           "440.64",
			currencySubunitPrecision: 2,
			roundType:                RoundingMethodHalfToEven,
		},
		"precision 2 half to even...": {
			inputAmount:              "440.6451",
			expectedAmount:           "440.65",
			currencySubunitPrecision: 2,
			roundType:                RoundingMethodHalfToEven,
		},
		"test zero padding, precision 1": {
			inputAmount:              "12",
			expectedAmount:           "12.0",
			currencySubunitPrecision: 1,
			roundType:                RoundingMethodHalfToEven,
		},
		"test zero padding, precision 2": {
			inputAmount:              "12.3",
			expectedAmount:           "12.30",
			currencySubunitPrecision: 2,
			roundType:                RoundingMethodFloor,
		},
		"test zero padding, precision 5": {
			inputAmount:              "12.3",
			expectedAmount:           "12.30000",
			currencySubunitPrecision: 5,
			roundType:                RoundingMethodFloor,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			in := models.Amount{
				Value: &types.Decimal{
					Value: tc.inputAmount,
				},
				Currency: &models.Currency{
					Subunit: &models.CurrencySubunit{
						Precision: int32(tc.currencySubunitPrecision),
					},
				},
			}
			expected := models.Amount{
				Value: &types.Decimal{
					Value: tc.expectedAmount,
				},
			}

			output, err := RoundAmount(&in, tc.roundType, tc.overridePrecision)
			assert.NoError(t, err)
			assert.Equal(t, expected.Value.Value, output.Value.Value)
		})
	}
}

func TestEqualTo(t *testing.T) {
	positive100NilCurrency := &models.Amount{
		Value: &types.Decimal{
			Value: "100",
		},
	}
	positive100InvalidCurrency := &models.Amount{
		Value: &types.Decimal{
			Value: "100",
		},
		Currency: &models.Currency{
			Id: "",
		},
	}
	positive100ValidCurrency := &models.Amount{
		Value: &types.Decimal{
			Value: "100",
		},
		Currency: &models.Currency{
			Id: "1",
		},
	}
	positive100OtherCurrency := &models.Amount{
		Value: &types.Decimal{
			Value: "100",
		},
		Currency: &models.Currency{
			Id: "2",
		},
	}
	positive200ValidCurrency := &models.Amount{
		Value: &types.Decimal{
			Value: "200",
		},
		Currency: &models.Currency{
			Id: "1",
		},
	}
	tests := map[string]struct {
		leftAmount  *models.Amount
		rightAmount *models.Amount
		expected    bool
		err         error
	}{
		"nil_currency": {
			leftAmount:  positive100NilCurrency,
			rightAmount: positive100ValidCurrency,
			expected:    false,
			err:         fmt.Errorf("missing currency: leftAmount"),
		},
		"invalid_currency": {
			leftAmount:  positive100InvalidCurrency,
			rightAmount: positive100ValidCurrency,
			expected:    false,
			err:         fmt.Errorf("missing currency: leftAmount"),
		},
		"other_currency": {
			leftAmount:  positive100ValidCurrency,
			rightAmount: positive100OtherCurrency,
			expected:    false,
			err:         fmt.Errorf("cannot compare amounts with different currencies, got left currency ID '1' and right currency ID '2'"),
		},
		"different_amount": {
			leftAmount:  positive100ValidCurrency,
			rightAmount: positive200ValidCurrency,
			expected:    false,
			err:         nil,
		},
		"equal": {
			leftAmount:  positive100ValidCurrency,
			rightAmount: positive100ValidCurrency,
			expected:    true,
			err:         nil,
		},
		"less_than": {
			leftAmount:  positive100ValidCurrency,
			rightAmount: positive200ValidCurrency,
			expected:    false,
			err:         nil,
		},
		"greater_than": {
			leftAmount:  positive100ValidCurrency,
			rightAmount: positive200ValidCurrency,
			expected:    false,
			err:         nil,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			cmp, err := EqualTo(test.leftAmount, test.rightAmount)
			if test.err != nil {
				assert.NotNil(t, err)
				assert.Equal(t, test.err.Error(), err.Error())
			} else {
				assert.Equal(t, test.expected, *cmp)
			}
		})
	}
}

func TestLessThan(t *testing.T) {
	positive100NilCurrency := &models.Amount{
		Value: &types.Decimal{
			Value: "100",
		},
	}
	positive100InvalidCurrency := &models.Amount{
		Value: &types.Decimal{
			Value: "100",
		},
		Currency: &models.Currency{
			Id: "",
		},
	}
	positive100ValidCurrency := &models.Amount{
		Value: &types.Decimal{
			Value: "100",
		},
		Currency: &models.Currency{
			Id: "1",
		},
	}
	positive100OtherCurrency := &models.Amount{
		Value: &types.Decimal{
			Value: "100",
		},
		Currency: &models.Currency{
			Id: "2",
		},
	}
	positive200ValidCurrency := &models.Amount{
		Value: &types.Decimal{
			Value: "200",
		},
		Currency: &models.Currency{
			Id: "1",
		},
	}
	tests := map[string]struct {
		leftAmount  *models.Amount
		rightAmount *models.Amount
		expected    bool
		err         error
	}{
		"nil_currency": {
			leftAmount:  positive100NilCurrency,
			rightAmount: positive100ValidCurrency,
			expected:    false,
			err:         fmt.Errorf("missing currency: leftAmount"),
		},
		"invalid_currency": {
			leftAmount:  positive100InvalidCurrency,
			rightAmount: positive100ValidCurrency,
			expected:    false,
			err:         fmt.Errorf("missing currency: leftAmount"),
		},
		"other_currency": {
			leftAmount:  positive100ValidCurrency,
			rightAmount: positive100OtherCurrency,
			expected:    false,
			err:         fmt.Errorf("cannot compare amounts with different currencies, got left currency ID '1' and right currency ID '2'"),
		},
		"different_amount": {
			leftAmount:  positive100ValidCurrency,
			rightAmount: positive200ValidCurrency,
			expected:    true,
			err:         nil,
		},
		"equal": {
			leftAmount:  positive100ValidCurrency,
			rightAmount: positive100ValidCurrency,
			expected:    false,
			err:         nil,
		},
		"less_than": {
			leftAmount:  positive100ValidCurrency,
			rightAmount: positive200ValidCurrency,
			expected:    true,
			err:         nil,
		},
		"greater_than": {
			leftAmount:  positive200ValidCurrency,
			rightAmount: positive100ValidCurrency,
			expected:    false,
			err:         nil,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			cmp, err := LessThan(test.leftAmount, test.rightAmount)
			if test.err != nil {
				assert.NotNil(t, err)
				assert.Equal(t, test.err.Error(), err.Error())
			} else {
				assert.Equal(t, test.expected, *cmp)
			}
		})
	}
}

func TestGreaterThan(t *testing.T) {
	positive100NilCurrency := &models.Amount{
		Value: &types.Decimal{
			Value: "100",
		},
	}
	positive100InvalidCurrency := &models.Amount{
		Value: &types.Decimal{
			Value: "100",
		},
		Currency: &models.Currency{
			Id: "",
		},
	}
	positive100ValidCurrency := &models.Amount{
		Value: &types.Decimal{
			Value: "100",
		},
		Currency: &models.Currency{
			Id: "1",
		},
	}
	positive100OtherCurrency := &models.Amount{
		Value: &types.Decimal{
			Value: "100",
		},
		Currency: &models.Currency{
			Id: "2",
		},
	}
	positive200ValidCurrency := &models.Amount{
		Value: &types.Decimal{
			Value: "200",
		},
		Currency: &models.Currency{
			Id: "1",
		},
	}
	tests := map[string]struct {
		leftAmount  *models.Amount
		rightAmount *models.Amount
		expected    bool
		err         error
	}{
		"nil_currency": {
			leftAmount:  positive100NilCurrency,
			rightAmount: positive100ValidCurrency,
			expected:    false,
			err:         fmt.Errorf("missing currency: leftAmount"),
		},
		"invalid_currency": {
			leftAmount:  positive100InvalidCurrency,
			rightAmount: positive100ValidCurrency,
			expected:    false,
			err:         fmt.Errorf("missing currency: leftAmount"),
		},
		"other_currency": {
			leftAmount:  positive100ValidCurrency,
			rightAmount: positive100OtherCurrency,
			expected:    false,
			err:         fmt.Errorf("cannot compare amounts with different currencies, got left currency ID '1' and right currency ID '2'"),
		},
		"different_amount": {
			leftAmount:  positive100ValidCurrency,
			rightAmount: positive200ValidCurrency,
			expected:    false,
			err:         nil,
		},
		"equal": {
			leftAmount:  positive100ValidCurrency,
			rightAmount: positive100ValidCurrency,
			expected:    false,
			err:         nil,
		},
		"less_than": {
			leftAmount:  positive100ValidCurrency,
			rightAmount: positive200ValidCurrency,
			expected:    false,
			err:         nil,
		},
		"greater_than": {
			leftAmount:  positive200ValidCurrency,
			rightAmount: positive100ValidCurrency,
			expected:    true,
			err:         nil,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			cmp, err := GreaterThan(test.leftAmount, test.rightAmount)
			if test.err != nil {
				assert.NotNil(t, err)
				assert.Equal(t, test.err.Error(), err.Error())
			} else {
				assert.Equal(t, test.expected, *cmp)
			}
		})
	}
}

func TestConvertCurrency(t *testing.T) {

	currencyA := models.Currency{Id: uuid.NewString()}
	currencyB := models.Currency{Id: uuid.NewString()}
	currencyC := models.Currency{Id: uuid.NewString()}

	type args struct {
		amount *models.Amount
		rate   *models.ExchangeRate
	}
	tests := map[string]struct {
		args    args
		want    *models.Amount
		wantErr bool
	}{
		"error_nil_inputs": {
			args:    args{amount: nil, rate: nil},
			want:    nil,
			wantErr: true,
		},
		"valid_2X_rate": {
			args: args{
				amount: &models.Amount{Value: &types.Decimal{Value: "7"}, Currency: &currencyA},
				rate:   &models.ExchangeRate{CurrencyFrom: &currencyA, CurrencyTo: &currencyB, Value: 2},
			},
			want:    &models.Amount{Value: &types.Decimal{Value: "14"}, Currency: &currencyB},
			wantErr: false,
		},
		"valid_2X_rate_reverse": {
			args: args{
				amount: &models.Amount{Value: &types.Decimal{Value: "26"}, Currency: &currencyB},
				rate:   &models.ExchangeRate{CurrencyFrom: &currencyA, CurrencyTo: &currencyB, Value: 2},
			},
			want:    &models.Amount{Value: &types.Decimal{Value: "13"}, Currency: &currencyA},
			wantErr: false,
		},
		"valid_zero_rate": {
			args: args{
				amount: &models.Amount{Value: &types.Decimal{Value: "26"}, Currency: &currencyB},
				rate:   &models.ExchangeRate{CurrencyFrom: &currencyA, CurrencyTo: &currencyB, Value: 0},
			},
			wantErr: true,
		},
		"valid_3X_rate_reverse": {
			args: args{
				amount: &models.Amount{Value: &types.Decimal{Value: "26"}, Currency: &currencyB},
				rate:   &models.ExchangeRate{CurrencyFrom: &currencyA, CurrencyTo: &currencyB, Value: 3},
			},
			want:    &models.Amount{Value: &types.Decimal{Value: "8.666666667"}, Currency: &currencyA},
			wantErr: false,
		},
		"error_rate_currency_mismatch": {
			args: args{
				amount: &models.Amount{Value: &types.Decimal{Value: "26"}, Currency: &currencyC},
				rate:   &models.ExchangeRate{CurrencyFrom: &currencyA, CurrencyTo: &currencyB, Value: 3},
			},
			wantErr: true,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := ConvertCurrency(tt.args.amount, tt.args.rate)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConvertCurrency() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConvertCurrency() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSumTransactionFees(t *testing.T) {
	testCases := map[string]struct {
		paymentGatewayFees  []*models.PaymentGatewayFeeAmount
		processingFees      []*models.ProcessingFeeAmount
		servicingFees       []*models.ServicingFeeAmount
		expectedAmountValue string
	}{
		"test1": {
			paymentGatewayFees: []*models.PaymentGatewayFeeAmount{
				{Amount: &models.Amount{Value: &types.Decimal{Value: "10"}}},
				{Amount: &models.Amount{Value: &types.Decimal{Value: "30"}}},
			},
			processingFees: []*models.ProcessingFeeAmount{
				{Amount: &models.Amount{Value: &types.Decimal{Value: "20"}}},
			},
			servicingFees: []*models.ServicingFeeAmount{
				{Amount: &models.Amount{Value: &types.Decimal{Value: "30.3"}}},
			},
			expectedAmountValue: "90.3",
		},
		"test2": {
			paymentGatewayFees: []*models.PaymentGatewayFeeAmount{
				{Amount: &models.Amount{Value: &types.Decimal{Value: "1"}}},
				{Amount: &models.Amount{Value: &types.Decimal{Value: "2"}}},
				{Amount: &models.Amount{Value: &types.Decimal{Value: "3"}}},
			},
			servicingFees: []*models.ServicingFeeAmount{
				{Amount: &models.Amount{Value: &types.Decimal{Value: "30.3"}}},
			},
			expectedAmountValue: "36.3",
		},
		"test3": {
			paymentGatewayFees: []*models.PaymentGatewayFeeAmount{
				{Amount: &models.Amount{Value: &types.Decimal{Value: "10"}}},
				{Amount: &models.Amount{Value: &types.Decimal{Value: "20"}}},
				{Amount: &models.Amount{Value: &types.Decimal{Value: "30"}}},
			},
			expectedAmountValue: "60",
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			pgFeeAmount, err := SumFeeAmounts(tc.paymentGatewayFees, nil)
			assert.NoError(t, err)
			processingFeeAmount, err := SumFeeAmounts(tc.processingFees, nil)
			assert.NoError(t, err)
			servicingFeeAmount, err := SumFeeAmounts(tc.servicingFees, nil)
			assert.NoError(t, err)

			amount, err := SumAmounts(pgFeeAmount, processingFeeAmount, servicingFeeAmount)
			assert.NoError(t, err)
			assert.Equal(t, tc.expectedAmountValue, amount.GetValue().GetValue())
		})
	}
}

func uintPtr(v uint) *uint {
	return &v
}
