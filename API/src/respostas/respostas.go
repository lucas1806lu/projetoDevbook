package respostas

import (
	"encoding/json"
	"log"
	"net/http"
)

// json retorna uma resposta em json para a requisição
func JSON(w http.ResponseWriter, statusCode int, dados interface{}) {

	//a resposta no postman retorna em json content TYpe de resposta com dois parametros em string
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if dados != nil {

		if erro := json.NewEncoder(w).Encode(dados); erro != nil {
			log.Fatal(erro)
		}

	}
}

//Erro retorna um erro em formato JSON
func Erro(w http.ResponseWriter, statusCode int, erro error) {
	JSON(w, statusCode, struct {
		Erro string `json: "erro"`
	}{
		Erro: erro.Error(),
	})
}
