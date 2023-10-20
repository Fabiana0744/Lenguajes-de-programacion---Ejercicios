class ContactoT1 extends Contacto {
    private String atributoAdicionalT1;

    public ContactoT1(String nombre, String direccion, String telefono, String atributoAdicionalT1) {
        super(nombre, direccion, telefono);
        this.atributoAdicionalT1 = atributoAdicionalT1;
    }

    public String getAtributoAdicionalT1() {
        return atributoAdicionalT1;
    }

    @Override
    public String toString() {
        return "Conracto 1: " + getNombre() + ", Dirección: " + getDireccion() + ", Teléfono: " + getTelefono() + ", Atributo Adicional T1: " + atributoAdicionalT1;
    }
}
