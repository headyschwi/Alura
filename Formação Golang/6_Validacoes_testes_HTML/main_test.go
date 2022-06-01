package main

import (
	"Alura/api-rest-gin/controllers"
	"Alura/api-rest-gin/database"
	"Alura/api-rest-gin/models"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetupRotasDeTeste() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	return r
}

func TestSaudacao(t *testing.T) {
	r := SetupRotasDeTeste()
	r.GET("/:nome", controllers.Saudacao)

	req, _ := http.NewRequest("GET", "/eder", nil)
	resposta := httptest.NewRecorder()

	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code)

	mockResposta := `{"API diz":"Olá, eder, tudo bem?"}`
	respostaBody, _ := ioutil.ReadAll(resposta.Body)

	assert.Equal(t, mockResposta, string(respostaBody))
}

var ID int

func CriaAlunoMock() {
	aluno := models.Aluno{Nome: "Aluno teste", CPF: "12345678901", RG: "123456789"}
	database.DB.Create(&aluno)
	ID = int(aluno.ID)
}

func DeletaAlunoMock() {
	var aluno models.Aluno
	database.DB.Delete(&aluno, ID)
}

func TestExibeTodosAlunos(t *testing.T) {
	database.ConectaComBancoDeDados()
	CriaAlunoMock()
	defer DeletaAlunoMock()

	r := SetupRotasDeTeste()

	r.GET("/alunos", controllers.ExibeTodosAlunos)
	req, _ := http.NewRequest("GET", "/alunos", nil)
	resposta := httptest.NewRecorder()

	r.ServeHTTP(resposta, req)

	assert.Equal(t, http.StatusOK, resposta.Code)
	fmt.Println(resposta.Body)
}

func TestBuscaAlunoPorCPF(t *testing.T) {
	database.ConectaComBancoDeDados()
	CriaAlunoMock()
	defer DeletaAlunoMock()

	r := SetupRotasDeTeste()

	r.GET("/alunos/cpf/:cpf", controllers.BuscaAlunoPorCPF)

	req, _ := http.NewRequest("GET", "/alunos/cpf/12345678901", nil)
	resposta := httptest.NewRecorder()

	r.ServeHTTP(resposta, req)

	assert.Equal(t, http.StatusOK, resposta.Code)
}

func TestBuscaAlunoPorId(t *testing.T) {
	database.ConectaComBancoDeDados()
	CriaAlunoMock()
	defer DeletaAlunoMock()
	r := SetupRotasDeTeste()

	r.GET("/alunos/:id", controllers.BuscaAlunoPorId)

	pathBusca := fmt.Sprintf("/alunos/%d", ID)

	req, _ := http.NewRequest("GET", pathBusca, nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)

	var alunoMock models.Aluno
	json.Unmarshal(resposta.Body.Bytes(), &alunoMock)

	assert.Equal(t, "Aluno teste", alunoMock.Nome)
}

func TestDeletaAluno(t *testing.T) {
	database.ConectaComBancoDeDados()
	CriaAlunoMock()
	r := SetupRotasDeTeste()
	r.DELETE("/alunos/:id", controllers.DeletaAluno)

	path := fmt.Sprintf("/alunos/%d", ID)

	req, _ := http.NewRequest("DELETE", path, nil)
	resposta := httptest.NewRecorder()

	r.ServeHTTP(resposta, req)

	assert.Equal(t, http.StatusOK, resposta.Code)
}

func TestAtualizaAluno(t *testing.T) {
	database.ConectaComBancoDeDados()
	CriaAlunoMock()
	defer DeletaAlunoMock()

	r := SetupRotasDeTeste()
	r.PATCH("/alunos/:id", controllers.AtualizaAluno)
	aluno := models.Aluno{Nome: "Aluno teste", CPF: "12345678901", RG: "123456700"}
	alunojson, _ := json.Marshal(aluno)

	path := fmt.Sprintf("/alunos/%d", ID)
	req, _ := http.NewRequest("PATCH", path, bytes.NewBuffer(alunojson))
	resposta := httptest.NewRecorder()

	r.ServeHTTP(resposta, req)

	var alunoMockAtualizado models.Aluno
	json.Unmarshal(resposta.Body.Bytes(), &alunoMockAtualizado)

	assert.Equal(t, alunoMockAtualizado.CPF, aluno.CPF)
}
