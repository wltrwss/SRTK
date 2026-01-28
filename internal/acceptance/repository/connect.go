package repository

import (
	"database/sql"
	"fmt"
	"log/slog"
	"os"

	"github.com/joho/godotenv"
)

func ConnectDB(logger *slog.Logger) (*sql.DB, error) {
	//Подгрузка коннекта в переменную серды окружения
	_ = godotenv.Load()

	connStr := os.Getenv("DATABASE_URL")
	if connStr == "" {
		return nil, fmt.Errorf("DATABASE_URL is not set")
	}

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("При попытке подключения к БД произошла ошибка:%w", err)
	}
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("При проверке соединения с БД произошла ошибка:%w", err)
	}
	logger.Info("Соединение с БД установлено!")
	return db, nil
}
