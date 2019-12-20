package util

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func RandomInMinMax(min, max int) int {
	return rand.Intn(max-min) + min
}

func RandomIn(a, b int) int {
	var min, max int
	if a < b {
		min = a
		max = b
	} else {
		min = b
		max = a
	}
	return RandomInMinMax(min, max)
}

func RandomIndexList(max, count int) []int {
	m := make([]int, count)
	// In the following loop, the iteration when i=0 always swaps m[0] with m[0].
	// A change to remove this useless iteration is to assign 1 to i in the init
	// statement. But Perm also effects r. Making this change will affect
	// the final state of r. So this change can't be made for compatibility
	// reasons for Go 1.
	for i := 0; i < count; i++ {
		v := rand.Intn(max)
		m[i] = v
	}
	return m
}

func HitRate(rate float64) bool {
	k := rand.Float64()
	if rate < k {
		return true
	}
	return false
}
