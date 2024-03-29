// Code generated by "go generate github.com/coderme/gonum/unit/constant”; DO NOT EDIT.

// Copyright ©2019 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package constant

import (
	"fmt"

	"github.com/coderme/gonum/unit"
)

// Avogadro is the Avogadro constant (A), the number of constituent particles contained in one mole of a substance.
// The dimension of Avogadro is mol^-1. The standard uncertainty of the constant is 7.4e+15 mol^-1.
const Avogadro = avogadroUnits(6.022140857e+23)

type avogadroUnits float64

// Unit converts the avogadroUnits to a *unit.Unit
func (cnst avogadroUnits) Unit() *unit.Unit {
	return unit.New(float64(cnst), unit.Dimensions{
		unit.MoleDim: -1,
	})
}

func (cnst avogadroUnits) Format(fs fmt.State, c rune) {
	switch c {
	case 'v':
		if fs.Flag('#') {
			fmt.Fprintf(fs, "%T(%v)", cnst, float64(cnst))
			return
		}
		fallthrough
	case 'e', 'E', 'f', 'F', 'g', 'G':
		p, pOk := fs.Precision()
		w, wOk := fs.Width()
		switch {
		case pOk && wOk:
			fmt.Fprintf(fs, "%*.*"+string(c), w, p, cnst.Unit())
		case pOk:
			fmt.Fprintf(fs, "%.*"+string(c), p, cnst.Unit())
		case wOk:
			fmt.Fprintf(fs, "%*"+string(c), w, cnst.Unit())
		default:
			fmt.Fprintf(fs, "%"+string(c), cnst.Unit())
		}
	default:
		fmt.Fprintf(fs, "%%!"+string(c)+"(constant.avogadroUnits=%v mol^-1)", float64(cnst))
	}
}
