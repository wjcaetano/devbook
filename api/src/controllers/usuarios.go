package controllers

import (
	"api/src/banco"
	"api/src/modelos"
	"api/src/repository"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

// CriarUsuario insere um usuário no banco de dados
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	corpoRequest, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	var usuario modelos.Usuario
	if err = json.Unmarshal(corpoRequest, &usuario); err != nil {
		log.Fatal(err)
	}

	db, err := banco.Conectar()
	if err != nil {
		log.Fatal(err)
	}

	repositorio := repository.NovoRepositorioDeUsuarios(db)
	repositorio.Criar(usuario)
}

// BuscarUsuarios busca todos os usuário no banco de dados
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando todos os usuários"))
}

// BuscarUsuario busca um usuário no banco de dados
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando um usuário"))
}

// AtualizarUsuario altera dados de um usuário no banco de dados
func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizando usuário"))
}

// DeletarUsuario deleta um usuário no banco de dados
func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deletando usuários"))
}
