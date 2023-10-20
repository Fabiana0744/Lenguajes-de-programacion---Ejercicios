abstract class ElementoAgenda {
    private String nombre;

    public ElementoAgenda(String nombre) {
        this.nombre = nombre;
    }

    public String getNombre() {
        return nombre;
    }

    public abstract String toString();
}
