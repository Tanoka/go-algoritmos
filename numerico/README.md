## Algoritmos cálculo

### Máximo común divisor
```go
func Mcd(a, b int) int 
```

Utilizando el algoritmo de Euclides

https://es.wikipedia.org/wiki/M%C3%A1ximo_com%C3%BAn_divisor


### Exponenciación
```go
func Exponente(a float64, b int)  float64 
```
En lugar de A^n = A * A * A * A * A ... O(N)

Utilizamos las propiedades los exponentes O(log N):
```
A⁶ = A⁽²*³⁾ = (A²)³
A³ = A⁽²⁺¹⁾ = A² * A¹
```

Si N =1000.000 O(N) es 1000.000 pero O(log N) es aproximadamente 20 
