package db

import (
	"context"
	"time"

	"github.com/Synchro/YoutubeSynchroniser/constants"
	"github.com/Synchro/YoutubeSynchroniser/models"
	"github.com/Synchro/common/database"

	cmodels "github.com/Synchro/common/models"
)

func InsertUniqueSongs(plws models.SongsWithPlaylist) error {
	mcl, err := database.New(constants.GC.MongoDBURI)
	if err != nil {
		return err
	}

	scl := database.NewSongsDataController(mcl)

	for _, pls := range plws.Songs {
		ctx, cancelFunc := context.WithTimeout(context.Background(), time.Duration(5)*time.Minute)
		defer cancelFunc()

		present, err := scl.IsPresent(ctx, pls.VideoId)
		if err != nil {
			return err
		}

		if present {
			continue
		}

		s := &cmodels.Song{
			VideoId:      pls.VideoId,
			Name:         pls.Title,
			PlaylistName: plws.Title,
		}

		err = scl.InsertSong(ctx, s)
		if err != nil {
			return err
		}

	}

	return nil
}
