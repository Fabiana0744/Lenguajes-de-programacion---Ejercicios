package Data;

import java.util.LinkedList;

public class AgendaSingleton {
    private static AgendaSingleton instance = null;
    private LinkedList<Object> listaObjetos;
    private LinkedList<String> contactos;
    private LinkedList<String> eventos;

    private AgendaSingleton() {
        listaObjetos = new LinkedList<>();
        contactos = new LinkedList<>();
        eventos = new LinkedList<>();
    }

    public static AgendaSingleton getInstance() {
        if (instance == null) {
            synchronized (AgendaSingleton.class) {
                if (instance == null) {
                    instance = new AgendaSingleton();
                }
            }
        }
        return instance;
    }

    public LinkedList<String> getContactos() {
        return contactos;
    }

    public LinkedList<String> getEventos() {
        return eventos;
    }

    public LinkedList<Object> getListaObjetos() {
        return listaObjetos;
    }
}