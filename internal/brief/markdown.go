package brief

import (
	"fmt"
	"strings"

	"techpulse/internal/model"
)

func Markdown(articles []model.Article, hours int) string {
	var output strings.Builder

	fmt.Fprintf(&output, "# TechPulse — Last %d Hours\n\n", hours)

	if len(articles) == 0 {
		output.WriteString("No articles found in this time window.\n")
		return output.String()
	}

	for _, article := range articles {
		fmt.Fprintf(&output, "## %s\n\n", article.Title)
		fmt.Fprintf(&output, "%s\n\n", article.Description)
		fmt.Fprintf(&output, "Source: %s  \n", article.Source)
		fmt.Fprintf(
			&output,
			"Published: %s  \n",
			article.PublishedAt.UTC().Format("2006-01-02 15:04 MST"),
		)
		fmt.Fprintf(&output, "Read more: %s\n\n---\n\n", article.URL)
	}

	return output.String()
}
