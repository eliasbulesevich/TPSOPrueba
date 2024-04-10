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

	muxInterface := http.NewServeMux()

	muxInterface.HandleFunc("/mensaje", utils.RecibirMensaje)

	//panic("no implementado!")
	errInterface := http.ListenAndServe(":8005", muxInterface)
	if errInterface != nil {
		panic(errInterface)
	}

}
