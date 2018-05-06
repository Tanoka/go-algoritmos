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
	mc := Mcd(fl.a, fl.b)

	if mc != fl.re {
		t.Errorf("Expected %d | Actual %d ",fl.re, mc)
	}
}

}

func TestExponente(t *testing.T) {
	flagtest := []struct{
		a float64
		b int
		re float64
	}{
			{3,1,3},
			{7,3, 343},
			{4,4,256},
			{3,7,2187},
			{7,6,117649},
			{3,40,12157665459056928801},
	}

	var pot float64
	for _, fl := range flagtest {
		pot = Exponente(fl.a, fl.b)
		if (pot != fl.re) {
			t.Errorf("Values:%f, %d | Expoted:%f, Actual:%f",fl.a, fl.b, fl.re, pot)
		}
	}
}

func TestFactn(t *testing.T) {
	flagtest := []struct {
		b   int
		re1 int
		re2 int
	}{
		{1, 0, 1},
		{3, 1, 2},
		{4, 2, 4},
		{7, 2, 4},
		{8, 3, 8},
		{11, 3, 8},
		{15, 3, 8},
		{16, 4, 16},
		{17, 4, 16},
	}

	var pot, val int
	for _, fl := range flagtest {
		pot, val = factn(fl.b, 2)
		if pot != fl.re1 {
			t.Errorf("Error exponente. Values:%d | Expoted:%d, Actual:%d", fl.b, fl.re1, pot)
		}
		if val != fl.re2 {
			t.Errorf("Error total. Values:%d | Expoted:%d, Actual:%d", fl.b, fl.re2, val)
		}
	}
}
