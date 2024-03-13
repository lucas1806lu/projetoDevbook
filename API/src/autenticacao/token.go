package autenticacao

import (
	"API/src/config"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// esse pacote exetrno github.com/dgrijalva/jwt-go é responsavel por converter string em token e vice e verso.

//CriarToken retorna m token assinadocom as permissoes de usuario
func CriarToken(usuarioID uint64) (string, error) {
	permissoes := jwt.MapClaims{}
	permissoes["authorized"] = true
	permissoes["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permissoes["usuarioId"] = usuarioID
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissoes)

	return token.SignedString([]byte(config.SecretKey)) //secret
}

//ValidarToken verifica se o token passado n requisissão é valido
func ValidrToken(r *http.Request) error {
	tokenString := extrairToken(r)
	token, erro := jwt.Parse(tokenString, retornaChaveVerificada)
	if erro != nil {
		return erro
	}
	if  _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid{
		return nil
	} 
	return errors.New("Token invalido")

}

//ExtrairUsuarioPorID retorna o usuarioId qe está salvo no token
func ExtrairUsuarioPorID(r *http.Request) (uint64, error){

    tokenString := extrairToken(r)
	token, erro := jwt.Parse(tokenString, retornaChaveVerificada)
	if erro != nil {
		return 0, erro
	}

	if  permissoes, ok := token.Claims.(jwt.MapClaims); ok && token.Valid{
		usuarioID, erro := strconv.ParseUint(fmt.Sprintf("%.0f", permissoes["usuarioId"]), 10, 64)
		if erro != nil{
			return 0, erro
		}
		return usuarioID, nil
	} 
	return 0, errors.New("Token Invalido! ")

}


func extrairToken(r *http.Request) string {
	token := r.Header.Get("Authorization")

	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}

	return ""
}

func retornaChaveVerificada(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Método de assinatura inesperado! %v", token.Header["alg"])
	}
	return config.SecretKey, nil
}
