package sdk

import (
	"fmt"
	"testing"

	"github.com/kr/pretty"
	"github.com/stretchr/testify/assert"
)

func TestThatYoutubePlaylistsAreGet(t *testing.T) {
	ys := New()

	ps, err := ys.GetPlaylists("@shazam5705")
	assert.Nil(t, err, "error should be nil")

	fmt.Println(pretty.Sprint(ps))
}

func TestThatYoutubePlaylistSongsAreGet(t *testing.T) {
	ys := New()

	ps, err := ys.GetPlaylistItems("https://www.youtube.com/playlist?list=PLepvTsiQW3Wxo5_dKQJvBgWMQ_LnMCiKi")
	assert.Nil(t, err, "error should be nil")

	fmt.Println(pretty.Sprint(ps))
}
