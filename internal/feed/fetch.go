package feed

import (
	"context"
	"fmt"

	"github.com/mmcdole/gofeed"

	"techpulse/internal/model"
)

func Fetch(ctx context.Context, url string, source string) ([]model.Article, error) {
	parser := gofeed.NewParser()

	result, err := parser.ParseURLWithContext(url, ctx)
	if err != nil {
		return nil, fmt.Errorf("fetch %s: %w", source, err)
	}

	articles := make([]model.Article, 0, len(result.Items))

	for _, item := range result.Items {
		if item.PublishedParsed == nil {
			continue
		}

		articles = append(articles, model.Article{
			Title:       item.Title,
			Description: cleanDescription(item.Description),
			URL:         item.Link,
			Source:      source,
			PublishedAt: item.PublishedParsed.UTC(),
			Categories:  item.Categories,
		})
	}

	return articles, nil
}
