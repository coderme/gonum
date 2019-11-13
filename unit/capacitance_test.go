// Code generated by "go generate github.com/coderme/gonum/unit; DO NOT EDIT.

// Copyright ©2019 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package unit

import (
	"fmt"
	"testing"
)

func TestCapacitanceFormat(t *testing.T) {
	for _, test := range []struct {
		value  Capacitance
		format string
		want   string
	}{
		{1.23456789, "%v", "1.23456789 F"},
		{1.23456789, "%.1v", "1 F"},
		{1.23456789, "%20.1v", "                 1 F"},
		{1.23456789, "%20v", "        1.23456789 F"},
		{1.23456789, "%1v", "1.23456789 F"},
		{1.23456789, "%#v", "unit.Capacitance(1.23456789)"},
		{1.23456789, "%s", "%!s(unit.Capacitance=1.23456789 F)"},
	} {
		got := fmt.Sprintf(test.format, test.value)
		if got != test.want {
			t.Errorf("Format %q %v: got: %q want: %q", test.format, float64(test.value), got, test.want)
		}
	}
}
