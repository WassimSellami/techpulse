package brief

import (
	"strings"

	"techpulse/internal/model"
)

func FilterTech(articles []model.Article) []model.Article {
	filtered := make([]model.Article, 0, len(articles))

	for _, article := range articles {
		matched := false

		for _, category := range article.Categories {
			key := strings.ToLower(strings.TrimSpace(category))

			if _, allowed := allowedCategories[key]; allowed {
				matched = true
				break
			}
		}

		if matched {
			filtered = append(filtered, article)
		}
	}

	return filtered
}
