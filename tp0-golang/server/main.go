package main

import (
	"net/http"
	"server/utils"
)

func main() {

	//mux := http.NewServeMux()

	//mux.HandleFunc("/paquetes", utils.RecibirPaquetes)
	//mux.HandleFunc("/mensaje", utils.RecibirMensaje)

	//panic("no implementado!")
	//err := http.ListenAndServe(":8080", mux)
	//if err != nil {
	//	panic(err)
	//}

	// ----- Levantamos los 4 modulos ------//

	muxKernel := http.NewServeMux()

	muxKernel.HandleFunc("/mensaje", utils.RecibirMensaje)

	//panic("no implementado!") -- PONER VARIABLE EN PUERTO
	errKernel := http.ListenAndServe(":8001", muxKernel)
	if errKernel != nil {
		panic(errKernel)
	}

}
