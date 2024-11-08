package api

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/klausborkowski/apploader/internal/app"
)

type AuthRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type AuthResponse struct {
	Token string `json:"token"`
}

type UploadResponse struct {
	Status string `json:"status"`
}

func MakeAuthHandler(appInstance *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req AuthRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, `{"error": "invalid request body"}`, http.StatusBadRequest)
			return
		}

		token, err := appInstance.AuthenticateUser(req.Login, req.Password)
		if err != nil {
			http.Error(w, `{"error": "invalid login/password"}`, http.StatusUnauthorized)
			return
		}

		response := AuthResponse{Token: token}
		json.NewEncoder(w).Encode(response)
	}
}

func MakeUploadHandler(appInstance *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer ")
		name := strings.TrimPrefix(r.URL.Path, "/api/upload-asset/")
		data, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, `{"error": "failed to read data"}`, http.StatusBadRequest)
			return
		}

		if err := appInstance.UploadAsset(token, name, data); err != nil {
			http.Error(w, `{"error": "upload failed"}`, http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(UploadResponse{Status: "ok"})
	}
}

func MakeDownloadHandler(appInstance *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer ")
		name := strings.TrimPrefix(r.URL.Path, "/api/asset/")

		data, err := appInstance.DownloadAsset(token, name)
		if err != nil {
			http.Error(w, `{"error": "asset not found"}`, http.StatusNotFound)
			return
		}

		w.Write(data)
	}
}
