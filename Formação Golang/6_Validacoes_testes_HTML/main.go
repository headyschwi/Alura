package main

import (
	"Alura/api-rest-gin/database"
	"Alura/api-rest-gin/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequests()
}
