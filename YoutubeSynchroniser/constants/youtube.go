package constants

import (
	"github.com/Synchro/YoutubeSynchroniser/models"
)

const (
	ENV_CHECK_INTERVAL     string = "INTERVAL"
	ENV_CHANNEL_NAME       string = "CHANNEL_NAME"
	DEFAULT_CHECK_INTERVAL int    = 30 // Check every 30 minutes
)

var (
	YT_PART []string      = []string{"contentDetails", "snippet"}
	GC      models.Config = models.Config{
		CheckInterval: DEFAULT_CHECK_INTERVAL,
	}
)

// formats
const (
	YT_CHANNEL_FORMAT string = "https://www.youtube.com/%s/playlists"
)

const (
	FILTER_KEY_VIDEOID string = "videoId"
)

// Error constants
const (
	PAGE_ENDED       string = "all pages traversed" // can be used where request is being made when pagination has ended
	APIKEY_UNSET     string = "Youtube API Key not set, set %s as APIKey for Youtube"
	CHANNEL_ID_UNSET string = "Channel Id not set, set %s as ChannelId for Youtube"
)
