package routes

import (
	"Alura/api-rest-gin/controllers"

	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.POST("/alunos", controllers.CriaAluno)
	r.GET("/alunos", controllers.ExibeTodosAlunos)
	r.GET("/alunos/:id", controllers.BuscaAlunoPorId)
	r.GET("/alunos/cpf/:cpf", controllers.BuscaAlunoPorCPF)
	r.PATCH("alunos/:id", controllers.AtualizaAluno)
	r.DELETE("/alunos/:id", controllers.DeletaAluno)
	r.GET("/:nome", controllers.Saudacao)
	r.Run(":5000")
}
