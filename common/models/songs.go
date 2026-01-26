package models

import (
	"time"
)

type Song struct {
	VideoId        string    `bson:"videoId,omitempty"`        // VideoId of a youtube video
	Name           string    `bson:"title,omitempty"`          // Name/Title of song
	PlaylistName   string    `bson:"playlistTitle,omitempty"`  // title of playlist song belongs to
	CreatedAt      time.Time `bson:"createdAt,omitempty"`      // first sync time
	UpdatedAt      time.Time `bson:"updatedAt,omitempty"`      // latest sync time
	LocationOnDisk string    `bson:"locationOnDisk,omitempty"` // location on disk if the song is downloaded
	Downloaded     bool      `bson:"downloaded"`               // Whether the song has been downloaded
	Synced         bool      `bson:"synced"`                   // whether the song is synced by Agent
}
