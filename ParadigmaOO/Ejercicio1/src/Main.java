public class Main {
    public static void main(String[] args) {
        Agenda agenda1 = new Agenda();
        Agenda agenda2 = new Agenda();

        ContactoT1 contacto1 = new ContactoT1("Juan", "123 Main St", "123-456-7890", "Atributo1");
        ContactoT2 contacto2 = new ContactoT2("Maria", "456 Elm St", "987-654-3210", "Atributo2");
        Evento evento1 = new Evento("Fiesta de Cumpleaños", "2023-10-20", "AtributoEvento1");
        Evento evento2 = new Evento("Conferencia", "2023-11-15", "AtributoEvento2");

        agenda1.agregarElemento(contacto1);
        agenda1.agregarElemento(evento1);
        agenda2.agregarElemento(contacto2);
        agenda2.agregarElemento(evento2);

        agenda1.mostrarElementos();
        System.out.println();
        agenda2.mostrarElementos();
    }
}
