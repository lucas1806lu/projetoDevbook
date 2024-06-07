package modelos

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
	"webapp/src/config"
	"webapp/src/requisicoes"
)

//Usuario representa uma pessoa na rede social
type Usuario struct {
	ID          uint64       `json:"id"`
	Nome        string       `json:"nome"`
	Nick        string       `json:"nick"`
	Email       string       `json:"email"`
	CriadoEm    time.Time    `json:"criadoEm"`
	Seguidores  []Usuario    `json:"seguidores"`
	Seguindo    []Usuario    `json:"seguindo"`
	Publicacoes []Publicacao `json:"publicacoes"`
}

//BuscarUsuarioCompleto faz 4 requisiçoes na API para mostrar o usuario
func BuscarUsuarioCompleto(UsuarioID uint64, r *http.Request) (Usuario, error) {
	canalUsuario := make(chan Usuario)
	canalSeguidores := make(chan []Usuario)
	canalSeguindo := make(chan []Usuario)
	canalPublicacoes := make(chan []Publicacao)

	go BuscarDadosDoUsuario(canalUsuario, UsuarioID, r)
	go BuscarSeguidores(canalSeguidores, UsuarioID, r)
	go BuscarSeguidores(canalSeguindo, UsuarioID, r)
	go BuscarPublicacoes(canalPublicacoes, UsuarioID, r)

	var (
		usuario     Usuario
		seguidores  []Usuario
		seguindo    []Usuario
		publicacoes []Publicacao
	)

	for i := 0; i < 4; i++ {
		select {
		case usuarioCarregado := <-canalUsuario:
			if usuarioCarregado.ID == 0 {
				return Usuario{}, errors.New("Erro ao buscar usuario")
			}
			usuario = usuarioCarregado

		case seguidoresCarregados := <-canalSeguidores:
			if seguidoresCarregados == nil {
				return Usuario{}, errors.New("Erro ao buscar os seguidores")
			}

			seguidores = seguidoresCarregados

		case seguindoCarregados := <-canalSeguindo:
			if seguindoCarregados == nil {
				return Usuario{}, errors.New("Erro ao buscar os usuarios seguindo")
			}

			seguindo = seguindoCarregados

		case publicacoesCarregados := <-canalPublicacoes:
			if publicacoesCarregados == nil {
				return Usuario{}, errors.New("Erro ao buscar as Publicações ")
			}

			publicacoes = publicacoesCarregados
		}

	}

	usuario.Seguidores = seguidores
	usuario.Seguidores = seguindo
	usuario.Publicacoes = publicacoes

	return usuario, nil
}

//BuscarDadosDoUsuario chama a API para buscar dadosa base de usuario
func BuscarDadosDoUsuario(canal chan<- Usuario, usuarioID uint64, r *http.Request) {

	url := fmt.Sprintf("%s/usuarios/%d", config.APIURL, usuarioID)
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	if erro != nil {

		canal <- Usuario{}
		return
	}
	defer response.Body.Close()

	var usuario Usuario
	if erro = json.NewDecoder(response.Body).Decode(&usuario); erro != nil {

		canal <- Usuario{}
		return
	}

	canal <- usuario
}

//BuscarSeguidores chama a API para buscar dados de seguidores do usuario
func BuscarSeguidores(canal chan<- []Usuario, usuarioID uint64, r *http.Request) {

	url := fmt.Sprintf("%s/usuarios/%d/seguidores", config.APIURL, usuarioID)
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	if erro != nil {

		canal <- nil
		return
	}
	defer response.Body.Close()

	var seguidores []Usuario
	if erro = json.NewDecoder(response.Body).Decode(&seguidores); erro != nil {
		canal <- nil
		return
	}

	if seguidores == nil {
		canal <- make([]Usuario, 0)
		return
	}

	canal <- seguidores
}

//BuscarSeguindo chama a API para buscar dados de usuario sendo seguidos
func BuscarSeguindo(canal chan<- []Usuario, usuarioID uint64, r *http.Request) {

	url := fmt.Sprintf("%s/usuarios/%d/seguindo", config.APIURL, usuarioID)
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	if erro != nil {

		canal <- nil
		return
	}
	defer response.Body.Close()

	var seguindo []Usuario
	if erro = json.NewDecoder(response.Body).Decode(&seguindo); erro != nil {
		canal <- nil
		return

	}

	if seguindo == nil {
		canal <- make([]Usuario, 0)
		return
	}
	canal <- seguindo
}

//BuscarPublicacoes chama a API para buscar dados da publicaçoes de um usuario
func BuscarPublicacoes(canal chan<- []Publicacao, usuarioID uint64, r *http.Request) {
	url := fmt.Sprintf("%s/usuarios/%d/publicacoes", config.APIURL, usuarioID)
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	if erro != nil {
		canal <- nil
		return
	}

	defer response.Body.Close()

	var publicacoes []Publicacao
	if erro = json.NewDecoder(response.Body).Decode(&publicacoes); erro != nil {
		canal <- nil
		return
	}

	if publicacoes == nil {
		canal <- make([]Publicacao, 0)
		return
	}

	canal <- publicacoes
}
