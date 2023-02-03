package errors

import "time"

type (
	// LocalError defines the contract that allows implementing systems to define their own Reason codes
	LocalError interface {
		Reason() Reason
		Type() Type
		Domain() Domain
		Message() string
	}
)

// NewLocal is a constructor to create a minimally viable Error instance given an a root cause and implementation of
// LocalError.
func NewLocal(cause error, local LocalError) Error {
	return Error{
		Err:       cause,
		Type:      local.Type(),
		Reason:    local.Reason(),
		Domain:    local.Domain(),
		Message:   local.Message(),
		Timestamp: time.Now().UTC(),
	}
}

// WithMetaMap is a fluent style method to set the Metadata map on the Error instance.
func (e Error) WithMetaMap(data map[string]string) Error {
	e.Metadata = data
	return e
}

// WithMetaPairs is a fluent style convenience method to define the Metadata map as a series of key-value pairs.
// Destructive, sets the Metadata field with a new Map before iterating over the parameters.
// This will iterate over the variadic in pairs, and take the even index as the key, and the odd index as the value,
// as such, an odd number of values passed will see the last value ignored.
func (e Error) WithMetaPairs(pairs ...string) Error {
	e.Metadata = map[string]string{}
	for idx := 0; idx < len(pairs); idx += 2 {
		e.Metadata[pairs[idx]] = pairs[idx+1]
	}
	return e
}

// WithMetaPair is a fluent style convenience method to append to the Metadata map.
// If the Metadata map doesn't exist, it sets a new one, otherwise appends to the existing map.
func (e Error) WithMetaPair(key, value string) Error {
	if e.Metadata == nil {
		e.Metadata = make(map[string]string)
	}
	e.Metadata[key] = value
	return e
}
