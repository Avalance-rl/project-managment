package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	// TODO: сделать структруы данных для проектов, задач и тп. Настроить связи

	// TODO: интегрировать psql

	// TODO: API для управления, эндпоинты

	// TODO: бизнес-логика
	http.ListenAndServe(":8080", r)
}
