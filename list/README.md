## Listas

Lista doble de números enteros.


Hay dos ficheros de tests, uno con acceso privado, package con el mismo nombre que el package a testear y otro como si fuera acceso público, importando el package a testear.


Para lanzar los tests:
```
docker run --rm -v $PWD:/go/src/github.com/tanoka/go-algoritmos/list -w /go/src/github.com/tanoka/go-algoritmos/list golang go test -v
```

