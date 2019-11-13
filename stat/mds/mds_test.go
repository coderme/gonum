// Copyright ©2018 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mds

import (
	"testing"

	"github.com/coderme/gonum/floats"
	"github.com/coderme/gonum/mat"
)

var torgersonScalingTests = []struct {
	dis      mat.Symmetric
	wantK    int
	want     *mat.Dense
	wantVals []float64
}{
	// All expected values obtained from running R cmdscale with the input here.
	{
		// Data from http://rosetta.reltech.org/TC/v15/Mapping/data/dist-Aus.csv
		dis: mat.NewSymDense(8, []float64{
			0, 1328, 1600, 2616, 1161, 653, 2130, 1161,
			1328, 0, 1962, 1289, 2463, 1889, 1991, 2026,
			1600, 1962, 0, 2846, 1788, 1374, 3604, 732,
			2616, 1289, 2846, 0, 3734, 3146, 2652, 3146,
			1161, 2463, 1788, 3734, 0, 598, 3008, 1057,
			653, 1889, 1374, 3146, 598, 0, 2720, 713,
			2130, 1991, 3604, 2652, 3008, 2720, 0, 3288,
			1161, 2026, 732, 3146, 1057, 713, 3288, 0,
		}),
		wantK: 4,
		want: mat.NewDense(8, 4, []float64{
			-208.3266, 369.5373, 80.544010, 7.078974,
			904.8449, -356.0745, 92.309541, -19.077511,
			-925.9941, -1067.8589, -38.585847, -6.696700,
			1933.8035, -1129.8610, -50.099688, 7.344750,
			-1318.6333, 704.3759, -56.673300, -17.419249,
			-858.3951, 319.5948, 24.186572, 18.138426,
			1591.6571, 1511.3688, -43.308441, 1.043611,
			-1118.9564, -351.0825, -8.372847, 9.587699,
		}),
		wantVals: []float64{
			1.172027697614e+07, 5.686036184502e+06, 2.474981412369e+04, 1.238300279584e+03,
			-5.444455825182e-10, -2.471028925712e+02, -2.412961483429e+03, -8.452858567111e+04,
		},
	},
	{
		// R eurodist dataset.
		dis: mat.NewSymDense(21, []float64{
			0, 3313, 2963, 3175, 3339, 2762, 3276, 2610, 4485, 2977, 3030, 4532, 2753, 3949, 2865, 2282, 2179, 3000, 817, 3927, 1991,
			3313, 0, 1318, 1326, 1294, 1498, 2218, 803, 1172, 2018, 1490, 1305, 645, 636, 521, 1014, 1365, 1033, 1460, 2868, 1802,
			2963, 1318, 0, 204, 583, 206, 966, 677, 2256, 597, 172, 2084, 690, 1558, 1011, 925, 747, 285, 1511, 1616, 1175,
			3175, 1326, 204, 0, 460, 409, 1136, 747, 2224, 714, 330, 2052, 739, 1550, 1059, 1077, 977, 280, 1662, 1786, 1381,
			3339, 1294, 583, 460, 0, 785, 1545, 853, 2047, 1115, 731, 1827, 789, 1347, 1101, 1209, 1160, 340, 1794, 2196, 1588,
			2762, 1498, 206, 409, 785, 0, 760, 1662, 2436, 460, 269, 2290, 714, 1764, 1035, 911, 583, 465, 1497, 1403, 937,
			3276, 2218, 966, 1136, 1545, 760, 0, 1418, 3196, 460, 269, 2971, 1458, 2498, 1778, 1537, 1104, 1176, 2050, 650, 1455,
			2610, 803, 677, 747, 853, 1662, 1418, 0, 1975, 1118, 895, 1936, 158, 1439, 425, 328, 591, 513, 995, 2068, 1019,
			4485, 1172, 2256, 2224, 2047, 2436, 3196, 1975, 0, 2897, 2428, 676, 1817, 698, 1693, 2185, 2565, 1971, 2631, 3886, 2974,
			2977, 2018, 597, 714, 1115, 460, 460, 1118, 2897, 0, 550, 2671, 1159, 2198, 1479, 1238, 805, 877, 1751, 949, 1155,
			3030, 1490, 172, 330, 731, 269, 269, 895, 2428, 550, 0, 2280, 863, 1730, 1183, 1098, 851, 457, 1683, 1500, 1205,
			4532, 1305, 2084, 2052, 1827, 2290, 2971, 1936, 676, 2671, 2280, 0, 1178, 668, 1762, 2250, 2507, 1799, 2700, 3231, 2937,
			2753, 645, 690, 739, 789, 714, 1458, 158, 1817, 1159, 863, 1178, 0, 1281, 320, 328, 724, 471, 1048, 2108, 1157,
			3949, 636, 1558, 1550, 1347, 1764, 2498, 1439, 698, 2198, 1730, 668, 1281, 0, 1157, 1724, 2010, 1273, 2097, 3188, 2409,
			2865, 521, 1011, 1059, 1101, 1035, 1778, 425, 1693, 1479, 1183, 1762, 320, 1157, 0, 618, 1109, 792, 1011, 2428, 1363,
			2282, 1014, 925, 1077, 1209, 911, 1537, 328, 2185, 1238, 1098, 2250, 328, 1724, 618, 0, 331, 856, 586, 2187, 898,
			2179, 1365, 747, 977, 1160, 583, 1104, 591, 2565, 805, 851, 2507, 724, 2010, 1109, 331, 0, 821, 946, 1754, 428,
			3000, 1033, 285, 280, 340, 465, 1176, 513, 1971, 877, 457, 1799, 471, 1273, 792, 856, 821, 0, 1476, 1827, 1249,
			817, 1460, 1511, 1662, 1794, 1497, 2050, 995, 2631, 1751, 1683, 2700, 1048, 2097, 1011, 586, 946, 1476, 0, 2707, 1209,
			3927, 2868, 1616, 1786, 2196, 1403, 650, 2068, 3886, 949, 1500, 3231, 2108, 3188, 2428, 2187, 1754, 1827, 2707, 0, 2105,
			1991, 1802, 1175, 1381, 1588, 937, 1455, 1019, 2974, 1155, 1205, 2937, 1157, 2409, 1363, 898, 428, 1249, 1209, 2105, 0,
		}),
		// Note that k here is 12 despite the result from R's cmdscale.
		// This is due to disparity between the ED performed by R and Gonum.
		// See https://github.com/gonum/gonum/issues/768
		wantK: 12,
		want: mat.NewDense(21, 12, []float64{
			2290.274680, 1798.80293, 53.79314, -103.826958, -156.955115, 54.755434, -47.6768205, 1.241284, -14.893196, -6.366664, 4.818373, 0,
			-825.382790, 546.81148, -113.85842, 84.585831, 291.440759, -33.046236, -74.5267190, 3.766233, 225.620420, -21.270973, 22.050484, 0,
			59.183341, -367.08135, 177.55291, 38.797514, -95.620447, 40.058268, 2.3212184, 34.351715, -2.262151, -129.298820, 75.568686, 0,
			-82.845973, -429.91466, 300.19274, 106.353695, -180.446140, 31.336985, 88.6540176, 5.102633, 87.657362, 86.867202, 156.097411, 0,
			-352.499435, -290.90843, 457.35294, 111.449150, -417.496682, -138.972847, 46.7759317, 23.806769, 108.307203, 99.578699, -103.623270, 0,
			293.689633, -405.31194, 360.09323, -636.202379, 159.392662, -9.560652, 33.1089500, 9.193680, 19.417287, 7.925230, -1.272483, 0,
			681.931545, -1108.64478, 26.09257, 151.693056, 254.114892, 197.254208, -176.0359152, -6.508596, -59.829991, 64.904841, -26.466935, 0,
			-9.423364, 240.40600, -344.20659, 656.121110, -138.082259, -19.891841, -0.3128244, -20.167733, -18.250209, -11.175662, 2.160076, 0,
			-2048.449113, 642.45854, 167.86631, 78.621423, 239.881353, -51.009028, 60.7527717, 125.492010, -144.940590, 35.754867, 13.633563, 0,
			561.108970, -773.36929, 80.91722, 48.548472, -129.512374, 20.426493, 207.7920058, 23.646986, -166.184382, -64.027105, -5.187863, 0,
			164.921799, -549.36704, 270.82327, 116.886334, 62.711250, 117.489888, -252.4906676, -47.193881, -62.053055, 75.466644, 4.907670, 0,
			-1935.040811, 49.12514, -483.02056, -315.241752, -277.696589, -92.444183, -76.2549266, -28.529861, -26.279840, -2.503630, 7.050184, 0,
			-226.423236, 187.08779, -358.43234, -257.737009, -190.609088, -110.666015, -119.7638108, -43.118326, -76.066370, 89.289877, 18.820576, 0,
			-1423.353697, 305.87513, 253.26763, 2.478812, 5.366334, 205.060613, -45.8020725, -75.244576, -4.303057, -150.778068, -23.118399, 0,
			-299.498710, 388.80726, -109.17417, 12.651217, 231.405239, 154.749482, 237.4507512, -227.372584, 25.449825, 71.404170, -35.167507, 0,
			260.878046, 416.67381, -171.52428, 20.926369, 194.345825, -19.740272, 121.3353507, 184.415363, -5.563374, 90.507909, -22.562758, 0,
			587.675679, 81.18224, -75.88485, 13.080496, 114.688714, -185.462926, -47.4707495, 209.170525, 39.452744, -65.352092, -23.127035, 0,
			-156.836257, -211.13911, 131.30852, 27.089432, -100.800743, 6.794791, -25.6913210, -7.904762, 39.731234, -79.091817, -75.116800, 0,
			709.413282, 1109.36665, -179.83052, -109.895049, -90.243055, 316.665218, 9.9070412, 9.976867, 6.316941, -13.285987, 1.496742, 0,
			839.445911, -1836.79055, -541.35188, -108.755016, 25.155430, 41.399927, 71.8793196, 9.438407, 73.554358, -35.415426, -6.021569, 0,
			911.230500, 205.93020, 98.02313, 62.375253, 198.960032, -525.197308, -13.9515306, -183.562155, -44.881159, -43.133195, 15.060856, 0,
		}),
		wantVals: []float64{
			1.953837708954e+07, 1.185655533400e+07, 1.528844467987e+06, 1.118741950509e+06,
			7.893472026801e+05, 5.816552067198e+05, 2.623192077011e+05, 1.925975616762e+05,
			1.450845349644e+05, 1.079673069262e+05, 5.139484110774e+04, 7.598368987955e-11,
			-9.496124219168e+03, -5.305819566947e+04, -1.322165749977e+05, -2.573360255637e+05,
			-3.326719007160e+05, -5.162522542344e+05, -9.191490984121e+05, -1.006503960172e+06,
			-2.251844331736e+06,
		},
	},
}

func TestTorgersonScaling(t *testing.T) {
	for i, test := range torgersonScalingTests {
		_, c := test.dis.Dims()
		var got mat.Dense
		gotK, gotVals := TorgersonScaling(&got, make([]float64, c), test.dis)
		if gotK == 0 {
			t.Error("unexpected scaling failure")
			continue
		}
		if gotK != test.wantK {
			t.Errorf("unexpected k for test %d: got:%d want:%d", i, gotK, test.wantK)
		}
		if !mat.EqualApprox(colAbs{&got}, colAbs{test.want}, 1e-5) {
			t.Errorf("unexpected result for test %d:\ngot:\n%.5f\nwant:\n%.5f",
				i, mat.Formatted(&got), mat.Formatted(test.want))
		}
		if !floats.EqualApprox(gotVals, test.wantVals, 1e-10) {
			t.Errorf("unexpected Eigenvalues for test %d:\ngot: %.12e\nwant:%.12e", i, gotVals, test.wantVals)
		}
	}
}

// colAbs returns the value of columns reflected
// such that the first row is positive.
type colAbs struct {
	mat.Matrix
}

func (m colAbs) At(i, j int) float64 {
	if m.Matrix.At(0, j) < 0 {
		return -m.Matrix.At(i, j)
	}
	return m.Matrix.At(i, j)
}
