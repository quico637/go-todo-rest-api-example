package handler

import (
	"encoding/json"
	"net/http"

	"github.com/quico637/go-todo-rest-api-example/app/model"
)

func GetAllProjects(w http.ResponseWriter, r *http.Request) {
	projects := model.Project{"hola", "adios", true}
	respondJSON(w, http.StatusOK, projects)
}

func CreateProject(w http.ResponseWriter, r *http.Request) {
	project := model.Project{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&project); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	respondJSON(w, http.StatusCreated, project)
}
