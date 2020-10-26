package contenidoWeb

import "fmt"

type Audio struct {
	Titulo   string
	Formato  string
	Duracion uint64
}

func (a *Audio) Mostrar() {
	fmt.Println("Audio")
	fmt.Println("Titulo: ", a.Titulo)
	fmt.Println("Formato: ", a.Formato)
	fmt.Println("Duracion: ", a.Duracion)
}
