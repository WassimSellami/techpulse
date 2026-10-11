package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"techpulse/internal/model"
)

func Brief(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	now := time.Now().UTC()

	articles := []model.Article{
		{
			Title:       "ChatGPT 16 released",
			Description: "not yet, don't worry.",
			URL:         "https://example.com/GPT16",
			Source:      "Mock Tech News",
			PublishedAt: now.Add(-time.Hour),
		},
		{
			Title:       "Iphone 25 released",
			Description: "another silly joke.",
			URL:         "https://example.com/iphone25",
			Source:      "Mock Tech News",
			PublishedAt: now.Add(-2 * time.Hour),
		},
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
