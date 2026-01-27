package constants

import "github.com/Synchro/YoutubeDownloader/models"

const (
	YT_VIDEO_URL_FORMAT    string = "https://youtube.com/watch?v=%s"
	ENV_ROOT_DIR           string = "ROOTDIR" // This is the base directory where all songs will be downloaded
	ENV_CHECK_INTERVAL     string = "INTERVAL"
	DEFAULT_CHECK_INTERVAL int    = 30
	DEFAULT_ROOT_DIR       string = "/opt/songs"
)

var (
	GC models.Config = models.Config{
		CheckInterval: DEFAULT_CHECK_INTERVAL,
		RootDir:       DEFAULT_ROOT_DIR,
	}
)
