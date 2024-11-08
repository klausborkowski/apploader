package app

import (
	"github.com/klausborkowski/apploader/internal/repo"
)

type App struct {
	Repo *repo.Database
}

func NewApp(repo *repo.Database) *App {
	return &App{
		Repo: repo,
	}
}

func (a *App) AuthenticateUser(login, password string) (string, error) {
	return a.Repo.GetUserToken(login, password)
}

func (a *App) UploadAsset(token, name string, data []byte) error {
	return a.Repo.StoreAsset(token, name, data)
}

func (a *App) DownloadAsset(token, name string) ([]byte, error) {
	return a.Repo.GetAsset(token, name)
}
