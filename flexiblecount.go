package gocourtlistener

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

// FlexibleCount is a custom type that can unmarshal both numbers and strings into an int.
type FlexibleCount int

// UnmarshalJSON implements the json.Unmarshaler interface.
func (fc *FlexibleCount) UnmarshalJSON(data []byte) error {
	// First, try to unmarshal as an integer.
	var i int
	if err := json.Unmarshal(data, &i); err == nil {
		*fc = FlexibleCount(i)
		return nil
	}

	// If that fails, unmarshal as a string.
	var s string
	if err := json.Unmarshal(data, &s); err == nil {
		// If the string looks like a URL (starts with "http"), then ignore it.
		if strings.HasPrefix(s, "http") {
			*fc = 0 // or you can set it to -1 to indicate "unknown"
			return nil
		}

		// Otherwise, attempt to convert the string to an int.
		i, err := strconv.Atoi(s)
		if err != nil {
			return fmt.Errorf("FlexibleCount: unable to convert string %q to int: %v", s, err)
		}
		*fc = FlexibleCount(i)
		return nil
	}

	return fmt.Errorf("FlexibleCount: unable to unmarshal %s", string(data))
}
