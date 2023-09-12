let n_esimo n lista =
    let resultado = List.map2 (fun x i -> if i = n then Some x else None) lista [0..(List.length lista - 1)]
    match List.tryFind (fun x -> Option.isSome x) resultado with
    | Some(Some(value)) -> value
    | _ -> -1

// Ejemplo de uso
let lista1 = [1; 2; 3; 4; 5]
let resultado1 = n_esimo 2 lista1 
let resultado2 = n_esimo 3 lista1 
let resultado3 = n_esimo 6 lista1 

printfn "Resultado 1: %A" resultado1
printfn "Resultado 2: %A" resultado2
printfn "Resultado 3: %A" resultado3
