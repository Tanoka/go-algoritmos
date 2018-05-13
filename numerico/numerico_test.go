package numerico

import "testing"

func TestMcd(t *testing.T) {

	flagtest := []struct {
		a  int
		b  int
		re int
	}{
		{34, 7, 1},
		{54, 12, 6},
		{56, 42, 14},
		{86329, 131, 131},
		{13, 13, 13},
		{273, 182, 91},
	}

	for _, fl := range flagtest {
		mc := Mcd(fl.a, fl.b)

		if mc != fl.re {
			t.Errorf("Expected %d | Actual %d ", fl.re, mc)
		}
	}

}

func TestExponente(t *testing.T) {
	flagtest := []struct {
		a  float64
		b  int
		re float64
	}{
		{3, 1, 3},
		{7, 3, 343},
		{4, 4, 256},
		{3, 7, 2187},
		{7, 6, 117649},
		{3, 40, 12157665459056928801},
	}

	var pot float64
	for _, fl := range flagtest {
		pot = Exponente(fl.a, fl.b)
		if pot != fl.re {
			t.Errorf("Values:%f, %d | Expoted:%f, Actual:%f", fl.a, fl.b, fl.re, pot)
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

type resFaPrimo [][2]int

func TestFactorPrimo(t *testing.T) {
	flagtest := []struct {
		val int
		res resFaPrimo
	}{
		{
			10,
			resFaPrimo{{2, 1}, {5, 1}},
		},
		{
			11,
			resFaPrimo{{11, 1}},
		},
		{
			123,
			resFaPrimo{{3, 1}, {41, 1}},
		},
		{
			256,
			resFaPrimo{{2, 8}},
		},
		{
			234,
			resFaPrimo{{2, 1}, {3, 2}, {13, 1}},
		},
	}
	var sol resFaPrimo
	for _, fp := range flagtest {
		sol = factorprimo(fp.val)
		for pos, r1 := range fp.res {
			if r1[0] != sol[pos][0] || r1[1] != sol[pos][1] {
				t.Errorf("Error, Values:%d | Expected:%v, Actual:%v", fp.val, fp.res, sol)
			}
		}
	}
}

func TestFindPrimos(t *testing.T) {

	firstprim := []bool{true, true, true, true, false, true, false, true, false, false, false}

	flagtest := []struct {
		val int
		res []bool
	}{
		{val: 10, res: firstprim},
		{val: 16, res: append(firstprim, []bool{true, false, true, false, false}...)},
	}

	var res []bool
	for _, fl := range flagtest {
		res = findPrimos(fl.val)
		for i, so := range fl.res {
			if len(res) < i && so != res[i] {
				t.Errorf("Error. value:%d | Expected:%v Actual:%v", fl.val, fl.res, res)
			}
		}
	}
}
