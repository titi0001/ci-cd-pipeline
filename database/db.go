package database

import (
	"log"
	"os"

	"github.com/titi0001/golang-automated-delivery/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"github.com/joho/godotenv"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaComBancoDeDados() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Erro ao carregar .env", err)
		os.Exit(1)
	}

	stringDeConexao := "host="+os.Getenv("HOST")+" user="+os.Getenv("USERDB")+" password="+os.Getenv("PASSWORD")+" dbname="+os.Getenv("DBNAME")+" port="+os.Getenv("PORT")+" sslmode=disable"

	DB, err = gorm.Open(postgres.Open(stringDeConexao))
	if err != nil {
		log.Panic("Erro ao conectar com banco de dados")
	}

	DB.AutoMigrate(&models.Aluno{})
}
