package http

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"net/url"
	"repo-stat/api/internal/dto"
	"repo-stat/api/internal/usecase"
)

func NewGetRepositoryInfoHandler(log *slog.Logger, getRepositoryInfo *usecase.GetRepositoryInfo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Info("handling get repository info request", "method", r.Method, "url", r.URL.Path)

		repoURL := r.URL.Query().Get("url")
		if repoURL == "" {
			WriteErrorResponse(w, http.StatusBadRequest, "missing 'url' query parameter")
			return
		}

		parsedURL, err := url.Parse(repoURL)
		if err != nil || parsedURL.Host != "github.com" {
			WriteErrorResponse(w, http.StatusBadRequest, "invalid repository URL")
			return
		}

		info, err := getRepositoryInfo.Execute(r.Context(), repoURL)
		if err != nil {
			log.Error("failed to get repository info", "error", err)
			WriteErrorResponse(w, http.StatusInternalServerError, "failed to get repository info")
			return
		}

		dtoResponse := dto.RepositoryInfoResponse{
			FullName:    info.FullName,
			Description: info.Description,
			Stars:       info.Stars,
			Forks:       info.Forks,
			CreatedAt:   info.CreatedAt,
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(dtoResponse); err != nil {
			log.Error("failed to encode response", "error", err)
			WriteErrorResponse(w, http.StatusInternalServerError, "failed to encode response")
			return
		}
	}
}
