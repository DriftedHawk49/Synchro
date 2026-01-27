package db

import (
	"context"
	"time"

	"github.com/Synchro/YoutubeDownloader/constants"
	"github.com/Synchro/common/database"
	"github.com/Synchro/common/models"
)

func GetSongsToBeDownloaded() ([]models.Song, error) {
	mcl, err := database.New(constants.GC.MongodbURI)
	if err != nil {
		return nil, err
	}

	sc := database.NewSongsDataController(mcl)
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Minute)
	defer cancelFunc()

	filter := make(map[string]any)
	filter["downloaded"] = false
	songs, err := sc.GetSongs(ctx, filter)
	if err != nil {
		return nil, err
	}

	return songs, nil
}
