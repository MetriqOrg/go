package amount_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/lantah/go/amount"
	"github.com/lantah/go/xdr"
)

var Tests = []struct {
	S     string
	I     xdr.Int64
	valid bool
}{
	{"1000.000000", 1000000000, true},
	{"-1000.000000", -1000000000, true},
	{"1000.000001", 1000000001, true},
	{"1230.000001", 1230000001, true},
	{"1230.0000001", 0, false},
	{"9223372036854.775807", 9223372036854775807, true},
	{"9223372036854.775808", 0, false},
	{"9223372036855", 0, false},
	{"-9223372036854.775808", -9223372036854775808, true},
	{"-9223372036854.775809", 0, false},
	{"-9223372036855", 0, false},
	{"10000000000000.000000", 0, false},
	{"10000000000000", 0, false},
	{"-5.000000", -5000000, true},
	{"5.000000", 5000000, true},
	{"1.2345678", 0, false},
	// Expensive inputs:
	{strings.Repeat("1", 1000000), 0, false},
	{"1E9223372036854775807", 0, false},
	{"1e9223372036854775807", 0, false},
	{"Inf", 0, false},
}

func TestParse(t *testing.T) {
	for _, v := range Tests {
		o, err := amount.Parse(v.S)
		if !v.valid && err == nil {
			t.Errorf("expected err for input %s", v.S)
			continue
		}
		if v.valid && err != nil {
			t.Errorf("couldn't parse %s: %v", v.S, err)
			continue
		}

		if o != v.I {
			t.Errorf("%s parsed to %d, not %d", v.S, o, v.I)
		}
	}
}

func TestString(t *testing.T) {
	for _, v := range Tests {
		if !v.valid {
			continue
		}

		o := amount.String(v.I)

		if o != v.S {
			t.Errorf("%d stringified to %s, not %s", v.I, o, v.S)
		}
	}
}

func TestIntStringToAmount(t *testing.T) {
	var testCases = []struct {
		Output string
		Input  string
		Valid  bool
	}{
		{"1000.000000", "1000000000", true},
		{"-1000.000000", "-1000000000", true},
		{"1000.000001", "1000000001", true},
		{"1230.000001", "1230000001", true},
		{"9223372036854.775807", "9223372036854775807", true},
		{"9223372036854.775808", "9223372036854775808", true},
		{"922337.203686", "922337203686", true},
		{"-9223372036854.775808", "-9223372036854775808", true},
		{"-9223372036854.775809", "-9223372036854775809", true},
		{"-922337.203686", "-922337203686", true},
		{"10000000000000.000000", "10000000000000000000", true},
		{"0.000000", "0", true},
		// Expensive inputs when using big.Rat:
		{"100000000000000.000000", "1" + strings.Repeat("0", 20), true},
		{"-100000000000000.000000", "-1" + strings.Repeat("0", 20), true},
		{"1" + strings.Repeat("0", 1000-7) + ".000000", "1" + strings.Repeat("0", 1000), true},
		{"1" + strings.Repeat("0", 1000000-7) + ".000000", "1" + strings.Repeat("0", 1000000), true},
		// Invalid inputs
		{"", "nan", false},
		{"", "", false},
		{"", "-", false},
		{"", "1E9223372036854775807", false},
		{"", "1e9223372036854775807", false},
		{"", "Inf", false},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%s to %s (valid = %t)", tc.Input, tc.Output, tc.Valid), func(t *testing.T) {
			o, err := amount.IntStringToAmount(tc.Input)

			if !tc.Valid && err == nil {
				t.Errorf("expected err for input %s (output: %s)", tc.Input, tc.Output)
				return
			}
			if tc.Valid && err != nil {
				t.Errorf("couldn't parse %s: %v", tc.Input, err)
				return
			}

			if o != tc.Output {
				t.Errorf("%s converted to %s, not %s", tc.Input, o, tc.Output)
			}
		})
	}

}
