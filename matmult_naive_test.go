// Copyright 2018 Changkun Ou. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package gomat

import (
	"fmt"
	"testing"
)

func TestMatrix_MultNaive(t *testing.T) {
	A, err := New(2, 3)(
		1, 2, 3,
		2, 3, 1,
	)
	if err != nil {
		t.Errorf("New(2, 3) error, expect nil")
	}
	B, err := New(3, 1)(
		3,
		2,
		1,
	)
	if err != nil {
		t.Errorf("New(3, 1) error, expect nil")
	}

	t.Run("DotNaiveIJK()", func(t *testing.T) {
		C, err := New(2, 1)(
			0,
			0,
		)
		T, err := New(2, 1)(
			10,
			13,
		)
		if err != nil {
			t.Errorf("New(2, 1) error, expect nil")
		}
		err = A.DotNaiveIJK(B, C)
		if err != nil {
			t.Errorf("DotNaiveIJK(B, C) error, expect nil")
		}
		if !T.Equal(C) {
			t.Errorf("DotNaiveIJK() not euqal, expect euqal, got:")
			C.Print()
		}
	})
	t.Run("DotNaiveIKJ()", func(t *testing.T) {
		C, err := New(2, 1)(
			0,
			0,
		)
		T, err := New(2, 1)(
			10,
			13,
		)
		if err != nil {
			t.Errorf("New(2, 1) error, expect nil")
		}
		err = A.DotNaiveIKJ(B, C)
		if err != nil {
			t.Errorf("DotNaiveIKJ(B, C) error, expect nil")
		}
		if !T.Equal(C) {
			t.Errorf("DotNaiveIKJ() not euqal, expect euqal, got:")
			C.Print()
		}
	})
	t.Run("DotNaiveJIK()", func(t *testing.T) {
		C, err := New(2, 1)(
			0,
			0,
		)
		T, err := New(2, 1)(
			10,
			13,
		)
		if err != nil {
			t.Errorf("New(2, 1) error, expect nil")
		}
		err = A.DotNaiveJIK(B, C)
		if err != nil {
			t.Errorf("DotNaiveJIK(B, C) error, expect nil")
		}
		if !T.Equal(C) {
			t.Errorf("DotNaiveJIK() not euqal, expect euqal, got:")
			C.Print()
		}
	})
	t.Run("DotNaiveJKI()", func(t *testing.T) {
		C, err := New(2, 1)(
			0,
			0,
		)
		T, err := New(2, 1)(
			10,
			13,
		)
		if err != nil {
			t.Errorf("New(2, 1) error, expect nil")
		}
		err = A.DotNaiveJKI(B, C)
		if err != nil {
			t.Errorf("DotNaiveJKI(B, C) error, expect nil")
		}
		if !T.Equal(C) {
			t.Errorf("DotNaiveJKI() not euqal, expect euqal, got:")
			C.Print()
		}
	})
	t.Run("DotNaiveKIJ()", func(t *testing.T) {
		C, err := New(2, 1)(
			0,
			0,
		)
		T, err := New(2, 1)(
			10,
			13,
		)
		if err != nil {
			t.Errorf("New(2, 1) error, expect nil")
		}
		err = A.DotNaiveKIJ(B, C)
		if err != nil {
			t.Errorf("DotNaiveKIJ(B, C) error, expect nil")
		}
		if !T.Equal(C) {
			t.Errorf("DotNaiveKIJ() not euqal, expect euqal, got:")
			C.Print()
		}
	})
	t.Run("DotNaiveKJI()", func(t *testing.T) {
		C, err := New(2, 1)(
			0,
			0,
		)
		T, err := New(2, 1)(
			10,
			13,
		)
		if err != nil {
			t.Errorf("New(2, 1) error, expect nil")
		}
		err = A.DotNaiveKJI(B, C)
		if err != nil {
			t.Errorf("DotNaiveKJI(B, C) error, expect nil")
		}
		if !T.Equal(C) {
			t.Errorf("DotNaiveKJI() not euqal, expect euqal, got:")
			C.Print()
		}
	})
	t.Run("DotNaiveIJKP()", func(t *testing.T) {
		C, err := New(2, 1)(
			0,
			0,
		)
		T, err := New(2, 1)(
			10,
			13,
		)
		if err != nil {
			t.Errorf("New(2, 1) error, expect nil")
		}
		err = A.DotNaiveIJKP(B, C)
		if err != nil {
			t.Errorf("DotNaiveIJKP(B, C) error, expect nil")
		}
		if !T.Equal(C) {
			t.Errorf("DotNaiveIJKP() not euqal, expect euqal, got:")
			C.Print()
		}
	})
}

// ----------------------- benchmarks ----------------------------

func BenchmarkNaive(b *testing.B) {
	for n := 0; n < 100; n++ {
		A := Rand(n, n)
		B := Rand(n, n)
		C := Zero(n, n)

		b.Run(fmt.Sprintf("DotNaiveIJK() %dx%d", n, n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				A.DotNaiveIJK(B, C)
			}
		})
		b.Run(fmt.Sprintf("DotNaiveIKJ()-%dx%d", n, n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				A.DotNaiveIKJ(B, C)
			}
		})
		b.Run(fmt.Sprintf("DotNaiveJIK()-%dx%d", n, n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				A.DotNaiveJIK(B, C)
			}
		})
		b.Run(fmt.Sprintf("DotNaiveJKI()-%dx%d", n, n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				A.DotNaiveJKI(B, C)
			}
		})
		b.Run(fmt.Sprintf("DotNaiveKIJ()-%dx%d", n, n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				A.DotNaiveKIJ(B, C)
			}
		})
		b.Run(fmt.Sprintf("DotNaiveKJI()-%dx%d", n, n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				A.DotNaiveKJI(B, C)
			}
		})
	}
}

func BenchmarkNaiveP(b *testing.B) {
	for n := 55; n < 100; n++ {
		A := Rand(n, n)
		B := Rand(n, n)
		C := Zero(n, n)

		b.Run(fmt.Sprintf("DotNaiveIJKP() %dx%d", n, n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				A.DotNaiveIJKP(B, C)
			}
		})
		b.Run(fmt.Sprintf("DotNaiveIKJP()-%dx%d", n, n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				A.DotNaiveIKJP(B, C)
			}
		})
		b.Run(fmt.Sprintf("DotNaiveJIKP()-%dx%d", n, n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				A.DotNaiveJIKP(B, C)
			}
		})
		b.Run(fmt.Sprintf("DotNaiveJKIP()-%dx%d", n, n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				A.DotNaiveJKIP(B, C)
			}
		})
		b.Run(fmt.Sprintf("DotNaiveKIJP()-%dx%d", n, n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				A.DotNaiveKIJP(B, C)
			}
		})
		b.Run(fmt.Sprintf("DotNaiveKJIP()-%dx%d", n, n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				A.DotNaiveKJIP(B, C)
			}
		})
	}
}
