package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"techpulse/internal/feed"
	"techpulse/internal/model"
)

func Brief(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()

	articles, err := feed.Fetch(ctx, feed.TechCrunchURL, "TechCrunch")
	if err != nil {
		log.Printf("Failed to fetch news: %v", err)
		http.Error(w, "Failed to fetch news", http.StatusBadGateway)
		return
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
