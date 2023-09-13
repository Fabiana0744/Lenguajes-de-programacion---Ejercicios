let laberinto = [
    (1, [2; 7]); (2, [1; 3; 8]); (3, [2; 4; 9]);
    (4, [3; 10]); (5, [6; 11]); (6, [5; 12]);
    (7, [1; 13]); (8, [2; 9; 14]); (9, [3; 8; 15]);
    (10, [4; 16]); (11, [5; 17]); (12, [6; 18]);
    (13, [7; 14; 19]); (14, [8; 13; 15; 20]); (15, [9; 14; 21]);
    (16, [10; 22]); (17, [11; 23]); (18, [12; 24]);
    (19, [13; 25]); (20, [14; 26]); (21, [15; 22]);
    (22, [16; 21; 28]); (23, [17; 29]); (24, [18; 30]);
    (25, [19; 31]); (26, [20; 27]); (27, [26; 28]);
    (28, [22; 27; 29; 34]); (29, [23; 28; 35]); (30, [24; 36]);
    (31, [25; 32]); (32, [31; 33]); (33, [32; 34]);
    (34, [28; 33; 35; 36]); (35, [29; 34; 36]); (36, [30; 35])
]

let laberintoConParedesModificadas = [
    (1, [2; 7]); (2, [1; 3; 8]); (3, [2; 4; 9]);
    (4, [3; 10]); (5, [6; 11]); (6, [5; 12]);
    (7, [1; 13]); (8, [2; 9; 14]); (9, [3; 8; 15]);
    (10, [4; 16]); (11, [5; 17]); (12, [6; 18]);
    (13, [7; 14; 19; 20]); (14, [8; 13; 15; 21]); (15, [9; 14; 22]);
    (16, [10; 23]); (17, [11; 24]); (18, [12; 25]);
    (19, [13; 26]); (20, [14; 27]); (21, [15; 22]);
    (22, [21; 28]); (23, [16; 29]); (24, [17; 30]);
    (25, [18; 31]); (26, [19; 32]); (27, [20; 33]);
    (28, [22; 27; 34]); (29, [23; 35]); (30, [24; 36]);
    (31, [25; 32]); (32, [31; 33]); (33, [32; 34]);
    (34, [28; 33; 35; 36]); (35, [29; 34]); (36, [30; 35])
]


// Función para verificar si un elemento está en una lista
let miembro elem lista =
    List.exists (fun x -> x = elem) lista

// Función para generar vecinos
let rec vecinos nodo grafo =
    match grafo with
    | [] -> []
    | (head, neighbors) :: rest ->
        if head = nodo then neighbors
        else vecinos nodo rest

// Función para extender una ruta
let extender ruta grafo = 
    List.filter
        (fun x -> x <> [])
        (List.map  (fun x -> if (miembro x ruta) then [] else x::ruta) (vecinos (List.head ruta) grafo)) 

// Función principal de búsqueda en profundidad
let rec prof2 ini fin grafo =
    let rec prof_aux fin ruta grafo =
        if List.isEmpty ruta then []
        elif (List.head (List.head ruta)) = fin then
            List.rev (List.head ruta) //:: prof_aux fin ((extender (List.head ruta) grafo) @ (List.tail ruta)) grafo       
        else
            prof_aux fin ((List.tail ruta) @ (extender (List.head ruta) grafo)  ) grafo
    prof_aux fin [[ini]] grafo


// Función principal de búsqueda en profundidad
let obtenerTodasRutas ini fin grafo =
    let rec prof_aux fin ruta grafo =
        if List.isEmpty ruta then []
        elif (List.head (List.head ruta)) = fin then
            [List.rev (List.head ruta)] @ (prof_aux fin ((extender (List.head ruta) grafo) @ (List.tail ruta)) grafo)       
        else
            prof_aux fin ((List.tail ruta) @ (extender (List.head ruta) grafo)) grafo
    prof_aux fin [[ini]] grafo

let rutas = obtenerTodasRutas 2 32 laberinto  
printfn "Todas las rutas posibles son:"
List.iter (fun path -> printfn "%A" path) rutas

let ruta1 = prof2 2 32 laberinto
printfn "\nLa ruta más corta es:"
printfn "%A" ruta1


//------------------OTRA FORMA DE HACERLO-------------------------
(*
// Función para obtener todas las rutas posibles
let imprimirRutas ini fin grafo =
    let rec dfs ruta rutas =
        let node = List.head ruta
        if node = fin then
            List.rev ruta :: rutas
        else
            let neighbors = vecinos node grafo |> List.filter (fun vecino -> not (miembro vecino ruta))
            List.fold (fun acc vecino ->
                dfs (vecino :: ruta) acc
            ) rutas neighbors
    in
    dfs [ini] []

let rutasPosibles = imprimirRutas 2 32 laberinto
printfn "Todas las rutas posible son:"
List.iter (fun path -> printfn "%A" path) rutasPosibles

let ruta = prof2 2 32 laberinto
printfn "\nLa ruta más corta es:"
printfn "%A" ruta
*)
