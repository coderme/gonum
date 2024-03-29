// Copyright ©2017 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package distuv

import (
	"sort"
	"testing"

	"golang.org/x/exp/rand"
	"github.com/coderme/gonum/floats"
)

func TestParetoProb(t *testing.T) {
	for _, test := range []struct {
		x, xm, alpha, want float64
	}{
		{0, 1, 1, 0},
		{0.5, 1, 1, 0},
		{1, 1, 1, 1.0},
		{1.5, 1, 1, 0.444444444444444},
		{2, 1, 1, 0.25},
		{2.5, 1, 1, 0.16},
		{3, 1, 1, 0.1111111111111111},
		{3.5, 1, 1, 0.081632653061224},
		{4, 1, 1, 0.0625},
		{4.5, 1, 1, 0.049382716049383},
		{5, 1, 1, 0.04},

		{0, 1, 2, 0},
		{0.5, 1, 2, 0},
		{1, 1, 2, 2},
		{1.5, 1, 2, 0.592592592592593},
		{2, 1, 2, 0.25},
		{2.5, 1, 2, 0.128},
		{3, 1, 2, 0.074074074074074},
		{3.5, 1, 2, 0.046647230320700},
		{4, 1, 2, 0.03125},
		{4.5, 1, 2, 0.021947873799726},
		{5, 1, 2, 0.016},

		{0, 1, 3, 0},
		{0.5, 1, 3, 0},
		{1, 1, 3, 3.0},
		{1.5, 1, 3, 0.592592592592593},
		{2, 1, 3, 0.1875},
		{2.5, 1, 3, 0.0768},
		{3, 1, 3, 0.037037037037037},
		{3.5, 1, 3, 0.019991670137443},
		{4, 1, 3, 0.011718750000000},
		{4.5, 1, 3, 0.007315957933242},
		{5, 1, 3, 0.0048},
	} {
		pdf := Pareto{test.xm, test.alpha, nil}.Prob(test.x)
		if !floats.EqualWithinAbsOrRel(pdf, test.want, 1e-10, 1e-10) {
			t.Errorf("Pdf mismatch, x = %v, xm = %v, alpha = %v. Got %v, want %v", test.x, test.xm, test.alpha, pdf, test.want)
		}
	}
}

func TestParetoCDF(t *testing.T) {
	for _, test := range []struct {
		x, xm, alpha, want float64
	}{
		{0, 1, 1, 0},
		{0.5, 1, 1, 0},
		{1, 1, 1, 0},
		{1.5, 1, 1, 0.333333333333333},
		{2, 1, 1, 0.5},
		{2.5, 1, 1, 0.6},
		{3, 1, 1, 0.666666666666667},
		{3.5, 1, 1, 0.714285714285714},
		{4, 1, 1, 0.75},
		{4.5, 1, 1, 0.777777777777778},
		{5, 1, 1, 0.80},
		{5.5, 1, 1, 0.818181818181818},
		{6, 1, 1, 0.833333333333333},
		{6.5, 1, 1, 0.846153846153846},
		{7, 1, 1, 0.857142857142857},
		{7.5, 1, 1, 0.866666666666667},
		{8, 1, 1, 0.875},
		{8.5, 1, 1, 0.882352941176471},
		{9, 1, 1, 0.888888888888889},
		{9.5, 1, 1, 0.894736842105263},
		{10, 1, 1, 0.90},

		{0, 1, 2, 0},
		{0.5, 1, 2, 0},
		{1, 1, 2, 0},
		{1.5, 1, 2, 0.555555555555556},
		{2, 1, 2, 0.75},
		{2.5, 1, 2, 0.84},
		{3, 1, 2, 0.888888888888889},
		{3.5, 1, 2, 0.918367346938776},
		{4, 1, 2, 0.9375},
		{4.5, 1, 2, 0.950617283950617},
		{5, 1, 2, 0.96},
		{5.5, 1, 2, 0.966942148760331},
		{6, 1, 2, 0.972222222222222},
		{6.5, 1, 2, 0.976331360946746},
		{7, 1, 2, 0.979591836734694},
		{7.5, 1, 2, 0.982222222222222},
		{8, 1, 2, 0.984375000000000},
		{8.5, 1, 2, 0.986159169550173},
		{9, 1, 2, 0.987654320987654},
		{9.5, 1, 2, 0.988919667590028},
		{10, 1, 2, 0.99},

		{0, 1, 3, 0},
		{0.5, 1, 3, 0},
		{1, 1, 3, 0},
		{1.5, 1, 3, 0.703703703703704},
		{2, 1, 3, 0.875},
		{2.5, 1, 3, 0.936},
		{3, 1, 3, 0.962962962962963},
		{3.5, 1, 3, 0.976676384839650},
		{4, 1, 3, 0.984375000000000},
		{4.5, 1, 3, 0.989026063100137},
		{5, 1, 3, 0.992},
		{5.5, 1, 3, 0.993989481592787},
		{6, 1, 3, 0.995370370370370},
		{6.5, 1, 3, 0.996358670914884},
		{7, 1, 3, 0.997084548104956},
		{7.5, 1, 3, 0.997629629629630},
		{8, 1, 3, 0.998046875000000},
		{8.5, 1, 3, 0.998371667005903},
		{9, 1, 3, 0.998628257887517},
		{9.5, 1, 3, 0.998833649220003},
		{10, 1, 3, 0.999},
	} {
		cdf := Pareto{test.xm, test.alpha, nil}.CDF(test.x)
		if !floats.EqualWithinAbsOrRel(cdf, test.want, 1e-10, 1e-10) {
			t.Errorf("CDF mismatch, x = %v, xm = %v, alpha = %v. Got %v, want %v", test.x, test.xm, test.alpha, cdf, test.want)
		}
	}
}

func TestPareto(t *testing.T) {
	src := rand.New(rand.NewSource(1))
	for i, p := range []Pareto{
		{1, 10, src},
		{1, 20, src},
	} {
		testPareto(t, p, i)
	}
}

func testPareto(t *testing.T, p Pareto, i int) {
	tol := 1e-2
	const n = 1e6
	const bins = 50
	x := make([]float64, n)
	generateSamples(x, p)
	sort.Float64s(x)

	testRandLogProbContinuous(t, i, 0, x, p, tol, bins)
	checkMean(t, i, x, p, tol)
	checkVarAndStd(t, i, x, p, tol)
	checkExKurtosis(t, i, x, p, 7e-2)
	checkProbContinuous(t, i, x, p, 1e-3)
}
