package main

import "fmt"

func main() {
	x:= 0
	increment := func() int { //funcion anonima asignada a incrementar
				  // func expression = asignar una funcion a una variable
		x++
		return x
	}
	fmt.Println(increment())
	fmt.Println(increment())
}