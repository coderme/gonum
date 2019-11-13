// Code generated by "go generate github.com/coderme/gonum/unit”; DO NOT EDIT.

// Copyright ©2014 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package unit

import (
	"errors"
	"fmt"
	"math"
	"unicode/utf8"
)

// Frequency represents a frequency in Hertz.
type Frequency float64

const Hertz Frequency = 1

// Unit converts the Frequency to a *Unit
func (f Frequency) Unit() *Unit {
	return New(float64(f), Dimensions{
		TimeDim: -1,
	})
}

// Frequency allows Frequency to implement a Frequencyer interface
func (f Frequency) Frequency() Frequency {
	return f
}

// From converts the unit into the receiver. From returns an
// error if there is a mismatch in dimension
func (f *Frequency) From(u Uniter) error {
	if !DimensionsMatch(u, Hertz) {
		*f = Frequency(math.NaN())
		return errors.New("Dimension mismatch")
	}
	*f = Frequency(u.Unit().Value())
	return nil
}

func (f Frequency) Format(fs fmt.State, c rune) {
	switch c {
	case 'v':
		if fs.Flag('#') {
			fmt.Fprintf(fs, "%T(%v)", f, float64(f))
			return
		}
		fallthrough
	case 'e', 'E', 'f', 'F', 'g', 'G':
		p, pOk := fs.Precision()
		w, wOk := fs.Width()
		const unit = " Hz"
		switch {
		case pOk && wOk:
			fmt.Fprintf(fs, "%*.*"+string(c), pos(w-utf8.RuneCount([]byte(unit))), p, float64(f))
		case pOk:
			fmt.Fprintf(fs, "%.*"+string(c), p, float64(f))
		case wOk:
			fmt.Fprintf(fs, "%*"+string(c), pos(w-utf8.RuneCount([]byte(unit))), float64(f))
		default:
			fmt.Fprintf(fs, "%"+string(c), float64(f))
		}
		fmt.Fprint(fs, unit)
	default:
		fmt.Fprintf(fs, "%%!%c(%T=%g Hz)", c, f, float64(f))
	}
}
