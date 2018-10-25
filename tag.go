package main

import "time"

// TagPreview model
type TagPreview struct {
	Context     string         `json:"context" description:"Context of the tag (i.e. person, place, thing, etc)."`
	Description string         `json:"description,omitempty" description:"Description of the tag."`
	DisplayName string         `json:"displayName" description:"Display name of the tag."`
	Name        string         `json:"name" description:"Name of the tag."`
	URLs        TagPreviewURLs `json:"urls" description:"URLs associated with the tag."`
}

// TagPreviewURL model
type TagPreviewURL struct {
	Brand string `json:"brand,omitempty" description:"Brand of the url."`
	Path  string `json:"path,omitempty" description:"Path of the url."`
}

// TagPreviewURLs model
type TagPreviewURLs struct {
	Canonical *TagPreviewURL           `json:"canonical,omitempty" description:"Canonical URL of the tag."`
	External  string                   `json:"external,omitempty" description:"External URL of the tag."`
	Published map[string]TagPreviewURL `json:"published,omitempty" description:"Published URLs of the tag."`
}

// TagTheme model
type TagTheme struct {
	ID          string `json:"id" description:"ID of the tag theme."`
	Description string `json:"description,omitempty" description:"Description of the tag theme."`
	Name        string `json:"name" description:"Name of the theme."`

	Dates TagDates `json:"dates" description:"dates associated with tag theme"`
}

// TagDates defines times stamps associated with the tag scheme, context and tag itself
type TagDates struct {
	Created  time.Time  `json:"created" description:"Creation date"`
	Imported time.Time  `json:"imported" description:"Imported date"`
	Modified *time.Time `json:"modified,omitempty" description:"Last modification date"`
}

// TagThemes model
type TagThemes struct {
	TagThemes []*TagTheme `json:"tagthemes" description:"Tag Themes"`
	Total     int         `json:"total" description:"Number of Tag Themes found"`
}

// TagContext model
type TagContext struct {
	ID          string `json:"id" description:"ID of the tag context."`
	Description string `json:"description,omitempty" description:"Description of the tag context."`
	Name        string `json:"name" description:"Name of the context."`

	Dates TagDates `json:"dates" description:"dates associated with tag context"`
}

// TagContexts model
type TagContexts struct {
	TagContexts []*TagContext `json:"tagcontexts" description:"Tag context"`
	Total       int           `json:"total" description:"Number of Tag Contexts found"`
}

// Tag model
type Tag struct {
	ContextID   string   `json:"contextID,omitempty" description:"Context ID of the tag"`
	ContextName string   `json:"contextName,omitempty" description:"Context name of the tag"`
	Dates       TagDates `json:"dates,omitempty" description:"dates associated with tag"`
	Description string   `json:"description,omitempty" description:"Description of the tag."`
	DisplayName string   `json:"displayName" description:"Display name of the tag."`
	ID          string   `json:"id,omitempty" description:"ID of the tag"`
	Name        string   `json:"name" description:"Name of the tag."`
	Slug        string   `json:"-" description:"Slug portion of the tag url."`
	Themes      []string `json:"themes,omitempty" description:"Themes associated with the tag."`
	Visible     bool     `json:"visible" description:"Whether or not the tag should be visible to render layers."`
}

// Tags model
type Tags struct {
	Page  int    `json:"page" description:"Page number"`
	Tags  []*Tag `json:"tags" description:"Tags"`
	Total int    `json:"total" description:"Number of Tags found"`
}
