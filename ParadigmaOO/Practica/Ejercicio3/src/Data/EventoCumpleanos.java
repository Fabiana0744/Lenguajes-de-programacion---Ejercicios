package Data;

import java.util.Date;

public class EventoCumpleanos extends Evento {
    private int edad;

    public EventoCumpleanos(Date fecha, String nombre, int edad) {
        super(fecha, nombre);
        this.edad = edad;
    }

    public int getEdad() {
        return edad;
    }

    public void setEdad(int edad) {
        this.edad = edad;
    }

    @Override
    public void imprimir() {
        System.out.println(this.toString());
    }

    @Override
    public String toString() {
        return "EventoCumpleanos:\nNombre= " + getNombre() +
                "\nFecha= " + getFecha() +
                "\nEdad= " + edad;
    }
}