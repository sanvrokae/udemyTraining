package main

import "fmt"

func main(){
	x := 42
	fmt.Println(x)
	{
		fmt.Println(x)
		y := "Quis Custodiet Ipsos Custodes"
		fmt.Println(y)
	}
	//fmt.Println(y)  //outside scope of y
}
