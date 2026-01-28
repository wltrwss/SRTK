package main

import (
	"log/slog"
	"net/http"
	"os"
	"srtk/internal/acceptance/handlers"
	"srtk/internal/acceptance/repository"

	_ "github.com/lib/pq"
)

func main() {
	logger := slog.New(
		slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level:     slog.LevelInfo,
			AddSource: true,
		}),
	)

	db, err := repository.ConnectDB(logger)
	if err != nil {
		logger.Error("Ошибка подключения к БД", slog.String("description", err.Error()))
		return
	}
	defer db.Close()
	defer logger.Info("Соединение с БД разорвано.")

	handler := handlers.NewHandler(db, logger) //Для использоавния db в хендлерах

	http.HandleFunc("/acceptance/scaner", handler.ScannerHandler)
	http.HandleFunc("/acceptance/position", handler.PositionsHandler)
	http.HandleFunc("/acceptance/position/send", handler.SendPositionsHandler)

	logger.Info("Сервер запущен. Порт: 1488")
	if err := http.ListenAndServe(":1488", nil); err != nil {
		logger.Error("Аварийное завершение работы сервера. Причина:", slog.String("description", err.Error()))
	}
}

//==============ЗАМЕТКИ==============//
//cmd.exe /c "chcp 1251 && psql -U postgres -d acceptance_service_db" - запуск БД в нужной кодировке

//============РАЗОБРАТЬСЯ============//
//ЧТО ВЫЗЫВАТЬ РАНЬШЕ СЕРВЕР ИЛИ БД? БД

//КАК СДЕЛАТЬ ЧТОБЫ ОСНОВНОЙ ПОТОК НЕ БЛОКИРОВАЛСЯ
//ПРИ ПОДНЯТИИ СЕРВЕРА? ОН ДОЛЖЕН БЛОКИРОВАТЬСЯ ЭТО НОРМАЛЬНО

//НУЖНО ЛИ ПОДНИМАТЬ ОШИБКУ ИЗ БД НА УРОВЕНЬ ХЕНДЛЕРА ЧЕРЕЗ
//СЕРВИС? ИЛИ ВЫВОДИТЬ СРАЗУ ЖЕ? ДА, ВЫВОДИТЬ В ХЕНДЛЕР.
