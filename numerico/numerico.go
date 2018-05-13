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
func Exponente(a float64, b int) float64 {

	var pa, final float64
	final = 1
	for b >= 2 {

		dos, mul := factn(b, 2)
		pa = a
		for x := 0; x < dos; x++ {
			pa = pa * pa
		}
		final *= pa

		b = b - mul
	}
	if b == 1 {
		final *= a
	}
	return final
}

func factorprimo(numero int) [][2]int {
	i := 2
	res := make([][2]int, 0, 3)
	var fact, times int
	for i < numero {
		for numero%i == 0 {
			fact = i
			times++
			numero = numero / i
		}
		if times > 0 {
			res = append(res, [2]int{fact, times})
		}
		fact, times = 0, 0
		i++
	}
	if numero > 1 {
		res = append(res, [2]int{numero, 1})
	}
	return res
}

//Criba de eratostenes
//Si el valor del slice es true entonces el número correspondiente al indice es primo.
func findPrimos(limit int) []bool {

	res := make([]bool, limit, limit)

	//Inicializamos el array a true, por defecto está a false
	//Todos los números pares no son primos.
	for i := range res {
		if i > 2 && i%2 != 0 {
			res[i] = true
		} else {
			res[i] = false
		}
	}

	//No hace falta tratar los números pares
	for x := 3; x < limit; x = x + 2 {
		//La criba empieza en el número siguiente
		for y := x + 1; y < limit; y++ {
			if y%x == 0 {
				res[y] = false
			}
		}

	}
	return res
}
