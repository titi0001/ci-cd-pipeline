package main

import (
	"github.com/titi0001/golang-automated-delivery/database"
	"github.com/titi0001/golang-automated-delivery/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequest()
}
