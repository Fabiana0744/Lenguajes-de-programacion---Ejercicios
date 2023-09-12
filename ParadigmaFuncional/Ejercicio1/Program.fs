let desplazarIzquierda lista n =
    let rec desplazarIzqAux lista n =
        match n with
        | 0 -> lista
        | _ ->
            let nuevaLista = List.tail lista @ [0]
            desplazarIzqAux nuevaLista (n - 1)
    desplazarIzqAux lista n

let desplazarDerecha lista n =
    let longitud = List.length lista
    let rec desplazarDerAux lista n =
        match n with
        | 0 -> lista
        | _ ->
            let nuevaLista = 0 :: List.init (longitud - 1) (fun i -> List.item i lista)
            desplazarDerAux nuevaLista (n - 1)
    desplazarDerAux lista n

let desplazar direccion n lista =
    let longitud = List.length lista

    if longitud = 0 then []
    elif direccion = "izq" then
        let desplazada = desplazarIzquierda lista n
        List.take longitud desplazada
    elif direccion = "der" then
        let desplazada = desplazarDerecha lista n
        List.take longitud desplazada
    else []

// Ejemplos de uso:
let lista1 = [1; 2; 3; 4; 5]
let resultado1 = desplazar "izq" 3 lista1
let resultado2 = desplazar "der" 2 lista1
let resultado3 = desplazar "izq" 6 lista1
let resultado4 = desplazar "der" 3 lista1
let resultado5 = desplazar "der" 4 lista1
let resultado6 = desplazar "izq" 4 lista1

printfn "Resultado 1: %A" resultado1
printfn "Resultado 2: %A" resultado2
printfn "Resultado 3: %A" resultado3
printfn "Resultado 4: %A" resultado4
printfn "Resultado 5: %A" resultado5
printfn "Resultado 6: %A" resultado6