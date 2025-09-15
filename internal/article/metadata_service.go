package article

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/chiquitav2/journalful/internal/db"
	"io"
	"log/slog"
	"net/http"
	"time"
)

// CrossRef API response structs
type CrossRefResponse struct {
	Message CrossRefMessage `json:"message"`
}

type CrossRefMessage struct {
	DOI             string           `json:"DOI"`
	Title           []string         `json:"title"`
	Author          []CrossRefAuthor `json:"author"`
	PublishedOnline CrossRefDate     `json:"published"`
	URL             string           `json:"URL"`
	Abstract        string           `json:"abstract"`
	ContainerTitle  []string         `json:"container-title"`
}

type CrossRefAuthor struct {
	Given  string `json:"given"`
	Family string `json:"family"`
}

type CrossRefDate struct {
	DateParts [][]int `json:"date-parts"`
}

type MetadataService struct {
	client *http.Client
}

func NewMetadataService() *MetadataService {
	return &MetadataService{
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (s *MetadataService) FetchAndPrepareArticle(doi string) (*db.CreateArticleParams, []string, error) {
	meta, err := s.fetchArticleMetadataFromDOI(doi)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to fetch article metadata from DOI: %w", err)
	}

	var articleTitle string
	if len(meta.Title) > 0 {
		articleTitle = meta.Title[0]
	}

	articleAbstract := sql.NullString{String: meta.Abstract, Valid: meta.Abstract != ""}

	var articlePublicationYear sql.NullInt32
	if len(meta.PublishedOnline.DateParts) > 0 && len(meta.PublishedOnline.DateParts[0]) > 0 {
		articlePublicationYear = sql.NullInt32{Int32: int32(meta.PublishedOnline.DateParts[0][0]), Valid: true}
	}

	var articleJournalName sql.NullString
	if len(meta.ContainerTitle) > 0 {
		articleJournalName = sql.NullString{String: meta.ContainerTitle[0], Valid: meta.ContainerTitle[0] != ""}
	}

	var normalizedAuthors []string
	for _, author := range meta.Author {
		fullName := fmt.Sprintf("%s %s", author.Given, author.Family)
		normalizedAuthors = append(normalizedAuthors, fullName)
	}

	return &db.CreateArticleParams{
		Title:           articleTitle,
		Doi:             doi,
		Url:             sql.NullString{String: meta.URL, Valid: meta.URL != ""},
		Abstract:        articleAbstract,
		PublicationYear: articlePublicationYear,
		JournalName:     articleJournalName,
	}, normalizedAuthors, nil
}

func (s *MetadataService) fetchArticleMetadataFromDOI(doi string) (*CrossRefMessage, error) {
	url := fmt.Sprintf("https://api.crossref.org/v1/works/%s", doi)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP request: %w", err)
	}

	req.Header.Set("User-Agent", "Journalful/1.0")

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make HTTP request to CrossRef API: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		slog.Error("CrossRef API returned non-OK status", "status", resp.Status, "doi", doi)
		return nil, fmt.Errorf("CrossRef API returned non-OK status: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read CrossRef API response body: %w", err)
	}

	var crossRefResponse CrossRefResponse
	err = json.Unmarshal(body, &crossRefResponse)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal CrossRef API response: %w", err)
	}

	slog.Info("Successfully fetched metadata from CrossRef", "doi", doi)
	return &crossRefResponse.Message, nil
}
