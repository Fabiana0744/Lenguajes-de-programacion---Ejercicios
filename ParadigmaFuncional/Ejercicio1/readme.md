# Ejercicio de Desplazamiento de Listas en F#

Este ejercicio involucra la creación de funciones en F# para desplazar elementos en una lista hacia la izquierda o la derecha, llenando los espacios vacíos con ceros según se indique por el usuario.

## Descripción

Este ejercicio consta de tres funciones principales:

`desplazarIzquierda lista n`: Esta función toma una lista `lista` y un número entero `n` como argumentos y desplaza los elementos de la lista hacia la izquierda `n` posiciones. Los espacios vacíos se llenan con ceros.

`desplazarDerecha lista n`: Esta función toma una lista `lista` y un número entero `n` como argumentos y desplaza los elementos de la lista hacia la derecha `n` posiciones. Los espacios vacíos se llenan con ceros.

`desplazar direccion n lista`: Esta función principal toma tres argumentos: `direccion` (una cadena que puede ser "izq" o "der" para indicar la dirección del desplazamiento), `n` (el número de posiciones a desplazar) y `lista` (la lista de elementos). Dependiendo de la dirección, utiliza las funciones `desplazarIzquierda` o `desplazarDerecha` para realizar el desplazamiento y devuelve la lista resultante.
