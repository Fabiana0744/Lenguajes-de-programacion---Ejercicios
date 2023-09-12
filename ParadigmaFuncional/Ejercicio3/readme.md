# Función para Obtener el N-ésimo Elemento de una Lista en F#

Esta es una función simple en F# que te permite obtener el N-ésimo elemento de una lista sin utilizar recursión, sino haciendo uso de la función `List.map2`.

## Descripción

La función `n_esimo` toma dos argumentos: `n` y `lista`. `n` es el índice del elemento que deseas obtener (comenzando en 0 para el primer elemento), y `lista` es la lista de la cual deseas extraer el elemento.

La función utiliza la función `List.map2` para crear una nueva lista que consiste en pares de elementos, donde el primer elemento de cada par es el elemento de la lista original y el segundo elemento es su índice correspondiente. Luego, filtra esta lista para obtener solo los pares cuyo índice coincide con el valor de `n`, y finalmente devuelve el primer elemento de ese par, que es el elemento deseado.
