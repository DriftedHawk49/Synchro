package models

type Playlist struct {
	Title string // name of Playlist
	Id    string // Youtube Playlist Id
}

type Song struct {
	VideoId string // VideoId for download
	Title   string // Title/Name of Song
}

type SongsWithPlaylist struct {
	Title string // Name/Title of playlist
	Songs []*Song
}
