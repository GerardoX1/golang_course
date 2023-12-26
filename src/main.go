package main

import "fmt"

func main() {
	// Declaracion de variables
	helloMessage := "Hello"
	wordMessage := "word"

	// Printnl
	fmt.Println(helloMessage, wordMessage)
	fmt.Println(helloMessage, wordMessage)

	// Printf
	nombre := "Luis"
	edad := 20
	hijos := true

	fmt.Printf("%s tiene mas de %d años \n", nombre, edad)
	fmt.Printf("%#v tiene mas de %#v años \n", nombre, edad)

	//Sprintf
	message := fmt.Sprintf("%v tiene mas de %v años", nombre, edad)

	fmt.Println(message)

	// Tipo de datos
	fmt.Printf("helloMessage: %T \n", helloMessage)
	fmt.Printf("hijos: %T \n", hijos)
	fmt.Printf("edad: %T \n", edad)
}
