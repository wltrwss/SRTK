package repository

import (
	"database/sql"
	"fmt"
	"log/slog"
	"math/rand"
	"srtk/internal/acceptance/items"
	"time"
)

var currentTableName string

func CreateTableName() string {
	randomNum := rand.Intn(900000000) + 100000000
	todayDate := time.Now().Format("20060102")
	return fmt.Sprintf("positions_%s_%d", todayDate, randomNum)
}

func CreateTable(db *sql.DB, logger *slog.Logger) error {
	tableName := CreateTableName()
	currentTableName = tableName
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id				 BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
		barcode          TEXT, 
		name             TEXT,   
		unit_measurement TEXT,   
		quantity         FLOAT,  
		price_buy      	 FLOAT,  
		price_sell       FLOAT, 
		date             DATE
	)`, tableName)
	_, err := db.Exec(query)
	if err != nil {
		return fmt.Errorf("Не удалось создать таблицу: %w", err)
	}
	logger.Info(fmt.Sprintf("Таблица %s успешно создана!", tableName))
	return nil
}

func InsertTable(object items.Position, db *sql.DB, logger *slog.Logger) error {
	query := fmt.Sprintf(`
        INSERT INTO %s (barcode, name, unit_measurement, quantity, price_buy, price_sell, date)
        VALUES ($1, $2, $3, $4, $5, $6, $7);`, currentTableName)

	_, err := db.Exec(query,
		object.Barcode,
		object.Name,
		object.UnitMeasurement,
		object.Quantity,
		object.Price_buy,
		object.Price_sell,
		object.Date,
	)

	if err != nil {
		return fmt.Errorf("Не удалось записать структуру в таблицу: %w", err)
	}

	logger.Info(fmt.Sprintf("Внесена запись в таблицу %s!", currentTableName))
	return nil
}

func UploadTable(db *sql.DB) ([]items.Position, error) {
	query := fmt.Sprintf(`
		SELECT
			barcode,
			name,
			unit_measurement,
			quantity,
			price_buy,
			price_sell,
			date
		FROM %s
	`, currentTableName)

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []items.Position

	for rows.Next() {
		var r items.Position
		if err := rows.Scan(
			&r.Barcode,
			&r.Name,
			&r.UnitMeasurement,
			&r.Quantity,
			&r.Price_buy,
			&r.Price_sell,
			&r.Date,
		); err != nil {
			return nil, err
		}
		result = append(result, r)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
