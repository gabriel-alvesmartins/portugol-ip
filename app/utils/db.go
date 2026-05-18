package utils

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB

// ConnectToDB inicializa a conexão com o banco de dados PostgreSQL
func ConnectToDB() {
	// Carrega as variáveis de ambiente
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Aviso: arquivo .env não encontrado, usando variáveis de ambiente do sistema.")
	}

	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")

	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		user, password, dbname, host, port)

	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Erro ao conectar no banco de dados: ", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("Erro ao fazer ping no banco de dados: ", err)
	}

	fmt.Println("Conexão com o banco de dados estabelecida com sucesso.")
	
	createTables()
}

// createTables cria a tabela de pacientes se não existir
func createTables() {
	query := `
	CREATE TABLE IF NOT EXISTS pacientes (
		id SERIAL PRIMARY KEY,
		nome VARCHAR(100) NOT NULL,
		email VARCHAR(100) UNIQUE NOT NULL,
		idade INT NOT NULL,
		diagnostico TEXT,
		telefone VARCHAR(20),
		cpf VARCHAR(14) UNIQUE
	);`

	_, err := DB.Exec(query)
	if err != nil {
		log.Fatal("Erro ao criar a tabela pacientes: ", err)
	}
	fmt.Println("Tabela pacientes verificada/criada com sucesso.")
}
