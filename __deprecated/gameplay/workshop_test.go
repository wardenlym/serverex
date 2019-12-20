package gameplay

import (
	"fmt"
	"testing"
)

func Test_lot(t *testing.T) {
	countMap := map[int]int{}

	check := 50000
	for i := 0; i < check; i++ {
		v := lot.Lots(up_prob...)
		countMap[v]++
	}

	for i, count := range countMap {
		result := float64(count) / float64(check) * 100
		prob := float64(up_prob[i].Prob())

		name := fmt.Sprint(up_prob[i].(upgrade_prob).star)

		// 0.1 check
		if (prob-1) <= result && result < (prob+1) {
			fmt.Printf("ok %3.5f%%(%7d) : star%s\n", result, count, name)
		} else {
			t.Errorf("error %3.5f%%(%7d) : star%s\n", result, count, name)
		}
	}
}
