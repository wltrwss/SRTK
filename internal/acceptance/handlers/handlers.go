package handlers

import (
	"database/sql"
	"encoding/json"
	"log/slog"
	"net/http"
	"srtk/internal/acceptance/items"
	"srtk/internal/acceptance/service"
)

type Handler struct {
	db     *sql.DB
	logger *slog.Logger
}

func NewHandler(db *sql.DB, logger *slog.Logger) Handler {
	return Handler{db: db, logger: logger}
}

func JSONDecode[T any](r *http.Request, object *T) error {
	return json.NewDecoder(r.Body).Decode(object)
}

// внедрить проверку на метод!
func (h Handler) PositionsHandler(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("Поступил запрос PositionsHandler")
	var object items.Position

	if err := JSONDecode(r, &object); err != nil {
		h.logger.Error("Сбой при расшифровке данных!", slog.String("description", err.Error()))
		http.Error(w, "Сбой при расшифровке данных! Проверьте правильность отправляемых данных.", http.StatusBadRequest)
		return
	}

	if err := service.SavePosition(object, h.db, h.logger); err != nil {
		h.logger.Error("Направленные данные имеют ошибку или не соотвествуют требованиям!", slog.String("description", err.Error()))
		http.Error(w, "Направленные данные имеют ошибку или не соотвествуют требованиям!", http.StatusUnprocessableEntity)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ОК"))
	h.logger.Info("Запрос успешно обработан PositionsHandler")
}

func (h Handler) ScannerHandler(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("Поступил запрос ScannerHandler")
	var object items.Scanner

	if err := JSONDecode(r, &object); err != nil {
		h.logger.Error("Сбой при расшифровке данных!", slog.String("description", err.Error()))
		http.Error(w, "Сбой при расшифровке данных! Проверьте правильность отправляемых данных.", http.StatusBadRequest)
		return
	}

	if err := service.CheckScanner(object, h.db, h.logger); err != nil {
		h.logger.Error("Сбой в работе функций!", slog.String("description", err.Error()))
		http.Error(w, "Проверьте подключение сканнера!", http.StatusUnprocessableEntity)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ОК"))
	h.logger.Info("Запрос успешно обработан ScannerHandler")
}

func (h Handler) SendPositionsHandler(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("Поступил запрос PositionsHandler")
	var object items.Signal

	if err := JSONDecode(r, &object); err != nil {
		h.logger.Error("Сбой при расшифровке данных!", slog.String("description", err.Error()))
		http.Error(w, "Сбой при расшифровке данных! Проверьте правильность отправляемых данных.", http.StatusBadRequest)
		return
	}

	if p, err := service.CheckSignal(object, h.db, h.logger); err != nil {
		h.logger.Error("Направленные данные имеют ошибку или не соотвествуют требованиям!", slog.String("description", err.Error()))
		http.Error(w, "Направленные данные имеют ошибку или не соотвествуют требованиям!", http.StatusUnprocessableEntity)
		return
	} else {

	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ОК"))
	h.logger.Info("Запрос успешно обработан PositionsHandler")
}
