package handlers

import (
	"banana/pkg/web"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type EmployeeHandler struct {
	st map[string]string // meu banco de dados
}

func NewEmployeeHanlder(bancoDedados map[string]string) *EmployeeHandler {
	return &EmployeeHandler{
		st: bancoDedados,
	}
}

func (e *EmployeeHandler) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		web.ReponseJson(w, http.StatusOK, e.st, "Lista de usuarios encontrada com sucesso !")
	}
}

func (e *EmployeeHandler) GetById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//URL PARAMS
		id := chi.URLParam(r, "id")
		fmt.Println("ID: ", id)

		//Query Params
		// userdId := r.URL.Query().Get("userID")
		// fmt.Println("userID: ", userdId)

		//filtrando do banco de dados
		result, exists := e.st[id]

		if !exists {
			web.ReponseJson(w, http.StatusOK, nil, "Usuário não encontado !")
			return
		}

		web.ReponseJson(w, http.StatusOK, result, "usuario encontrado com sucesso !")
	}
}
