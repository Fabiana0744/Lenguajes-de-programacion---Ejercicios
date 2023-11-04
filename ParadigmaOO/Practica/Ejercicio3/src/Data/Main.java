package Data;

import java.util.Date;

public class Main {
    public static void main(String[] args) {
        // Crear una instancia de la fábrica
        ObjAgendaFactory factory = new ObjAgendaFactoryImpl();

        // Crear un contacto amigo
        Contacto contactoAmigo = factory.createContacto("Amigo Nombre", "987-654-321", "amigo");

        // Crear un evento de cumpleaños
        Evento eventoCumpleanos = factory.createEvento(new Date(), "Cumpleaños de Juan", "cumpleanos");

        // Imprimir los detalles de los objetos
        System.out.println("Detalles del Contacto Amigo:");
        contactoAmigo.imprimir();

        System.out.println("\nDetalles del Evento de Cumpleaños:");
        eventoCumpleanos.imprimir();
    }
}
