// Copyright ©2015 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build cblas

package mat

import (
	"github.com/coderme/gonum/blas/blas64"
	"github.com/coderme/netlib/blas/netlib"
)

func init() {
	blas64.Use(netlib.Implementation{})
}
