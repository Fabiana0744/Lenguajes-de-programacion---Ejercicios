class ContactoT2 extends Contacto {
    private String atributoAdicionalT2;

    public ContactoT2(String nombre, String direccion, String telefono, String atributoAdicionalT2) {
        super(nombre, direccion, telefono);
        this.atributoAdicionalT2 = atributoAdicionalT2;
    }

    public String getAtributoAdicionalT2() {
        return atributoAdicionalT2;
    }

    @Override
    public String toString() {
        return "Tipo 2: " + getNombre() + ", Dirección: " + getDireccion() + ", Teléfono: " + getTelefono() + ", Atributo Adicional T2: " + atributoAdicionalT2;
    }
}
