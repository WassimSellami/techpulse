package brief

import (
	"sort"
	"time"

	"techpulse/internal/model"
)

func Select(articles []model.Article, hours int, now time.Time) []model.Article {
	cutoff := now.Add(-time.Duration(hours) * time.Hour)
	selected := make([]model.Article, 0)

	for _, article := range articles {
		if !article.PublishedAt.Before(cutoff) &&
			!article.PublishedAt.After(now) {
			selected = append(selected, article)
		}
	}

	sort.Slice(selected, func(i, j int) bool {
		return selected[i].PublishedAt.After(selected[j].PublishedAt)
	})

	return selected
}
