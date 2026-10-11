package brief

import (
	"strings"

	"techpulse/internal/model"
)

func FilterTech(articles []model.Article) []model.Article {
	filtered := make([]model.Article, 0, len(articles))

	for _, article := range articles {
		for _, category := range article.Categories {
			key := strings.ToLower(strings.TrimSpace(category))

			if _, allowed := allowedCategories[key]; allowed {
				filtered = append(filtered, article)
				break
			}
		}
	}

	return filtered
}
