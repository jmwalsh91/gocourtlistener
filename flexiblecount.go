package gocourtlistener

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// FlexibleCount is a custom type that can unmarshal both numbers and strings into an int.
type FlexibleCount int

// UnmarshalJSON implements the json.Unmarshaler interface.
func (fc *FlexibleCount) UnmarshalJSON(data []byte) error {
	// Try to unmarshal as an int.
	var i int
	if err := json.Unmarshal(data, &i); err == nil {
		*fc = FlexibleCount(i)
		return nil
	}

	// If that fails, try to unmarshal as a string.
	var s string
	if err := json.Unmarshal(data, &s); err == nil {
		i, err := strconv.Atoi(s)
		if err != nil {
			return fmt.Errorf("FlexibleCount: unable to convert string %q to int: %v", s, err)
		}
		*fc = FlexibleCount(i)
		return nil
	}

	return fmt.Errorf("FlexibleCount: unable to unmarshal %s", string(data))
}
