package controllers

import (
	"net/http"
	"webapp/src/cookies"
)

//FazerLogout remove os dados de autenticaçao salvos no browser do usuario
func FazerLogout(w http.ResponseWriter, r *http.Request) {
	cookies.Deletar(w)
	http.Redirect(w, r, "login", 302)
}
