package numerico

//máximo comun divisor //GCD greatest commom divisor
func Mcd(a, b int) int {
	if b > 0 {
		r := a % b
		a = b
		return Mcd(a, r)
	}
	return a
}


func factn(b int, n int) (pot int, s int) {
	m := b
	s = 1
	for m >= n {
		m = int(m / n)
		if m > 0 {
			pot++
			s *= n
		}
	}
	return
}

// exponente. a elevado a b 
func Exponente(a float64, b int)  float64 {

	var pa, final float64
	final = 1
	for b >= 2 {

		dos, mul := factn(b, 2)
		pa = a
		for x:=0; x < dos; x++ {
			pa = pa*pa
		}
		final *= pa

		b = b - mul
	}
	if b == 1 {
		final *= a
	}
	return final
}

