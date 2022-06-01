package main

// Este arquivo serve para exemplificar operações CRUD usando o GORM

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Curso struct {
	gorm.Model
	Nome, Area, Professor string
}

func main() {
	db, err := gorm.Open(postgres.Open("artigo.db"), &gorm.Config{})
	if err != nil {
		panic("erro ao conectar com o banco de dados")
	}

	db.AutoMigrate(&Curso{})

	//Create Curso
	db.Create(&Curso{Nome: "Go: Fundamentos de uma aplicacao web", Area: "Programacao", Professor: "Guilherme Lima"})

	//Read Curso
	var curso Curso
	db.First(&curso, 1)

	//Update Curso
	db.Model(&curso).Update("Nome", "Go: Fundamentos de uma aplicação web com GORM")

	//Delete Curso
	db.Delete(&curso, 1)
}
