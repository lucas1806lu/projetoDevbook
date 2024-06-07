package modelos

import "time"

//Publicacao representa uma publicação feita por um usuario
type Publicacao struct {
	ID        uint64    `json: "id,omitempy"`
	Titulo    string    `json: "titulo,omitempy"`
	Conteudo  string    `json: "conteudo,omitempy"`
	AutorID   uint64    `json: "autorId,omitempy"`
	AutorNick string    `json: "autorNick,omitempy"`
	Curtidas  uint64    `json: "curtidas"`
	CriadaEm  time.Time `json: "criadaEm,omitempy"`
}
