package db

import (
	"context"
	"time"

	"github.com/Synchro/YoutubeSynchroniser/constants"
	"github.com/Synchro/YoutubeSynchroniser/models"
	"github.com/Synchro/common/database"

	cmodels "github.com/Synchro/common/models"
)

func InsertUniqueSongs(pl *models.Playlist) error {
	mcl, err := database.New(constants.GC.MongoDBURI)
	if err != nil {
		return err
	}

	scl := database.NewSongsDataController(mcl)

	for _, pls := range pl.Entries {
		ctx, cancelFunc := context.WithTimeout(context.Background(), time.Duration(5)*time.Minute)
		defer cancelFunc()

		present, err := scl.IsPresent(ctx, pls.URL)
		if err != nil {
			return err
		}

		if present {
			continue
		}

		s := &cmodels.Song{
			VideoURL:     pls.URL,
			Name:         pls.Title,
			PlaylistName: pl.Title,
		}

		err = scl.InsertSong(ctx, s)
		if err != nil {
			return err
		}

	}

	return nil
}
