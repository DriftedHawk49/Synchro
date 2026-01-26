package sdk

import (
	"fmt"
	"log/slog"
	"testing"

	"github.com/kr/pretty"
	"github.com/stretchr/testify/assert"
)

func TestThatYoutubePlaylistsAreGet(t *testing.T) {
	ys, err := New("dummyApiKey", slog.Default())
	assert.Nil(t, err, "error should be nil")

	ps, err := ys.GetPlaylists("channel id")
	assert.Nil(t, err, "error should be nil")

	fmt.Println(pretty.Sprint(ps))
}

func TestThatYoutubePlaylistSongsAreGet(t *testing.T) {
	ys, err := New("dummy API key", slog.Default())
	assert.Nil(t, err, "error should be nil")

	ps, err := ys.GetPlaylistItems("playlistId")
	assert.Nil(t, err, "error should be nil")

	fmt.Println(pretty.Sprint(ps))
}
