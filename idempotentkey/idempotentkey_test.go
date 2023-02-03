package idempotentkey

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestIdempotentUUID(t *testing.T) {
	testCases := []struct {
		name       string
		input1Seed string
		input2Seed string
		equal      bool
	}{
		{
			name:       "equal inputs 1",
			input1Seed: "dc559b7b-e2f7-4863-8a60-98dce53fb93f-funding",
			input2Seed: "dc559b7b-e2f7-4863-8a60-98dce53fb93f-funding",
			equal:      true,
		},
		{
			name:       "equal inputs 2",
			input1Seed: "dc559b7b-e2f7-4863-8a60-98dce53fb93f-execution",
			input2Seed: "dc559b7b-e2f7-4863-8a60-98dce53fb93f-execution",
			equal:      true,
		},
		{
			name:       "unequal inputs 1",
			input1Seed: "dc559b7b-e2f7-4863-8a60-98dce53fb93f-funding",
			input2Seed: "dc559b7b-e2f7-4863-8a60-98dce53fb93f-payout",
			equal:      false,
		},
		{
			name:       "unequal inputs 2",
			input1Seed: "dc559b7b-e2f7-4863-8a60-98dce53fb93f-kjasdhauhdajskkfsdkhkajsdjaksdhaf",
			input2Seed: "dc559b7b-e2f7-4863-8a60-98dce53fb93f-kjasdhauhdajskkfsdkhkajsdjaksdhaf-ssasdjjkkasd",
			equal:      false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got1 := GenerateIdempotentUUID(tc.input1Seed)
			got2 := GenerateIdempotentUUID(tc.input2Seed)
			require.Equal(t, tc.equal, got1 == got2)
		})
	}
}
