package feed

const (
	TechCrunchURL  = "https://techcrunch.com/feed/"
	ArsTechnicaURL = "https://feeds.arstechnica.com/arstechnica/index"
	TheVergeURL    = "https://www.theverge.com/rss/index.xml"
)

type Source struct {
	Name string
	URL  string
}

var Sources = []Source{
	{Name: "TechCrunch", URL: TechCrunchURL},
	{Name: "Ars Technica", URL: ArsTechnicaURL},
	{Name: "The Verge", URL: TheVergeURL},
}
