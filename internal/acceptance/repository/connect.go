package repository

import (
	"database/sql"
	"fmt"
	"log/slog"
)

func ConnectDB(logger *slog.Logger) (*sql.DB, error) {
	connStr := "postgres://postgres:1488@localhost:5432/acceptance_service_db?sslmode=disable"
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
