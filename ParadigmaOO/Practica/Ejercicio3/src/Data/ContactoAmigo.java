package Data;

public class ContactoAmigo extends Contacto {
    private String email;

    public ContactoAmigo(String nombre, String telefono, String email) {
        super(nombre, telefono);
        this.email = email;
    }

    public String getEmail() {
        return email;
    }

    @Override
    public void imprimir() {
        System.out.println(this.toString());
    }

    @Override
    public String toString() {
        return "ContactoAmigo:\nNombre= " + getNombre() +
                "\nTeléfono= " + getTelefono() +
                "\nEmail= " + email;
    }
}