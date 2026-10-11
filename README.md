# TechPulse

A Go backend that collects technology news from RSS feeds and returns a daily brief in JSON or Markdown.

## Features

- Fetches TechCrunch, Ars Technica, and The Verge on demand
- Filters articles using a technology category allowlist
- Cleans HTML from article descriptions
- Filters by publication time and sorts newest first
- Removes duplicate URLs
- Supports configurable result limits

Category filtering relies on publisher tags and may exclude relevant articles.

## Run

Requires Go.

```bash
go mod download
go run ./cmd/api
```

The server listens on `http://localhost:8080`.

## API

```http
GET /brief
```

| Parameter | Default | Accepted values |
|-----------|---------|-----------------|
| `hours` | `24` | Integer from 1–48 |
| `limit` | `20` | Integer from 1–100 |
| `format` | `json` | `json` or `markdown` |

### JSON brief

```bash
curl "http://localhost:8080/brief?hours=24&limit=10"
```

### Markdown brief

```bash
curl "http://localhost:8080/brief?hours=24&limit=10&format=markdown"
```

### Save a brief

```bash
curl -s "http://localhost:8080/brief?format=markdown" -o brief.md
```

Invalid parameters return HTTP 400. A feed-fetch failure returns HTTP 502.

## Project structure

```text
cmd/api/           Application entry point
internal/handler/  HTTP request handling
internal/feed/     RSS fetching, source URLs, and description cleanup
internal/model/    Article model
internal/brief/    Filtering, sorting, deduplication, and Markdown output
```