package main

import (
	"fmt"

	contenido "./contenido"
)

func crearImagen() contenido.Imagen {
	var titulo string
	var formato string
	var canales string

	fmt.Println("Ingrese el título: ")
	fmt.Scan(&titulo)
	fmt.Println("Ingrese el formato: ")
	fmt.Scan(&formato)
	fmt.Println("Ingrese los canales: ")
	fmt.Scan(&canales)
	return contenido.Imagen{
		Titulo:  titulo,
		Formato: formato,
		Canales: canales}
}

func crearAudio() contenido.Audio {
	var titulo string
	var formato string
	var duracion uint64

	fmt.Println("Ingrese el título: ")
	fmt.Scan(&titulo)
	fmt.Println("Ingrese el formato: ")
	fmt.Scan(&formato)
	fmt.Println("Ingrese la duracion: ")
	fmt.Scan(&duracion)
	return contenido.Audio{
		Titulo:   titulo,
		Formato:  formato,
		Duracion: duracion}
}

func crearVideo() contenido.Video {
	var titulo string
	var formato string
	var frames uint64

	fmt.Println("Ingrese el título: ")
	fmt.Scan(&titulo)
	fmt.Println("Ingrese el formato: ")
	fmt.Scan(&formato)
	fmt.Println("Ingrese la duracion: ")
	fmt.Scan(&frames)
	return contenido.Video{
		Titulo:  titulo,
		Formato: formato,
		Frames:  frames}
}

func main() {
	miContenidoWeb := contenido.ContenidoWeb{}
	var option string
	noSalir := true
	for noSalir {
		fmt.Println("Menu")
		fmt.Println("Elija una opción: ")
		fmt.Println("1) Crear Imagen")
		fmt.Println("2) Crear Audio")
		fmt.Println("3) Crear Video")
		fmt.Println("4) Mostrar todo")
		fmt.Println("5) Salir")
		fmt.Print(">>> ")
		fmt.Scan(&option)
		switch option {
		case "1":
			imagen := crearImagen()
			miContenidoWeb.Insertar(&imagen)
		case "2":
			audio := crearAudio()
			miContenidoWeb.Insertar(&audio)
		case "3":
			video := crearVideo()
			miContenidoWeb.Insertar(&video)
		case "4":
			miContenidoWeb.Mostrar()
		case "5":
			noSalir = false
		default:
			fmt.Println("Esa opción no existe")
		}
	}
}
