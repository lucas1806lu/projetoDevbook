package controllers

import (
	"net/http"
	"webapp/src/utils"
)

// CarregarTelaDeLogin vai redirecionar para tela de login
func CarregarTelaDeLogin(w http.ResponseWriter, r *http.Request) {

	utils.ExecutarTemplate(w, "login.html", nil)
}

//CarregarPaginaDeCadastroDeUsuario vai carregar para pagina de castro de usuario
func CarregarPaginaDeCadastroDeUsuario(w http.ResponseWriter, r *http.Request) {

	utils.ExecutarTemplate(w, "cadastro.html", nil)

}
