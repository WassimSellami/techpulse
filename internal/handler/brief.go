package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"techpulse/internal/brief"
	"techpulse/internal/feed"
	"techpulse/internal/model"
)

func Brief(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	hours := 24

	if value := r.URL.Query().Get("hours"); value != "" {
		parsed, err := strconv.Atoi(value)
		if err != nil || parsed < 1 || parsed > 48 {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{
				"error": "hours must be an integer between 1 and 48",
			})
			return
		}

		hours = parsed
	}

	articles := make([]model.Article, 0)

	for _, source := range feed.Sources {
		ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)

		sourceArticles, err := feed.Fetch(ctx, source.URL, source.Name)
		cancel()

		if err != nil {
			log.Printf("Failed to fetch %s: %v", source.Name, err)
			http.Error(w, "Failed to fetch news", http.StatusBadGateway)
			return
		}

		articles = append(articles, sourceArticles...)
		articles = brief.Select(articles, hours, time.Now().UTC())
	}

	response := struct {
		Count    int             `json:"count"`
		Articles []model.Article `json:"articles"`
	}{
		Count:    len(articles),
		Articles: articles,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
