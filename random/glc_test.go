package random

import "testing"
import "sort"

func TestGlcgen(t *testing.T) {
	val := make([]int, 15)
	glc := Glcgen(0)
	for x := 0; x < 15; x++ {
		val[x] = int(glc())
	}
	sort.Ints(val)
	an := -1
	for _, v := range val {
		if v <= an {
			t.Errorf("%d is lesser than %d", v, an)
		}
		an = v
	}
}
