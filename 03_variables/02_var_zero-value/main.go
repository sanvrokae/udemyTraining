package main

import "fmt"

func main() {

	//Declarando sin asignar valor
	var a int     //Imprime un entero sin declarar = 0
	var b string  //Imprime una cadena sin declarar = ''
	var c float64 //Imprime un flotante sin declarar = 0
	var d bool    //Imprime un booleano no declarado = falso

	fmt.Printf("%v \n", a)
	fmt.Printf("%v \n", b)
	fmt.Printf("%v \n", c)
	fmt.Printf("%v \n", d)

	fmt.Println()
}
