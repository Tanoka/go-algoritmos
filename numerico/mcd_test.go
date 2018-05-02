package numerico

import "testing"

func TestMcd(t *testing.T) {

	flagtest := []struct{
		a int
		b int
		re int
	}{
		{34,7,1},
		{54,12,6},
		{56,42,14},
		{86329, 131, 131},
		{13,13,13},
		{273, 182,91},

	}

	for _,fl := range flagtest {
	mc := mcd(fl.a, fl.b)

	if mc != fl.re {
		t.Errorf("Expected %d | Actual %d ",fl.re, mc)
	}
}

}
