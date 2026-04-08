package models

import "encoding/json"

type PlaylistWithSong struct {
}

type Song struct {
	Title   string
	VideoId string
}

type Version struct {
	Version        string  `json:"version"`
	CurrentGitHead *string `json:"current_git_head"`
	ReleaseGitHead *string `json:"release_git_head"`
	Repository     string  `json:"repository"`
}

type Thumbnail struct {
	URL        string  `json:"url"`
	Height     int     `json:"height"`
	Width      int     `json:"width"`
	ID         *string `json:"id,omitempty"`
	Resolution *string `json:"resolution,omitempty"`
}

type Entry struct {
	Type              string      `json:"_type"`
	IeKey             string      `json:"ie_key"`
	ID                string      `json:"id"`
	URL               string      `json:"url"`
	Title             string      `json:"title"`
	Description       *string     `json:"description"`
	Duration          int         `json:"duration"`
	ChannelID         string      `json:"channel_id"`
	Channel           string      `json:"channel"`
	ChannelURL        string      `json:"channel_url"`
	Uploader          string      `json:"uploader"`
	UploaderID        *string     `json:"uploader_id"`
	UploaderURL       *string     `json:"uploader_url"`
	Thumbnails        []Thumbnail `json:"thumbnails"`
	Timestamp         *int64      `json:"timestamp"`
	ReleaseTimestamp  *int64      `json:"release_timestamp"`
	Availability      *string     `json:"availability"`
	ViewCount         int64       `json:"view_count"`
	LiveStatus        *string     `json:"live_status"`
	ChannelIsVerified *bool       `json:"channel_is_verified"`
	XForwardedForIP   *string     `json:"__x_forwarded_for_ip"`
}

type Playlist struct {
	ID                   string         `json:"id"`
	Title                string         `json:"title"`
	Availability         *string        `json:"availability"`
	ChannelFollowerCount *int64         `json:"channel_follower_count"`
	Description          string         `json:"description"`
	Tags                 []string       `json:"tags"`
	Thumbnails           []Thumbnail    `json:"thumbnails"`
	ModifiedDate         string         `json:"modified_date"`
	ViewCount            *int64         `json:"view_count"`
	PlaylistCount        int            `json:"playlist_count"`
	Channel              string         `json:"channel"`
	ChannelID            string         `json:"channel_id"`
	UploaderID           string         `json:"uploader_id"`
	Uploader             string         `json:"uploader"`
	ChannelURL           string         `json:"channel_url"`
	UploaderURL          string         `json:"uploader_url"`
	Type                 string         `json:"_type"`
	Entries              []Entry        `json:"entries"`
	ExtractorKey         string         `json:"extractor_key"`
	Extractor            string         `json:"extractor"`
	WebpageURL           string         `json:"webpage_url"`
	OriginalURL          string         `json:"original_url"`
	WebpageURLBasename   string         `json:"webpage_url_basename"`
	WebpageURLDomain     string         `json:"webpage_url_domain"`
	ReleaseYear          *int           `json:"release_year"`
	Epoch                int64          `json:"epoch"`
	FilesToMove          map[string]any `json:"__files_to_move"`
	Version              Version        `json:"_version"`
}

type PlaylistEntry struct {
	Type            string      `json:"_type"`
	IeKey           string      `json:"ie_key"`
	ID              string      `json:"id"`
	URL             string      `json:"url"`
	Title           string      `json:"title"`
	Thumbnails      []Thumbnail `json:"thumbnails"`
	Duration        *int        `json:"duration"`
	Timestamp       *int64      `json:"timestamp"`
	XForwardedForIP *string     `json:"__x_forwarded_for_ip"`
}

type Channel struct {
	ID                   string           `json:"id"`
	Channel              string           `json:"channel"`
	ChannelID            string           `json:"channel_id"`
	Title                string           `json:"title"`
	Availability         *string          `json:"availability"`
	ChannelFollowerCount *int64           `json:"channel_follower_count"`
	Description          string           `json:"description"`
	Tags                 []string         `json:"tags"`
	Thumbnails           []Thumbnail      `json:"thumbnails"`
	UploaderID           string           `json:"uploader_id"`
	UploaderURL          string           `json:"uploader_url"`
	ModifiedDate         *string          `json:"modified_date"`
	ViewCount            *int64           `json:"view_count"`
	PlaylistCount        int              `json:"playlist_count"`
	Uploader             string           `json:"uploader"`
	ChannelURL           string           `json:"channel_url"`
	Type                 string           `json:"_type"`
	Entries              []*PlaylistEntry `json:"entries"`
	ExtractorKey         string           `json:"extractor_key"`
	Extractor            string           `json:"extractor"`
	WebpageURL           string           `json:"webpage_url"`
	OriginalURL          string           `json:"original_url"`
	WebpageURLBasename   string           `json:"webpage_url_basename"`
	WebpageURLDomain     string           `json:"webpage_url_domain"`
	ReleaseYear          *int             `json:"release_year"`
	Epoch                int64            `json:"epoch"`
	FilesToMove          map[string]any   `json:"__files_to_move"`
	Version              Version          `json:"_version"`
}

func ParseChannel(data []byte) (*Channel, error) {
	var c Channel
	if err := json.Unmarshal(data, &c); err != nil {
		return nil, err
	}
	return &c, nil
}

func ParsePlaylist(data []byte) (*Playlist, error) {
	var p Playlist
	if err := json.Unmarshal(data, &p); err != nil {
		return nil, err
	}
	return &p, nil
}
