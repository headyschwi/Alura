package main

import (
	"ApiRest/database"
	"ApiRest/routes"
	"fmt"
)

func main() {
	database.ConectaComBancoDeDados()

	fmt.Println("Iniciando o servidor Rest com Go")
	routes.HandleRequest()
}
