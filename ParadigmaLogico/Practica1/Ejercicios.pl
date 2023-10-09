%--------------- Ejercicio 1 ------------------
sumlist([], 0).
sumlist([X|Xs], S) :-
    sumlist(Xs, S1), S is S1 + X.
% sumlist([1, 2, 3, 4], S).


%--------------- Ejercicio 2 ------------------
subconj([], []).
subconj(_, []).
subconj(S, [X|Xs]) :-
    member(X, S), subconj(S, Xs).
% subconj([1, 2, 3], [1, 2, 3, 4, 5]).


%--------------- Ejercicio 3 ------------------
aplanar([], []).
aplanar([X|Xs], L2) :-
    aplanar(X, Y), aplanar(Xs, Ys), append(Y, Ys, L2).
aplanar(X, [X]) :-
    \+ is_list(X).
aplanar([], []) :-
    \+ is_list([]).
% aplanar([1,[2, [3,4],5],[6,[7,8]]], L).


%--------------- Ejercicio 4 ------------------
partir([], _, [], []).
partir([X|Xs], Umbral, [X|Menores], Mayores) :-
    X =< Umbral, partir(Xs, Umbral, Menores, Mayores).
partir([X|Xs], Umbral, Menores, [X|Mayores]) :-
    X > Umbral, partir(Xs, Umbral, Menores, Mayores).
% partir([2, 7, 4, 8, 9, 1], 6, Menores, Mayores).


%--------------- Ejercicio 5 ------------------

% Forma1: Contando las silabas y palabras

contieneSubcadena(Cadena, Subcadena) :-
    sub_string(Cadena, _, _, _, Subcadena).

filtrarCadenasConSubcadena(Subcadena, ListaEntrada, Filtradas) :-
    findall(Cadena, (member(Cadena, ListaEntrada), contieneSubcadena(Cadena, Subcadena)), Filtradas).
% filtrarCadenasConSubcadena("la", ["la casa", "el perrola", "pintando la cerca"], Filtradas).


% Forma2: Contando solo las palabras

estaEnLista(Palabra, ListaPalabras) :-
    member(Palabra, ListaPalabras).

dividirEnPalabras(Cadena, ListaPalabras) :-
    split_string(Cadena, " ", " ", ListaPalabras).

filtrarCadenasConPalabraCompleta(Subcadena, ListaEntrada, Filtradas) :-
    findall(Cadena, (
        member(Cadena, ListaEntrada),
        dividirEnPalabras(Cadena, ListaPalabras),
        estaEnLista(Subcadena, ListaPalabras)
    ), Filtradas).

% filtrarCadenasConPalabraCompleta("la", ["la casa", "el perrola", "pintando la cerca"], Filtradas).


