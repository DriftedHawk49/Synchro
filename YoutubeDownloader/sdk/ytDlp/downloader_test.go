package ytdlp

import (
	"fmt"
	"testing"

	"github.com/Synchro/YoutubeDownloader/constants"
	"github.com/stretchr/testify/assert"
)

func TestThatSongIsDownloaded(t *testing.T) {
	yt := New(".")
	err := yt.Download(fmt.Sprintf(constants.YT_VIDEO_URL_FORMAT, "Y2lWEVyO-qE"), "testsample")
	assert.Nil(t, err, "error should be nil")
}
