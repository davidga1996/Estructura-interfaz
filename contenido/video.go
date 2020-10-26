package contenidoWeb

import "fmt"

type Video struct {
	Titulo  string
	Formato string
	Frames  uint64
}

func (v *Video) Mostrar() {
	fmt.Println("Video")
	fmt.Println("Titulo: ", v.Titulo)
	fmt.Println("Formato: ", v.Formato)
	fmt.Println("Frames: ", v.Frames)
}
