package modelos

import (
	seguranca "API/src/segunca"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

//Usuario representa um usuario utilizando a rede social
type Usuario struct {
	ID       uint64    `json: "id,omitempty"`
	Nome     string    `json: "nome,omitempty"`
	Nick     string    `json: "nick,omitempty"`
	Email    string    `json: "email,omitempty"`
	Senha    string    `json: "Senha,omitempty"`
	CriadoEm time.Time `json: "criadoEm,omitempty"`
}

// Preparar vai chamar os metodos para validar e formatar o usuario recebido
func (usuario *Usuario) Preparar(etapa string) error {
	if erro := usuario.validar(etapa); erro != nil {
		return erro
	}

	if erro := usuario.formatar(etapa); erro != nil {
		return erro
	}

	return nil
}

// função para validar os dados para não entrar vazio no banco de dados
func (usuario *Usuario) validar(etapa string) error {

	if usuario.Nome == "" {
		return errors.New("O nome é obrigatório e não pode estar em branco")
	}

	if usuario.Nick == "" {
		return errors.New("O nick é obrigatório e não pode estar em branco")
	}

	if usuario.Email == "" {
		return errors.New("O Email é obrigatório e não pode estar em branco")

	}
	//checkmail vem de um pacote de modulo externo do "github.com/badoux/checkmail v1.2.1"
	if erro := checkmail.ValidateFormat(usuario.Email); erro != nil {
		return errors.New("O e-mail inserido é invalido")
	}

	// esse if com metodo etapa deixa obrigatorio a senha obrigatorio só na fase de cadastro
	if etapa == "cadastro" && usuario.Senha == "" {
		return errors.New("O senha é obrigatório e não pode estar em branco")
	}

	return nil
}

// essa função formatar retira os espaços em branco do inicio e fim dos parametros
func (usuario *Usuario) formatar(etapa string) error {
	usuario.Nome = strings.TrimSpace(usuario.Nome)
	usuario.Nick = strings.TrimSpace(usuario.Nick)
	usuario.Email = strings.TrimSpace(usuario.Email)

	if etapa == "cadastro" {
		senhaComHash, erro := seguranca.Hash(usuario.Senha)
		if erro != nil {
			return erro
		}

		usuario.Senha = string(senhaComHash)
	}
	return nil
}
