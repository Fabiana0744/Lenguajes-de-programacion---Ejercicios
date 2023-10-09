persona('Persona 1', [0, 1, 1, 0, 0, 1, 1, 0, 0, 1]).
persona('Persona 2', [1, 1, 1, 1, 0, 0, 1, 0, 1, 1]).
persona('Persona 3', [0, 1, 0, 1, 1, 0, 0, 1, 1, 1]).

% Porcentaje de similitud entre dos cromosomas
porcentaje_similitud([], [], 0).
porcentaje_similitud([Gen1|Resto1], [Gen2|Resto2], Similitud) :-
    porcentaje_similitud(Resto1, Resto2, SimilitudResto),
    (Gen1 = Gen2 -> Similitud is SimilitudResto + 1 ; Similitud is SimilitudResto).

% Predicado para encontrar la persona más similar a una muestra
encontrar_persona_mas_parecida(Muestra) :-
    findall((Nombre, Porcentaje), (
        persona(Nombre, CromosomaCandidato),
        porcentaje_similitud(Muestra, CromosomaCandidato, Similitud),
        length(Muestra, LongitudMuestra),
        Porcentaje is (Similitud / LongitudMuestra) * 100
    ), Resultados),
    max_member((NombreMax, PorcentajeMax), Resultados),
    format("Persona más similar: ~w con ~2f%% de similitud~n", [NombreMax, PorcentajeMax]).


% Prueba:
% encontrar_persona_mas_parecida([0, 1, 0, 1, 0, 1, 0, 1, 0, 1]).
