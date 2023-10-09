# Ejercicios de Prolog


## Ejercicio 1: Suma de Elementos en una Lista

### Descripción
El ejercicio consiste en definir un predicado `sumlist` que tome una lista de números y calcule la suma de sus elem>

### Uso
?- sumlist([1, 2, 3, 4], S).

## Ejercicio 2: Subconjuntos de una Lista

### Descripción
El ejercicio consiste en definir un predicado `subconj` que determine si una lista es un subconjunto de otra lista.

### Uso
?- subconj([1, 2], [1, 2, 3, 4]).

## Ejercicio 3: Aplanar Listas Anidadas

### Descripción
El ejercicio consiste en definir un predicado `aplanar` que tome una lista con múltiples listas anidadas como eleme>

### Uso
?- aplanar([1, [2, [3, 4], 5], [6, [7, 8]]], L).

## Ejercicio 4: Dividir una Lista por Umbral

### Descripción
El ejercicio consiste en definir un predicado `partir` que divida una lista en dos listas, una con elementos menore>

### Uso
?- partir([2, 7, 4, 8, 9, 1], 6, Menores, Mayores).

## Ejercicio 5: Filtrar Cadenas con Subcadena

### Descripción
El ejercicio consiste en definir un predicado `sub_cadenas` que filtre cadenas que contengan una subcadena especifi>

### Uso
?- sub_cadenas("la", ["la casa", "el perro", "pintando la cerca"], Filtradas).
