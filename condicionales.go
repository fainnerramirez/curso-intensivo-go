package main

import (
	"fmt"
	"os"
)

func main() {

	edad := 20

	if edad >= 18 {
		fmt.Println("Eres mayor de edad")
	} else {
		fmt.Println("Eres menor de edad")
	}

	// bucles
	for i := 0; i < len(os.Args); i++ {
		fmt.Println("Argumento", i, ":", os.Args[i])
	}
}
