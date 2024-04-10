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

	muxCpu := http.NewServeMux()

	muxCpu.HandleFunc("/mensaje", utils.RecibirMensaje)

	//panic("no implementado!")
	errCpu := http.ListenAndServe(":8003", muxCpu)
	if errCpu != nil {
		panic(errCpu)
	}

}
