class Evento extends ElementoAgenda {
    private String fecha;
    private String atributoAdicionalEvento;

    public Evento(String nombre, String fecha, String atributoAdicionalEvento) {
        super(nombre);
        this.fecha = fecha;
        this.atributoAdicionalEvento = atributoAdicionalEvento;
    }

    public String getFecha() {
        return fecha;
    }

    public String getAtributoAdicionalEvento() {
        return atributoAdicionalEvento;
    }

    @Override
    public String toString() {
        return "Evento: " + getNombre() + ", Fecha: " + fecha + ", Atributo Adicional Evento: " + atributoAdicionalEvento;
    }
}
