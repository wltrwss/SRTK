package service

import (
	"database/sql"
	"fmt"
	"log/slog"
	"srtk/internal/acceptance/items"
	"srtk/internal/acceptance/repository"
	"time"
)

func CheckPosition(object items.Position) error {
	if object.Barcode == "" {
		return fmt.Errorf("Отсутсвует штрих-код!")
	}
	if object.Name == "" {
		return fmt.Errorf("Отсутсвует имя позиции товара!")
	}
	if object.UnitMeasurement == "" {
		return fmt.Errorf("Отсутсвует единица измерения")
	}
	if object.Quantity == 0 {
		return fmt.Errorf("Отсутсвует колличество")
	}
	if object.Price_buy == 0 {
		return fmt.Errorf("Отсутсвует цена закупки")
	}
	if object.Price_sell == 0 {
		return fmt.Errorf("Отсутсвует цена реализации")
	}
	return nil
}

func SavePosition(object items.Position, db *sql.DB, logger *slog.Logger) error {
	if err := CheckPosition(object); err != nil {
		return err
	}

	now := time.Now()
	object.Date = &now

	return repository.InsertTable(object, db, logger)
}

func CheckScanner(object items.Scanner, db *sql.DB, logger *slog.Logger) error {
	if !object.Status {
		return fmt.Errorf("Сканер не обнаружен! Проверьте подключение.")
	}
	return repository.CreateTable(db, logger)
}

func CheckSignal(object items.Signal, db *sql.DB, logger *slog.Logger) ([]items.Position, error) {
	if !object.Status {
		return nil, fmt.Errorf("Я НЕ ЗНАЮ КАКИМ ОБРАЗОМ ВОЗНИКЛА ЭТА ОШИБКА. ЭТОГО НЕ ДОЛЖНО БЫЛО БЫТЬ")
	}

	var p []items.Position
	p, err := repository.UploadTable(db)
	return p, err
}

/*ТЕСТОВЫЕ ЗАПРОСЫ ДЛЯ POSTMAN
{
    "barcode":"01101",
    "name":"string",
    "unit_measurement":"string",
    "quantity": 100.1,
    "price_buy": 200.43,
    "price_sell": 300.34,
    "date": null
}

{
    "status":false
}
*/
