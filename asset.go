package main

import (
	"bitbucket.org/ffxblue/api-content/lib/data"
	"time"
)

// Asset model
type Asset struct {
	Ads           AssetAds              `json:"ads" description:"Ads metadata associated with the asset."`
	AssetType     string                `json:"assetType" description:"Schema.org definition of the asset type."`
	Categories    []string              `json:"categories" description:"Categories associated with the asset."`
	CategoryID    int                   `json:"categoryId,omitempty" description:"Identifier of the category associated with the asset."`
	Data          AssetData             `json:"asset" description:"Data associated with the type of asset."`
	Dates         AssetDates            `json:"dates" description:"Dates associated with the asset."`
	EditingState  string                `json:"editingState" description:"Editing state of the asset (i.e. draft, ready, etc)."`
	ID            string                `json:"id" description:"Identifier of the asset."`
	Images        map[string]AssetImage `json:"featuredImages,omitempty" description:"Images associated with the asset."`
	Label         string                `json:"label,omitempty" description:"Label associated with the asset (i.e. exclusive, etc)."`
	Participants  *AssetParticipants    `json:"participants,omitempty" description:"Participants associated with the asset (i.e. authors, creator, etc)."`
	PromotedBrand *AssetPromotedBrand   `json:"promotedBrand,omitempty" description:"The brand to be identified with this asset."`
	PublicState   string                `json:"publicState" description:"Public visibility state of the asset (i.e. published, retracted, etc)."`
	Resources     []AssetResource       `json:"resources,omitempty" description:"Resources associated with the asset (i.e. additional widget configurations)."`
	SEO           *AssetSEO             `json:"seo,omitempty" description:"SEO associated with the asset."`
	Social        AssetSocial           `json:"social" description:"Social information of the asset."`
	SourceCMS     AssetSourceCMS        `json:"sourceCms" description:"Source CMS of the asset."`
	Sources       []AssetSource         `json:"sources,omitempty" description:"Sources of the asset."`
	Sponsor       *AssetSponsor         `json:"sponsor,omitempty" description:"Sponsor of the asset."`
	Tags          *AssetTags            `json:"tags,omitempty" description:"Tags associated with the asset."`
	URLs          AssetURLs             `json:"urls" description:"URLs associated with the asset."`
	Version       AssetVersion          `json:"version" description:"Version number of the asset."`
}

// AssetAds model
type AssetAds struct {
	ExclusionTopics []string `json:"exclusionTopics,omitempty" description:"Categories of ads that are excluded"`
	Suppress        bool     `json:"suppress" description:"Whether ads should be suppressed."`
}

// AssetAuthor model
type AssetAuthor struct {
	Bio    string                `json:"bio,omitempty" description:"Description about the author."`
	ID     string                `json:"id" description:"Identifier of the author."`
	Images map[string]AssetImage `json:"featuredImages,omitempty" description:"Author images"`
	Name   string                `json:"name" description:"Name of the author."`
	Title  string                `json:"title,omitempty" description:"Job title of the author."`
}

// AssetData model
type AssetData struct {
	About            string                     `json:"about,omitempty" description:"Description about the asset (i.e. write off used outside of the asset display)."`
	AllowComments    bool                       `json:"allowComments" description:"Whether or not comments are allowed on the asset."`
	Body             string                     `json:"body,omitempty" description:"Body content of the asset."`
	BodyPlaceholders map[string]BodyPlaceholder `json:"bodyPlaceholders,omitempty" description:"Resources associated with the asset (i.e. content to be substituted inside the asset body)."`
	Byline           string                     `json:"byline,omitempty" description:"Free text describing the author of the asset."`
	CloseComments    bool                       `json:"closeComments" description:"Whether or not comments are closed on the asset."`
	Headlines        AssetHeadlines             `json:"headlines" description:"Headline variations of the asset."`
	Intro            string                     `json:"intro,omitempty" description:"Introduction to the asset (i.e. write off)."`

	// Gallery Type fields
	Images []AssetImageData `json:"images,omitempty" description:"Images contained within the gallery"`

	// Live Article type fields
	IsLive            bool       `json:"isLive,omitempty" description:"Indicates if the live article is still live"`
	LastPostPublished *time.Time `json:"lastPostPublished,omitempty" description:"The time of the last published post for the live article"`

	// Video type fields
	Duration      int    `json:"duration,omitempty" description:"Video duration"`
	Geoblocked    bool   `json:"geoblocked,omitempty" description:"Indicates if the video is geo blocked"`
	LiveStreamURL string `json:"liveStreamURL,omitempty" description:"Live stream URL"`
	Producer      string `json:"producer,omitempty" description:"Producer of the video"`
	RenderMode    string `json:"renderMode,omitempty" description:"Video render mode (i.e. progressive, live, stream)"`
	VideoType     string `json:"videoType,omitempty" description:"Video type (i.e. standard, ad)"`
}

// AssetDates model
type AssetDates struct {
	Created        time.Time  `json:"created" description:"Date the asset was created."`
	FirstPublished *time.Time `json:"firstPublished,omitempty" description:"Date the asset was first published."`
	Imported       time.Time  `json:"imported,omitempty" description:"Date the asset was imported."`
	Modified       *time.Time `json:"modified,omitempty" description:"Date the asset was last modified."`
	Published      *time.Time `json:"published,omitempty" description:"Date the asset was last published."`
	Saved          *time.Time `json:"saved,omitempty" description:"Date the asset was last saved."`
	TimeToTakeDown *time.Time `json:"timeToTakeDown,omitempty" description:"Date the asset is due to be removed from publication."`
}

// AssetHeadlines model
type AssetHeadlines struct {
	Headline string `json:"headline" description:"Headline of the asset."`
	Medium   string `json:"medium,omitempty" description:"Medium headline of the asset (55 chars or less)."`
}

// AssetImage model
type AssetImage struct {
	Data AssetImageData `json:"data" description:"Data representing the image."`
}

// AssetImageData defines metadata for an image
type AssetImageData struct {
	AltText   string   `json:"altText,omitempty" description:"Alternate text description of the image"`
	Animated  *bool    `json:"animated,omitempty" description:"Whether the image is animated"`
	Aspect    *float64 `json:"aspect,omitempty" description:"Aspect ratio of the image"`
	AutoCrop  bool     `json:"autoCrop,omitempty" description:"Identifies whether the crop is applied via automated means"`
	Caption   string   `json:"caption,omitempty" description:"Caption describing the image"`
	Credit    string   `json:"credit,omitempty" description:"Person or organisation which is credited for the image"`
	CropWidth *uint    `json:"cropWidth,omitempty" description:"Width of the crop"`
	Filename  string   `json:"fileName,omitempty" description:"Filename of the image"`
	ID        string   `json:"id,omitempty" description:"Image ID"`
	MimeType  *string  `json:"mimeType,omitempty" description:"Detected mimeType of the image file"`
	OffsetX   *int     `json:"offsetX,omitempty" description:"X offset of the crop"`
	OffsetY   *int     `json:"offsetY,omitempty" description:"Y offset of the crop"`
	Source    string   `json:"source,omitempty" description:"Source of the image"`
	Zoom      *float64 `json:"zoom,omitempty" description:"Zoom factor to apply for renditions"`
}

// AssetParticipants model
type AssetParticipants struct {
	Authors []*AssetAuthor `json:"authors,omitempty" description:"Authors of the asset."`
}

// AssetPromotedBrand promoted brand details
type AssetPromotedBrand struct {
	ID    string `json:"id" description:"The short name of the promoted brand."`
	Label string `json:"label,omitempty" description:"The label of the promoted brand"`
}

// AssetResource model
type AssetResource struct {
	Data AssetResourceData `json:"data" description:"Data representing the resource."`
	Type string            `json:"type" description:"Type of resource."`
}

// AssetSEO model
type AssetSEO struct {
	Description  string `json:"description,omitempty" description:"SEO description of the asset."`
	DoNotIndex   bool   `json:"doNotIndex" description:"Whether the content must be marked as not to be indexed by search engines"`
	Keywords     string `json:"keywords,omitempty" description:"SEO Keywords associated with the asset."`
	NewsKeywords string `json:"newsKeywords,omitempty" description:"SEO news keywords of the asset"`
	Standout     bool   `json:"standout" description:"Indicates that the content should be marked as a 'standout' e.g. for Google News"`
	Title        string `json:"title,omitempty" description:"SEO title of the asset."`
}

// AssetSocial model
type AssetSocial struct {
	DoNotShare   bool                 `json:"doNotShare" description:"Whether the asset can be shared."`
	FacebookData *AssetSocialFacebook `json:"facebook,omitempty" description:"Facebook social preview"`
}

// AssetSocialFacebook model
type AssetSocialFacebook struct {
	Images map[string]AssetImage `json:"images,omitempty"`
}

// AssetSourceCMS model
type AssetSourceCMS struct {
	Type string `json:"cmsType" description:"Type of CMS that produced the asset."`
}

// AssetSponsor model
type AssetSponsor struct {
	Logos map[string]AssetImage `json:"logos,omitempty" description:"Images for asset's sponsor"`
	Name  string                `json:"name" description:"Name of the asset's sponsor."`
	Type  string                `json:"type,omitempty" description:"The type of sponsor"`
	URL   string                `json:"url,omitempty" description:"Optional URL to an external sponsor page."`
}

// AssetTags model
type AssetTags struct {
	Primary   *TagPreview  `json:"primary,omitempty" description:"Primary tag associated with the asset."`
	Secondary []TagPreview `json:"secondary,omitempty" description:"Secondary tags associated with the asset."`
}

// AssetURL model
type AssetURL struct {
	Brand    string `json:"brand,omitempty" description:"Brand of the url."`
	External string `json:"external,omitempty" description:"External canonical url."`
	Path     string `json:"path,omitempty" description:"Path of the url."`
}

// AssetURLs model
type AssetURLs struct {
	Canonical *AssetURL           `json:"canonical,omitempty" description:"Canonical URL of the asset."`
	External  string              `json:"external,omitempty" description:"External URL for assets with type 'url'"`
	Published map[string]AssetURL `json:"published,omitempty" description:"Published URLs of the asset."`
	WebSlug   string              `json:"webslug,omitempty" description:"Slug portion of the asset URL."`
}

// AssetVersion model
type AssetVersion struct {
	Internal  int `json:"internal,omitempty" description:"Version of the asset."`
	SourceCMS int `json:"sourceCms" description:"Version of the asset in the source CMS."`
}

// BodyPlaceholder model
type BodyPlaceholder struct {
	Data map[string]interface{} `json:"data" description:"Data representing the placeholder."`
	Type string                 `json:"type" description:"Type of placeholder."`
}

// AssetResourceData defines behavior which is common to all asset resources data
type AssetResourceData interface {
}

// AssetResourceScoreboard defines metadata for 'scoreboard' resource data
type AssetResourceScoreboard struct {
	GameID  string `json:"gameID,omitempty"`
	MatchID string `json:"matchID,omitempty"`
	Title   string `json:"title,omitempty"`
	Type    string `json:"type"`
}

// AssetResourceDataTalkingPointsBase defines metadata that all 'talking points' resource data includes
type AssetResourceDataTalkingPointsBase struct {
	Title string `json:"title,omitempty"`
	Type  string `json:"type"`
}

// AssetTalkingPointsFact model
type AssetTalkingPointsFact struct {
	AssetResourceDataTalkingPointsBase
	Items []AssetTalkingPointsFactItem `json:"items"`
}

// AssetTalkingPointsFactItem model
type AssetTalkingPointsFactItem struct {
	Quantity string `json:"quantity"`
	Text     string `json:"text"`
}

// AssetTalkingPointsList model
type AssetTalkingPointsList struct {
	AssetResourceDataTalkingPointsBase
	Items []AssetTalkingPointsListItem `json:"items"`
}

// AssetTalkingPointsListItem model
type AssetTalkingPointsListItem struct {
	Text string `json:"text"`
}

// AssetTalkingPointsText model
type AssetTalkingPointsText struct {
	AssetResourceDataTalkingPointsBase
	Text string `json:"text"`
}

// AssetSource model to hold source data
type AssetSource struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// FromDataAssetImageData returns a pointer to an
// AssetImageData based off a data.AssetImageData
func FromDataAssetImageData(d data.AssetImageData) *AssetImageData {
	return &AssetImageData{
		AltText:   d.AltText,
		Animated:  d.Animated,
		Aspect:    d.Aspect,
		AutoCrop:  d.AutoCrop,
		Caption:   d.Caption,
		Credit:    d.Credit,
		CropWidth: d.CropWidth,
		Filename:  d.Filename,
		ID:        d.ID,
		MimeType:  d.MimeType,
		OffsetX:   d.OffsetX,
		OffsetY:   d.OffsetY,
		Source:    d.Source,
		Zoom:      d.Zoom,
	}
}
