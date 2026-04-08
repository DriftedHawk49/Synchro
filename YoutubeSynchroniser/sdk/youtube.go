package sdk

import (
	"fmt"
	"os/exec"

	"github.com/Synchro/YoutubeSynchroniser/constants"
	"github.com/Synchro/YoutubeSynchroniser/models"
)

type YoutubeSDK struct {
}

func New() *YoutubeSDK {
	return &YoutubeSDK{}
}

func (ys *YoutubeSDK) GetPlaylists(channelName string) ([]*models.PlaylistEntry, error) {

	rawOutput, err := ytdlCommand(fmt.Sprintf(constants.YT_CHANNEL_FORMAT, channelName)).CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("error while getting playlists, output : %s, err : %s", string(rawOutput), err.Error())
	}

	ch, err := models.ParseChannel(rawOutput)
	if err != nil {
		return nil, fmt.Errorf("error while parsing channel data from raw form, err : %s", err.Error())
	}

	return ch.Entries, nil

}

func (ys *YoutubeSDK) GetPlaylistItems(playlistURL string) (*models.Playlist, error) {
	rawOutput, err := ytdlCommand(playlistURL).CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("error while getting playlistItems, output : %s, err : %s", string(rawOutput), err.Error())
	}

	pls, err := models.ParsePlaylist(rawOutput)
	if err != nil {
		return nil, fmt.Errorf("error while parsing playlist data from raw form, err : %s", err.Error())
	}

	return pls, nil
}

func ytdlCommand(url string) *exec.Cmd {
	return exec.Command("yt-dlp", "--flat-playlist", "-J", fmt.Sprintf("%s", url))
}
