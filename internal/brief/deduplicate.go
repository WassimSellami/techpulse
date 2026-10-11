package brief

import "techpulse/internal/model"

func Deduplicate(articles []model.Article) []model.Article {
	seen := make(map[string]struct{})
	unique := make([]model.Article, 0, len(articles))

	for _, article := range articles {
		if article.URL == "" {
			unique = append(unique, article)
			continue
		}

		if _, exists := seen[article.URL]; exists {
			continue
		}

		seen[article.URL] = struct{}{}
		unique = append(unique, article)
	}

	return unique
}
