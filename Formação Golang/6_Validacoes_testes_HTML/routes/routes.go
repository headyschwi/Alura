package routes

import (
	"Alura/api-rest-gin/controllers"

	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./assets")
	r.NoRoute(controllers.NotFound)

	r.POST("/alunos", controllers.CriaAluno)
	r.GET("/alunos", controllers.ExibeTodosAlunos)
	r.GET("/alunos/:id", controllers.BuscaAlunoPorId)
	r.GET("/alunos/cpf/:cpf", controllers.BuscaAlunoPorCPF)
	r.PATCH("alunos/:id", controllers.AtualizaAluno)
	r.DELETE("/alunos/:id", controllers.DeletaAluno)
	r.GET("/:nome", controllers.Saudacao)

	r.GET("/index", controllers.ExibeHTMLIndex)
	r.Run(":5000")
}
