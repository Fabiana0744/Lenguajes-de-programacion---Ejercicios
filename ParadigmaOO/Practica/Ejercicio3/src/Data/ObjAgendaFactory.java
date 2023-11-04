package Data;

import java.util.Date;

public interface ObjAgendaFactory {
    Contacto createContacto(String nombre, String telefono, String tipo);
    Evento createEvento(Date fecha, String nombre, String tipo);
}
