package main

import (
	extraer_Instagram "backend/funciones"
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/get-profile-data", extraer_Instagram.Enviar_info)

	// Envuelves tu handler original con el middleware CORS
	handler := extraer_Instagram.EnableCors(mux)

	fmt.Println("Servidor corriendo en http://localhost:8080")
	if err := http.ListenAndServe(":8080", handler); err != nil {
		fmt.Println("Error al iniciar el servidor:", err)
	}
}
