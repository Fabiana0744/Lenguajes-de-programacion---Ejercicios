module RutaCorta

open Recipientes

// Grafo de prueba con pesos en las aristas
let grafo = [
    ("i", [["a"; "b"]; ["3"; "6"]]);
    ("a", [["i"; "c"; "d"]; ["3"; "2"; "4"]]);
    ("b", [["i"; "c"; "d"]; ["6"; "5"; "1"]]);
    ("c", [["a"; "b"; "x"]; ["2"; "5"; "3"]]);
    ("d", [["a"; "b"; "f"]; ["4"; "1"; "2"]]);
    ("f", [["d"]; ["2"]]);
    ("x", [["c"]; ["3"]])
]


// Función para generar vecinos con pesos
let rec vecinos nodo grafo =
    match grafo with
    | [] -> []
    | (head, neighbors, weights) :: rest ->
        if head = nodo then
            List.map2 (fun neighbor weight -> (neighbor, weight)) neighbors weights
        else vecinos nodo rest

// Función para extender una ruta con pesos
let extender ruta grafo = 
    let (nodo_actual, suma_pesos) = ruta
    List.filter (fun (vecino, _) -> not (miembro vecino nodo_actual))
    (List.collect (fun (vecino, peso) ->
        let nuevo_peso = (int suma_pesos) + (int peso)
        if nuevo_peso = 0 then []
        else
            ((vecino::nodo_actual), nuevo_peso.ToString()) :: []
    ) (vecinos (List.head nodo_actual) grafo))

// Función principal de búsqueda en profundidad con pesos
let rec prof2 ini fin grafo =
    let rec prof_aux fin ruta grafo =
        if List.isEmpty ruta then []
        elif List.head (fst (List.head ruta)) = fin then
            List.rev (fst (List.head ruta)) :: prof_aux fin (extender (List.head ruta) grafo) grafo
        else
            prof_aux fin (extender (List.head ruta) grafo) grafo
    prof_aux fin [([ini], "0")] grafo

// Función para encontrar el camino más corto basado en la suma de pesos
let camino_mas_corto ini fin grafo =
    let resultados = prof2 ini fin grafo
    // Filtra rutas válidas y calcula sus sumas de pesos
    let rutas_validas =
        resultados
        |> List.filter (fun ruta -> not (List.isEmpty ruta))
        |> List.map (fun ruta -> (ruta, ruta |> List.map int |> List.sum))

    if List.isEmpty rutas_validas then
        None
    else
        let ruta_mas_corta, peso_mas_corto =
            List.minBy (fun (_, peso) -> peso) rutas_validas
        Some (ruta_mas_corta, peso_mas_corto)


        
    
