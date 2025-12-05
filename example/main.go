package main

import (
	"fmt"
	"time"

	"github.com/dbinario/rfc-mx"
)

func main() {
	datos := rfc.DatosFisica{
		Nombre:          "Juan",
		ApellidoPaterno: "Gomez",
		ApellidoMaterno: "Diaz",
		FechaNacimiento: time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC),
	}

	clave, err := rfc.GeneraRFCFisica(datos)
	if err != nil {
		panic(err)
	}

	fmt.Println("RFC:", clave)
}
