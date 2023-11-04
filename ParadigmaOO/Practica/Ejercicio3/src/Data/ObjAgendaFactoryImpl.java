package Data;

import java.util.Date;

public class ObjAgendaFactoryImpl implements ObjAgendaFactory {
    @Override
    public Contacto createContacto(String nombre, String telefono, String tipo) {
        if ("familiar".equals(tipo)) {
            return new ContactoFamiliar(nombre, telefono, "");
        } else if ("amigo".equals(tipo)) {
            return new ContactoAmigo(nombre, telefono, "");
        } else {
            throw new IllegalArgumentException("Tipo de contacto no válido: " + tipo);
        }
    }

    @Override
    public Evento createEvento(Date fecha, String nombre, String tipo) {
        if ("reunion".equals(tipo)) {
            return new EventoReunion(0, fecha, nombre);
        } else if ("cumpleanos".equals(tipo)) {
            return new EventoCumpleanos(fecha, nombre, 18);
        } else {
            throw new IllegalArgumentException("Tipo de evento no válido: " + tipo);
        }
    }
}
