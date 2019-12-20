package misc_test

import (
	"fmt"
	"math"
	"path/filepath"
	"runtime"
	"testing"
)

func expected(a, b float32) float32 {
	return float32(1 / (1 + math.Pow(float64(10), float64((b-a)/400))))
}

func elo(old, exp, score float32) float32 {
	k := float32(32) // The k-factor for Elo (default: 32)
	return float32(math.Round(float64(old + k*(score-exp))))
}

func round(v float32) float32 {
	return float32(math.Round(float64(v)*1000) / 1000)
}

func assert(tb testing.TB, condition bool, v ...interface{}) {
	if !condition {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("%s:%d: "+"\n", append([]interface{}{filepath.Base(file), line}, v...)...)
		tb.FailNow()
	}
}

func Test_elo(t *testing.T) {
	assert(t, round(expected(1613, 1609)) == 0.506)
	assert(t, round(expected(1613, 1477)) == 0.686)
	assert(t, round(expected(1613, 1388)) == 0.785)
	assert(t, round(expected(1613, 1586)) == 0.539)
	assert(t, round(expected(1613, 1720)) == 0.351)

	pairs := [][]float32{
		[]float32{0, 0},
		[]float32{1, 1},
		[]float32{10, 20},
		[]float32{123, 456},
		[]float32{2400, 2500},
	}

	for _, v := range pairs {
		assert(t, round(expected(v[0], v[1])+expected(v[1], v[0])) == 1.0)
	}

	exp := float32(0)
	exp += expected(1613, 1609)
	exp += expected(1613, 1477)
	exp += expected(1613, 1388)
	exp += expected(1613, 1586)
	exp += expected(1613, 1720)

	score := float32(0 + 0.5 + 1 + 1 + 0)

	assert(t, round(elo(1613, exp, score)) == 1601)
	assert(t, round(elo(1613, exp, 3)) == 1617)
}
