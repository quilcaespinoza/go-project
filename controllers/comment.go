package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/golang-es/prjecomments/commons"
	"github.com/golang-es/prjecomments/configuration"
	"github.com/golang-es/prjecomments/models"
)

// CommentCreate para crear comentarios
func CommentCreate(w http.ResponseWriter, r *http.Request) {
	comment := models.Comment{}
	m := models.Message{}

	err := json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = fmt.Sprintf("Error al leer el comentario %s", err)

		commons.DisplayMessage(w, m)
		return
	}
	db := configuration.GetConnection()
	defer db.Close()

	err = db.Create(&comment).Error
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = fmt.Sprintf("Error al crear el comentario %s", err)

		commons.DisplayMessage(w, m)
		return
	}
	m.Code = http.StatusCreated
	m.Message = "Se creo satisfactoriamente el registro"
	commons.DisplayMessage(w, m)
}
