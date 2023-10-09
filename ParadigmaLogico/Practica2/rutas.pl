% Definición de conexiones
conectado(i, a, 15).
conectado(i, b, 20).
conectado(a, b, 45).
conectado(a, c, 20).
conectado(b, f, 7).
conectado(b, c, 5).
conectado(c, f, 11).


% Predicado para encontrar la ruta más corta desde Ini hasta Fin
ruta_mas_corta(Ini, Fin, Ruta, Peso) :-
    ruta2(Ini, Fin, [], Ruta, 0, Peso).

% Predicado para encontrar la ruta más corta con pesos
ruta2(Fin, Fin, _, [Fin], Peso, Peso).

ruta2(Inicio, Fin, Visitados, [Inicio | Resto], PesoActual, PesoTotal) :-
    conectado(Inicio, Alguien, Peso),
    not(member(Alguien, Visitados)),
    NuevoPesoActual is PesoActual + Peso,
    ruta2(Alguien, Fin, [Inicio | Visitados], Resto, NuevoPesoActual, PesoTotal).

% Predicado para encontrar la ruta más corta e imprimir el resultado
encontrar_ruta_mas_corta(Inicio, Fin, Ruta, Peso) :-
    findall([R, P], ruta_mas_corta(Inicio, Fin, R, P), Rutas),
    min_peso(Rutas, [Ruta, Peso]).

min_peso([[Ruta, Peso]], [Ruta, Peso]).
min_peso([[_, P1] | Resto], [R2, P2]) :-
    min_peso(Resto, [R2, P2]),
    P1 >= P2.
min_peso([[R1, P1] | Resto], [R1, P1]) :-
    min_peso(Resto, [_, P2]),
    P1 < P2.

% Prueba
% encontrar_ruta_mas_corta(i, f, Ruta, Peso).
