package main

import (
	"fmt"
	"log/slog"

	"os"
	"os/exec"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/Synchro/YoutubeDownloader/constants"
	"github.com/Synchro/YoutubeDownloader/controllers/dispatcher"
	"github.com/Synchro/YoutubeDownloader/db"
	gc "github.com/Synchro/common/constants"
	"github.com/Synchro/common/log"
)

func init() {

	// check if yt dlp is present in environment, if it is not, then panic
	_, err := exec.Command("yt-dlp", "--version").CombinedOutput()
	if err != nil {
		panic(fmt.Sprintf("yt-dlp not found in path, it is required in path. err : %s", err.Error()))
	}

	muri := os.Getenv(gc.ENV_MONGODB_URI)
	if muri == "" {
		panic(fmt.Sprintf(gc.MONGO_URI_UNSET, gc.ENV_MONGODB_URI))
	}
	constants.GC.MongodbURI = muri

	rd := os.Getenv(constants.ENV_ROOT_DIR)
	if rd != "" {
		constants.GC.RootDir = rd
	}

	intr := os.Getenv(constants.ENV_CHECK_INTERVAL)
	if i, _ := strconv.Atoi(intr); i != 0 {
		constants.GC.CheckInterval = i
	}

}

func main() {

	/*
		 * Logic :
			* within interval, check if there are any un downloaded songs
			* send those un downloaded songs to job dispatcher
			* Each job should first download the song, if successful, then update the db.
	*/
	logger := log.New(slog.LevelInfo)
	sigChan := make(chan os.Signal, 1)

	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	tck := time.NewTicker(time.Second)
	jd, err := dispatcher.New(100, 50, logger)
	if err != nil {
		panic(err.Error())
	}

	jd.Initialize()
	logger.Info("starting youtube downloader", "config", constants.GC)

	for {
		select {
		case <-tck.C:
			tck.Reset(time.Duration(constants.GC.CheckInterval) * time.Minute)
			logger.Info("checking for any songs to be downloaded")
			// Check for un downloaded songs
			songs, err := db.GetSongsToBeDownloaded()
			if err != nil {
				logger.Error("error while getting songs to be downloaded", "err", err.Error())
				continue
			}

			logger.Info("total songs to be downloaded", "songs", len(songs))

			for _, s := range songs {
				jd.AddJob(s)
			}

			logger.Info("dispatched all songs for downloading")

		case <-sigChan:
			jd.Stop()
			tck.Stop()
			return
		}
	}

}
