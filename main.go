package main

import (
	"log"
	"net/http"

	"github.com/klausborkowski/apploader/internal/api"
	"github.com/klausborkowski/apploader/internal/app"
	"github.com/klausborkowski/apploader/internal/repo"
)

func main() {
	db, err := repo.NewDatabase("postgres://postgres:postgres@localhost:5432/postgres")
	if err != nil {
		log.Fatalf("Ошибка подключения к БД: %v", err)
	}

	app := app.NewApp(db)

	http.HandleFunc("/api/auth", api.MakeAuthHandler(app))
	http.HandleFunc("/api/upload-asset/", api.MakeUploadHandler(app))
	http.HandleFunc("/api/asset/", api.MakeDownloadHandler(app))

	log.Println("Сервер запущен на :8080")
	http.ListenAndServe(":8080", nil)
}
