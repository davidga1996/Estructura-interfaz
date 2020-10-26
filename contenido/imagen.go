package contenidoWeb

import "fmt"

type Imagen struct {
	Titulo  string
	Formato string
	Canales string
}

func (i *Imagen) Mostrar() {
	fmt.Println("Imagen")
	fmt.Println("Titulo: ", i.Titulo)
	fmt.Println("Formato: ", i.Formato)
	fmt.Println("Canales: ", i.Canales)
}
