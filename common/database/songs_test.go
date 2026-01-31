package database

import (
	"context"
	"fmt"
	"testing"

	"github.com/Synchro/common/models"
	"github.com/stretchr/testify/assert"
)

func TestThatSongIsInsertedInDatabase(t *testing.T) {
	cl, err := New("mongodb://admin:abc123@192.168.1.111:27017/")

	assert.Nil(t, err, "error should be nil")
	assert.NotNil(t, cl, "client should not be nil")

	sdc := NewSongsDataController(cl)

	s := models.Song{
		VideoId:      "asdada",
		Name:         "DummySong",
		PlaylistName: "Dummy Playlist",
		Downloaded:   false,
	}

	err = sdc.InsertSong(context.Background(), &s)
	assert.Nil(t, err, "error should be nil")

}

func TestThatSongIsGetWithoutFilter(t *testing.T) {
	cl, err := New("mongodb://admin:abc123@192.168.1.111:27017/")

	assert.Nil(t, err, "error should be nil")
	assert.NotNil(t, cl, "client should not be nil")

	sdc := NewSongsDataController(cl)

	songs, err := sdc.GetSongs(context.Background(), nil)
	if err != nil {
		fmt.Println(err.Error())
	}
	assert.Nil(t, err, "error should be nil")
	fmt.Println(songs)
}

func TestThatSongIsGetWithFilter(t *testing.T) {
	cl, err := New("mongodb://admin:abc123@192.168.1.111:27017/")
	assert.Nil(t, err, "error should be nil")
	assert.NotNil(t, cl, "client should not be nil")

	sdc := NewSongsDataController(cl)

	songs, err := sdc.GetSongs(context.Background(), map[string]any{"videoId": "samam"})
	if err != nil {
		fmt.Println(err.Error())
	}
	assert.Nil(t, err, "error should be nil")
	fmt.Println(songs)
}

func TestThatSongIsUpdated(t *testing.T) {
	cl, err := New("mongodb://admin:abc123@192.168.1.111:27017/")
	assert.Nil(t, err, "error should be nil")
	assert.NotNil(t, cl, "client should not be nil")

	sdc := NewSongsDataController(cl)

	songs, err := sdc.GetSongs(context.Background(), map[string]any{"videoId": "asdada"})
	if err != nil {
		fmt.Println(err.Error())
	}
	assert.Nil(t, err, "error should be nil")

	songs[0].Downloaded = false

	err = sdc.UpdateSong(context.Background(), &songs[0])
	if err != nil {
		fmt.Println(err.Error())
	}
	assert.Nil(t, err, "error should be nil")
	fmt.Println(songs)
}

func TestThatSongSyncStatusIsUpdated(t *testing.T) {
	cl, err := New("mongodb://admin:abc123@192.168.1.111:27017/")
	assert.Nil(t, err, "error should be nil")
	assert.NotNil(t, cl, "client should not be nil")

	sdc := NewSongsDataController(cl)

	err = sdc.SetSyncStatus(context.Background(), "test/S83uQdGqBSY.mp3")
	if err != nil {
		fmt.Println(err.Error())
	}
	assert.Nil(t, err, "error should be nil")

}
