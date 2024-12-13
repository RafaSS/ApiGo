package types

import "time"

type Manga struct {
	ID            string          `json:"id"`
	Type          string          `json:"type"`
	Attributes    MangaAttributes `json:"attributes"`
	Relationships []Relationship  `json:"relationships"`
}

type MangaAttributes struct {
	Title                          LocalizedString   `json:"title"`
	AltTitles                      []LocalizedString `json:"altTitles"`
	Description                    LocalizedString   `json:"description"`
	IsLocked                       bool              `json:"isLocked"`
	Links                          map[string]string `json:"links"`
	OriginalLanguage               string            `json:"originalLanguage"`
	LastVolume                     *string           `json:"lastVolume"`
	LastChapter                    *string           `json:"lastChapter"`
	PublicationDemographic         *string           `json:"publicationDemographic"`
	Status                         string            `json:"status"`
	Year                           *int              `json:"year"`
	ContentRating                  string            `json:"contentRating"`
	ChapterNumbersResetOnNewVolume bool              `json:"chapterNumbersResetOnNewVolume"`
	AvailableTranslatedLanguages   []string          `json:"availableTranslatedLanguages"`
	LatestUploadedChapter          string            `json:"latestUploadedChapter"`
	Tags                           []Tag             `json:"tags"`
	State                          string            `json:"state"`
	Version                        int64             `json:"version"`
	CreatedAt                      time.Time         `json:"createdAt"`
	UpdatedAt                      time.Time         `json:"updatedAt"`
}

type LocalizedString map[string]string

type Tag struct {
	ID            string         `json:"id"`
	Type          string         `json:"type"`
	Attributes    TagAttributes  `json:"attributes"`
	Relationships []Relationship `json:"relationships"`
}

type TagAttributes struct {
	Name        LocalizedString `json:"name"`
	Description LocalizedString `json:"description"`
	Group       string          `json:"group"`
	Version     int64           `json:"version"`
}

type Relationship struct {
	ID         string      `json:"id"`
	Type       string      `json:"type"`
	Related    *string     `json:"related,omitempty"`
	Attributes interface{} `json:"attributes,omitempty"`
}

const (
	Shounen = "shounen"
	Shoujo  = "shoujo"
	Josei   = "josei"
	Seinen  = "seinen"
)

const (
	Completed = "completed"
	Ongoing   = "ongoing"
	Cancelled = "cancelled"
	Hiatus    = "hiatus"
)

const (
	Safe         = "safe"
	Suggestive   = "suggestive"
	Erotica      = "erotica"
	Pornographic = "pornographic"
)

const (
	Content = "content"
	Format  = "format"
	Genre   = "genre"
	Theme   = "theme"
)

const (
	Draft     = "draft"
	Submitted = "submitted"
	Published = "published"
	Rejected  = "rejected"
)

const (
	Monochrome       = "monochrome"
	MainStory        = "main_story"
	AdaptedFrom      = "adapted_from"
	BasedOn          = "based_on"
	Prequel          = "prequel"
	SideStory        = "side_story"
	Doujinshi        = "doujinshi"
	SameFranchise    = "same_franchise"
	SharedUniverse   = "shared_universe"
	Sequel           = "sequel"
	SpinOff          = "spin_off"
	AlternateStory   = "alternate_story"
	AlternateVersion = "alternate_version"
	Preserialization = "preserialization"
	Colored          = "colored"
	Serialization    = "serialization"
)

type MangaRequest struct {
	Title                          LocalizedString   `json:"title"`
	AltTitles                      []LocalizedString `json:"altTitles"`
	Description                    LocalizedString   `json:"description"`
	Authors                        []string          `json:"authors"`
	Artists                        []string          `json:"artists"`
	Links                          map[string]string `json:"links"`
	OriginalLanguage               string            `json:"originalLanguage"`
	LastVolume                     *string           `json:"lastVolume,omitempty"`
	LastChapter                    *string           `json:"lastChapter,omitempty"`
	PublicationDemographic         *string           `json:"publicationDemographic,omitempty"`
	Status                         string            `json:"status"`
	Year                           *int              `json:"year,omitempty"`
	ContentRating                  string            `json:"contentRating"`
	ChapterNumbersResetOnNewVolume bool              `json:"chapterNumbersResetOnNewVolume"`
	Tags                           []string          `json:"tags"`
	PrimaryCover                   *string           `json:"primaryCover,omitempty"`
	Version                        int               `json:"version"`
}

type MangaResponse struct {
	Result   string `json:"result"`
	Response string `json:"response"`
	Data     Manga  `json:"data"`
}

// Add MangaViewModel struct
type MangaViewModel struct {
	*Manga
	AuthorName string
}

// Add MangaListViewModel struct
type MangaListViewModel struct {
	Mangas []*MangaViewModel
}
