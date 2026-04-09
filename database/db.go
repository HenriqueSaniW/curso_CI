package database

import (
	"log"
	"time"
	"github.com/guilhermeonrails/api-go-gin/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaComBancoDeDados() {
	stringDeConexao := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(stringDeConexao))

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
