package main

import (
	"fmt"
	"log/slog"
	"os"
	"strconv"
	"time"

	"github.com/Synchro/common/log"

	"github.com/Synchro/YoutubeSynchroniser/constants"
	"github.com/Synchro/YoutubeSynchroniser/db"
	"github.com/Synchro/YoutubeSynchroniser/models"
	"github.com/Synchro/YoutubeSynchroniser/sdk"
	cc "github.com/Synchro/common/constants"
)

func init() {
	muri := os.Getenv(cc.ENV_MONGODB_URI)
	if muri == "" {
		panic(fmt.Sprintf(cc.MONGO_URI_UNSET, cc.ENV_MONGODB_URI))
	}

	ak := os.Getenv(constants.ENV_APIKEY)
	if ak == "" {
		panic(fmt.Sprintf(constants.APIKEY_UNSET, constants.ENV_APIKEY))
	}

	chId := os.Getenv(constants.ENV_CHANNEL_ID)
	if chId == "" {
		panic(fmt.Sprintf(constants.APIKEY_UNSET, constants.ENV_CHANNEL_ID))
	}

	i := os.Getenv(constants.ENV_CHECK_INTERVAL)
	if in, _ := strconv.Atoi(i); in != 0 {
		constants.GC.CheckInterval = in
	}

	constants.GC.ChannelId = chId
	constants.GC.APIKey = ak
	constants.GC.MongoDBURI = muri

}

func main() {
	logger := log.New(slog.LevelInfo)
	for {
		// Get Youtube Music

		yt, err := sdk.New(constants.GC.APIKey, logger)
		if err != nil {
			panic(fmt.Sprintf("failed to create youtube SDK, %s", err.Error()))
		}

		logger.Info("starting Youtube synchroniser", "configured interval", constants.GC.CheckInterval)

		logger.Info("getting playlists")
		pls, err := yt.GetPlaylists(constants.GC.ChannelId)
		if err != nil {
			logger.Error("error while getting playlists", "err", err.Error())
			time.Sleep(time.Duration(constants.GC.CheckInterval) * time.Minute)
			continue
		}
		logger.Debug("got playlists", "playlists", pls)
		for _, pl := range pls {
			logger.Info("getting tracks for playlist", "playlist name", pl.Title, "playlist Id", pl.Id)
			songs, err := yt.GetPlaylistItems(pl.Id)
			if err != nil {
				//log and continue
				logger.Error("error while getting playlist items", "err", err.Error())
				continue
			}

			plws := models.SongsWithPlaylist{
				Title: pl.Title,
				Songs: songs,
			}
			logger.Info("inserting songs into db", "playlist", pl.Title, "number of songs", len(plws.Songs))
			err = db.InsertUniqueSongs(plws)
			if err != nil {
				logger.Error("error while inserting songs in db", "err", err.Error())
				continue
			}
			logger.Info("songs inserted successfully", "playlist name", pl.Title)
		}
		logger.Info("will check again after the interval", "interval", constants.GC.CheckInterval)
		time.Sleep(time.Duration(constants.GC.CheckInterval) * time.Minute)
	}
}
