//Generador lineal congruencial
package random

func Glcgen(seed uint32) func() uint32 {
	var a, c, m uint32
	a = 13 //9, 5 tambien funcionan....
	c = 3
	m = 16
	return func() uint32 {
		seed = (a*seed + c) % m
		return seed
	}
}
