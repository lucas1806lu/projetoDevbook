package main

import (
	"API/src/config"
	"API/src/router"
	
	"fmt"
	"log"
	"net/http"
)





func main() {

	config.Carregar()
	
	r := router.Gerar()

	fmt.Println(config.SecretKey)

	fmt.Printf("Escutando na porta : %d\n",config.Porta)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))

}
