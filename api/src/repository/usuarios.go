package repository

import (
	"api/src/modelos"
	"database/sql"
)

type Usuarios struct {
	db *sql.DB
}

// NovoRepositorioDeUsuarios cria um repositório de usuários
func NovoRepositorioDeUsuarios(db *sql.DB) *Usuarios {
	return &Usuarios{db: db}
}

// Criar insere um usuário no banco de dados
func (u Usuarios) Criar(usuarios modelos.Usuario) (uint64, error) {
	return 0, nil
}
