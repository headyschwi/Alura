package controllers

import (
	"Alura/api-rest-gin/database"
	"Alura/api-rest-gin/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CriaAluno(c *gin.Context) {
	var aluno models.Aluno
	err := c.ShouldBindJSON(&aluno)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Erro": err.Error(),
		})
		return
	}
	if err := models.ValidaDadosAluno(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Erro": err.Error(),
		})
		return
	}

	database.DB.Create(&aluno)
	c.JSON(http.StatusOK, aluno)
}

func ExibeTodosAlunos(c *gin.Context) {
	var alunos []models.Aluno
	database.DB.Find(&alunos)
	c.JSON(200, alunos)
}

func BuscaAlunoPorId(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")
	database.DB.First(&aluno, id)

	if aluno.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"NotFound": "Aluno não encontrado",
		})
		return
	}
	c.JSON(http.StatusOK, aluno)
}

func BuscaAlunoPorCPF(c *gin.Context) {
	var aluno models.Aluno
	cpf := c.Param("cpf")

	database.DB.Where(&models.Aluno{CPF: cpf}).First(&aluno)

	if aluno.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"NotFound:": "Aluno não encontrado."})
		return
	}

	c.JSON(http.StatusOK, aluno)
}

func AtualizaAluno(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")
	database.DB.First(&aluno, id)

	err := c.ShouldBindJSON(&aluno)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Erro": err.Error()})
		return
	}
	if err := models.ValidaDadosAluno(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Erro": err.Error(),
		})
		return
	}

	database.DB.Save(&aluno)
	c.JSON(http.StatusOK, aluno)
}

func DeletaAluno(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")
	database.DB.Delete(&aluno, id)
	c.JSON(http.StatusOK, gin.H{"data:": "Aluno deletado com sucesso."})
}

func Saudacao(c *gin.Context) {
	nome := c.Params.ByName("nome")

	c.JSON(200, gin.H{
		"API diz": "Olá, " + nome + ", tudo bem?",
	})
}

func ExibeHTMLIndex(c *gin.Context) {
	var alunos []models.Aluno
	database.DB.Find(&alunos)
	c.HTML(http.StatusOK, "index.html", gin.H{
		"alunos": alunos,
	})
}
func NotFound(c *gin.Context) {
	c.HTML(http.StatusNotFound, "404.html", nil)
}
