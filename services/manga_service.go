package services

import (
	"ApiGo/types"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// MangaServiceImpl is the concrete implementation of MangaService.
type MangaServiceImpl struct{}

// GetMangaList fetches the list of mangas for a given page.
func (s *MangaServiceImpl) GetMangaList(page int) (*types.MangaListViewModel, error) {
	apiURL := fmt.Sprintf("https://api.mangadex.org/manga?page=%d", page)

	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get manga list: %s", resp.Status)
	}

	var result types.MangaList
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	viewModel := &types.MangaListViewModel{
		Mangas:      make([]*types.MangaViewModel, len(result.Data)),
		CurrentPage: page,
	}
	for i, manga := range result.Data {
		viewModel.Mangas[i] = &types.MangaViewModel{
			Manga:      &manga,
			AuthorName: s.GetAuthorName(manga.Relationships),
		}
	}

	return viewModel, nil
}

// GetMangaListWithTitle fetches the list of mangas based on the title.
func (s *MangaServiceImpl) GetMangaListWithTitle(title string) (*types.MangaListViewModel, error) {
	apiURL := fmt.Sprintf("https://api.mangadex.org/manga?title=%s", url.QueryEscape(title))

	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get manga list: %s", resp.Status)
	}

	var result types.MangaList
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	viewModel := &types.MangaListViewModel{
		Mangas:      make([]*types.MangaViewModel, len(result.Data)),
		CurrentPage: 1, // Assuming search returns the first page
	}
	for i, manga := range result.Data {
		viewModel.Mangas[i] = &types.MangaViewModel{
			Manga:      &manga,
			AuthorName: s.GetAuthorName(manga.Relationships),
		}
	}

	return viewModel, nil
}

// GetMangaDetails fetches the details of a specific manga by ID.
func (s *MangaServiceImpl) GetMangaDetails(id string) (*types.MangaViewModel, error) {
	apiURL := fmt.Sprintf("https://api.mangadex.org/manga/%s", id)

	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get manga details: %s", resp.Status)
	}

	var manga types.Manga
	if err := json.NewDecoder(resp.Body).Decode(&manga); err != nil {
		return nil, err
	}

	viewModel := &types.MangaViewModel{
		Manga:      &manga,
		AuthorName: s.GetAuthorName(manga.Relationships),
	}

	return viewModel, nil
}

// Helper method to get the author's name.
func (s *MangaServiceImpl) GetAuthorName(relationships []types.Relationship) string {
	for _, rel := range relationships {
		if rel.Type == "author" {
			return rel.ID
		}
	}
	return "Unknown Author"
}
