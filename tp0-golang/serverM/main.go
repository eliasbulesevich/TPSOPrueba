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

	muxMemory := http.NewServeMux()

	muxMemory.HandleFunc("/mensaje", utils.RecibirMensaje)

	//panic("no implementado!")
	errMemory := http.ListenAndServe(":8002", muxMemory)
	if errMemory != nil {
		panic(errMemory)
	}

}
