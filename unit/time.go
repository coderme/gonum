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

// Time represents a duration in seconds.
type Time float64

const (
	Second Time = 1

	Minute = 60 * Second
	Hour   = 60 * Minute
)

// Unit converts the Time to a *Unit
func (t Time) Unit() *Unit {
	return New(float64(t), Dimensions{
		TimeDim: 1,
	})
}

// Time allows Time to implement a Timer interface
func (t Time) Time() Time {
	return t
}

// From converts the unit into the receiver. From returns an
// error if there is a mismatch in dimension
func (t *Time) From(u Uniter) error {
	if !DimensionsMatch(u, Second) {
		*t = Time(math.NaN())
		return errors.New("Dimension mismatch")
	}
	*t = Time(u.Unit().Value())
	return nil
}

func (t Time) Format(fs fmt.State, c rune) {
	switch c {
	case 'v':
		if fs.Flag('#') {
			fmt.Fprintf(fs, "%T(%v)", t, float64(t))
			return
		}
		fallthrough
	case 'e', 'E', 'f', 'F', 'g', 'G':
		p, pOk := fs.Precision()
		w, wOk := fs.Width()
		const unit = " s"
		switch {
		case pOk && wOk:
			fmt.Fprintf(fs, "%*.*"+string(c), pos(w-utf8.RuneCount([]byte(unit))), p, float64(t))
		case pOk:
			fmt.Fprintf(fs, "%.*"+string(c), p, float64(t))
		case wOk:
			fmt.Fprintf(fs, "%*"+string(c), pos(w-utf8.RuneCount([]byte(unit))), float64(t))
		default:
			fmt.Fprintf(fs, "%"+string(c), float64(t))
		}
		fmt.Fprint(fs, unit)
	default:
		fmt.Fprintf(fs, "%%!%c(%T=%g s)", c, t, float64(t))
	}
}
