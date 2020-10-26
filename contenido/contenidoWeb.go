package contenidoWeb

import "fmt"

type ContenidoWeb struct {
	multimedias []Multimedia
}

func (cw *ContenidoWeb) Insertar(multi Multimedia) {
	cw.multimedias = append(cw.multimedias, multi)
}

func (cw *ContenidoWeb) Mostrar() {
	if len(cw.multimedias) == 0 {
		fmt.Println("Está vacío...")
	}
	for i, multimedia := range cw.multimedias {
		fmt.Println("index: ", i)
		multimedia.Mostrar()
	}
}
