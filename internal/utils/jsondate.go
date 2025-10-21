package utils

import (
	"encoding/json"
	"errors"
	"strings"
	"time"
)

// JSONDate custom type untuk parsing tanggal dari JSON
// Mendukung format "YYYY-MM-DD"
type JSONDate time.Time

// UnmarshalJSON untuk parsing string ke time.Time otomatis
func (d *JSONDate) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), `"`)
	if s == "" {
		return nil
	}

	if t, err := time.Parse(time.RFC3339, s); err == nil {
		*d = JSONDate(t)
		return nil
	}

	if t, err := time.Parse("2006-01-02", s); err == nil {
		*d = JSONDate(t)
		return nil
	}

	return errors.New("invalid date format")
}

// MarshalJSON untuk konversi time.Time ke string otomatis
func (d JSONDate) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(d).Format(time.RFC3339))
}

// Time untuk konversi ke time.Time
func (d JSONDate) Time() time.Time {
	return time.Time(d)
}
