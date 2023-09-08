let esPalabraCompleta (subcadena: string) (cadena: string) =
    let palabras = cadena.Split([|' '|])
    Array.exists (fun palabra -> palabra = subcadena) palabras

let subCadenas (subcadena: string) (lista: string list) =
    lista |> List.filter (esPalabraCompleta subcadena)

let listaDeCadenas = ["la casa"; "el pelaje"; "pintando la cerca"]
let subcadenaBuscada = "la"

let resultado = subCadenas subcadenaBuscada listaDeCadenas
printfn "%A" resultado
