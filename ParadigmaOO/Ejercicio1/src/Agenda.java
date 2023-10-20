import java.util.ArrayList;
import java.util.List;

class Agenda {
    private List<ElementoAgenda> elementos;

    public Agenda() {
        elementos = new ArrayList<>();
    }

    public void agregarElemento(ElementoAgenda elemento) {
        elementos.add(elemento);
    }

    public void eliminarElemento(ElementoAgenda elemento) {
        elementos.remove(elemento);
    }

    public void modificarElemento(ElementoAgenda elementoViejo, ElementoAgenda elementoNuevo) {
        int index = elementos.indexOf(elementoViejo);
        if (index != -1) {
            elementos.set(index, elementoNuevo);
        }
    }

    public void mostrarElementos() {
        for (ElementoAgenda elemento : elementos) {
            System.out.println(elemento.toString());
        }
    }
}
