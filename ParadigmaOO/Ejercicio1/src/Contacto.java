abstract class Contacto extends ElementoAgenda {
    private String direccion;
    private String telefono;

    public Contacto(String nombre, String direccion, String telefono) {
        super(nombre);
        this.direccion = direccion;
        this.telefono = telefono;
    }

    public String getDireccion() {
        return direccion;
    }

    public String getTelefono() {
        return telefono;
    }

    @Override
    public abstract String toString();
}
