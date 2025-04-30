package config

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

type Database struct {
	DB     *sql.DB
	Logger *zap.Logger
}

func NewDatabase(logger *zap.Logger) (*Database, error) {
	// Carrega variáveis de ambiente
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	// Log das configurações (sem senha)
	logger.Info("Database configuration",
		zap.String("host", host),
		zap.String("port", port),
		zap.String("user", user),
		zap.String("dbname", dbname),
	)

	// String de conexão
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// Abre conexão com o banco
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %v", err)
	}

	// Configura conexão
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	// Testa a conexão
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error connecting to database: %v", err)
	}

	logger.Info("Successfully connected to database")

	return &Database{
		DB:     db,
		Logger: logger,
	}, nil
}

func (d *Database) Close() error {
	return d.DB.Close()
}
