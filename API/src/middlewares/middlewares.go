package middlewares

import (
	"API/src/autenticacao"
	"API/src/respostas"
	"log"
	"net/http"
)

//Logger escreve informações da requisição no terminal
func Logger(proximaFucao http.HandlerFunc) http.HandlerFunc{
	return func (w http.ResponseWriter, r *http.Request){
		log.Printf("\n %s %s %s", r.Method, r.RequestURI, r.Host)
		proximaFucao(w, r)
	}
}

// Autenticar erifica se ousuario fazendo a requisição está autenticado
func Autenticar(proximaFucao http.HandlerFunc) http.HandlerFunc{
	return func (w http.ResponseWriter, r *http.Request){
		if erro := autenticacao.ValidrToken(r); erro != nil{
			respostas.Erro(w, http.StatusUnauthorized, erro)
			return
		}
		
		proximaFucao(w, r)
	}
}