## Generador lineal congruencial

`xₙ := (a*xₙ₋₁ + c) % m`

los valores a,c y m deben ser unos concretos, sino la secuencia se puede repetir muy pronto, por ejemplo a=7, c=3, m=10 sacan la secuencia 0,3,4,1 y vuelve a empezar por el 0.
Si m^2, c impar y a=4c+1 debe funcionar con una secuenca de m números antes de volver a repetirse (full period)
https://es.wikipedia.org/wiki/Generador_lineal_congruencial explica cuales son los números GLC más comunes, quien los usa, etc
	a = 1103515245 
	ci = 12345
	m  =2147483647 = 2^31-1

