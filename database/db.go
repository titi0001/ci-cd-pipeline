package database

import (
	"log"
	// "os"
	"github.com/titi0001/golang-automated-delivery/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaComBancoDeDados() {

	stringDeConexao := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"
	// stringDeConexao := "host="+os.Getenv("HOST")+" user="+os.Getenv("USER")+" password="+os.Getenv("PASSWORD")+" dbname="+os.Getenv("DBNAME")+"port=5432"+os.Getenv("PORT")+"sslmode=disable"

	DB, err = gorm.Open(postgres.Open(stringDeConexao))
	if err != nil {
		log.Panic("Erro ao conectar com banco de dados")
	}

	DB.AutoMigrate(&models.Aluno{})
}
