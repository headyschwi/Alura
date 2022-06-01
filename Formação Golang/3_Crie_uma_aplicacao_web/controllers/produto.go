package controllers

import (
	"alura/aplicativo-web/models"
	"net/http"
	"strconv"
	"text/template"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, request *http.Request) {
	produtos := models.RetornaTodosProdutos()
	temp.ExecuteTemplate(w, "Index", produtos)
}

func New(w http.ResponseWriter, request *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, request *http.Request) {
	if request.Method == "POST" {
		nome := request.FormValue("nome")
		descricao := request.FormValue("descricao")
		valorStr := request.FormValue("valor")
		quantidadeStr := request.FormValue("quantidade")

		valor, _ := strconv.ParseFloat(valorStr, 64)
		quantidade, _ := strconv.Atoi(quantidadeStr)

		models.CriaNovoProduto(nome, descricao, valor, quantidade)
	}
	http.Redirect(w, request, "/", 301)
}

func Delete(w http.ResponseWriter, request *http.Request) {
	id := request.URL.Query().Get("id")
	models.DeletaProduto(id)
	http.Redirect(w, request, "/", 301)
}

func Edit(w http.ResponseWriter, request *http.Request) {
	id := request.URL.Query().Get("id")
	produto := models.EditaProduto(id)
	temp.ExecuteTemplate(w, "Edit", produto)
}

func Update(w http.ResponseWriter, request *http.Request) {
	if request.Method == "POST" {
		idStr := request.FormValue("id")
		nome := request.FormValue("nome")
		descricao := request.FormValue("descricao")
		valorStr := request.FormValue("valor")
		quantidadeStr := request.FormValue("quantidade")

		id, _ := strconv.Atoi(idStr)
		quantidade, _ := strconv.Atoi(quantidadeStr)
		valor, _ := strconv.ParseFloat(valorStr, 64)

		models.AtualizaProduto(nome, descricao, id, quantidade, valor)
	}
	http.Redirect(w, request, "/", 301)
}
