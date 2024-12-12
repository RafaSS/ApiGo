package types

import "time"

type Manga struct {
	Result   string  `json:"result"`
	Response string  `json:"response"`
	Data     []Datum `json:"data"`
	Limit    int64   `json:"limit"`
	Offset   int64   `json:"offset"`
	Total    int64   `json:"total"`
}

type Datum struct {
	ID            string           `json:"id"`
	Type          RelationshipType `json:"type"`
	Attributes    DatumAttributes  `json:"attributes"`
	Relationships []Relationship   `json:"relationships"`
}

type DatumAttributes struct {
	Title                          Title             `json:"title"`
	AltTitles                      []AltTitle        `json:"altTitles"`
	Description                    PurpleDescription `json:"description"`
	IsLocked                       bool              `json:"isLocked"`
	Links                          Links             `json:"links"`
	OriginalLanguage               OriginalLanguage  `json:"originalLanguage"`
	LastVolume                     *string           `json:"lastVolume"`
	LastChapter                    *string           `json:"lastChapter"`
	PublicationDemographic         *string           `json:"publicationDemographic"`
	Status                         Status            `json:"status"`
	Year                           int64             `json:"year"`
	ContentRating                  ContentRating     `json:"contentRating"`
	Tags                           []TagElement      `json:"tags"`
	State                          State             `json:"state"`
	ChapterNumbersResetOnNewVolume bool              `json:"chapterNumbersResetOnNewVolume"`
	CreatedAt                      time.Time         `json:"createdAt"`
	UpdatedAt                      time.Time         `json:"updatedAt"`
	Version                        int64             `json:"version"`
	AvailableTranslatedLanguages   []string          `json:"availableTranslatedLanguages"`
	LatestUploadedChapter          string            `json:"latestUploadedChapter"`
}

type AltTitle struct {
	En   *string `json:"en,omitempty"`
	Tr   *string `json:"tr,omitempty"`
	Ja   *string `json:"ja,omitempty"`
	ZhHk *string `json:"zh-hk,omitempty"`
	EsLa *string `json:"es-la,omitempty"`
	JaRo *string `json:"ja-ro,omitempty"`
	PtBr *string `json:"pt-br,omitempty"`
	Vi   *string `json:"vi,omitempty"`
	ID   *string `json:"id,omitempty"`
	Ko   *string `json:"ko,omitempty"`
	Fr   *string `json:"fr,omitempty"`
	Ru   *string `json:"ru,omitempty"`
	Zh   *string `json:"zh,omitempty"`
	ZhRo *string `json:"zh-ro,omitempty"`
	Uk   *string `json:"uk,omitempty"`
}

type PurpleDescription struct {
	En   string  `json:"en"`
	EsLa *string `json:"es-la,omitempty"`
	ID   *string `json:"id,omitempty"`
	PtBr *string `json:"pt-br,omitempty"`
	Ko   *string `json:"ko,omitempty"`
	Fr   *string `json:"fr,omitempty"`
	Zh   *string `json:"zh,omitempty"`
	Tr   *string `json:"tr,omitempty"`
	Ja   *string `json:"ja,omitempty"`
}

type Links struct {
	Al    *string `json:"al,omitempty"`
	Ap    *string `json:"ap,omitempty"`
	BW    *string `json:"bw,omitempty"`
	Kt    *string `json:"kt,omitempty"`
	Mu    *string `json:"mu,omitempty"`
	Amz   *string `json:"amz,omitempty"`
	Ebj   *string `json:"ebj,omitempty"`
	Mal   *string `json:"mal,omitempty"`
	Raw   *string `json:"raw,omitempty"`
	Engtl *string `json:"engtl,omitempty"`
	Nu    *string `json:"nu,omitempty"`
	Cdj   *string `json:"cdj,omitempty"`
}

type TagElement struct {
	ID            string        `json:"id"`
	Type          TagType       `json:"type"`
	Attributes    TagAttributes `json:"attributes"`
	Relationships []interface{} `json:"relationships"`
}

type TagAttributes struct {
	Name        Title             `json:"name"`
	Description FluffyDescription `json:"description"`
	Group       Group             `json:"group"`
	Version     int64             `json:"version"`
}

type FluffyDescription struct {
}

type Title struct {
	En string `json:"en"`
}

type Relationship struct {
	ID      string           `json:"id"`
	Type    RelationshipType `json:"type"`
	Related *string          `json:"related,omitempty"`
}

type ContentRating string

const (
	Erotica    ContentRating = "erotica"
	Safe       ContentRating = "safe"
	Suggestive ContentRating = "suggestive"
)

type OriginalLanguage string

const (
	Ja OriginalLanguage = "ja"
	Ko OriginalLanguage = "ko"
	Zh OriginalLanguage = "zh"
)

type State string

const (
	Published State = "published"
)

type Status string

const (
	Completed Status = "completed"
	Ongoing   Status = "ongoing"
)

type Group string

const (
	Content Group = "content"
	Format  Group = "format"
	Genre   Group = "genre"
	Theme   Group = "theme"
)

type TagType string

const (
	Tag TagType = "tag"
)

type RelationshipType string

const (
	Artist    RelationshipType = "artist"
	Author    RelationshipType = "author"
	CoverArt  RelationshipType = "cover_art"
	Creator   RelationshipType = "creator"
	TypeManga RelationshipType = "manga"
)
