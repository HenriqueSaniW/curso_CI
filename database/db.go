package database

import (
	"log"
	"time"
	"os"
	"github.com/guilhermeonrails/api-go-gin/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaComBancoDeDados() {
    host := os.Getenv("HOST")
    user := os.Getenv("USER")
    password := os.Getenv("PASSWORD")
    dbname := os.Getenv("DBNAME")
    port := os.Getenv("DBPORT")
    stringDeConexao := "host="+ host +" user=" + user + " password=" + password + " dbname=" + dbname + " port=" + port + " sslmode=require"

	for i := 0; i < 10; i++ {
		DB, err = gorm.Open(postgres.Open(stringDeConexao))
		if err == nil {
			break
		}
		log.Println("Aguardando banco de dados iniciar...")
		time.Sleep(2 * time.Second)
	}

	if err != nil {
		log.Panic("Erro ao se conectar com o banco de dados")
	}
	_ = DB.AutoMigrate(&models.Aluno{})
}
