package database

import (
	"context"
	"time"

	"github.com/Synchro/common/constants"
	"github.com/Synchro/common/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type SongsDataController struct {
	cl *mongo.Client
}

func NewSongsDataController(client *mongo.Client) *SongsDataController {
	return &SongsDataController{
		cl: client,
	}
}

func (sdc *SongsDataController) InsertSong(ctx context.Context, s *models.Song) error {
	coll := sdc.cl.Database(constants.DB_NAME).Collection(constants.COLLECTION_NAME)
	ct := time.Now()
	s.CreatedAt = ct
	s.UpdatedAt = ct
	_, err := coll.InsertOne(ctx, s)
	return err
}

func (sdc *SongsDataController) GetSongs(ctx context.Context, filter map[string]any) ([]models.Song, error) {
	coll := sdc.cl.Database(constants.DB_NAME).Collection(constants.COLLECTION_NAME)

	resp := make([]models.Song, 0)

	filterB := bson.M(filter)
	result, err := coll.Find(ctx, filterB)
	if err != nil {
		return resp, err
	}
	defer result.Close(ctx)

	err = result.All(ctx, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil

}

func (sdc *SongsDataController) UpdateSong(ctx context.Context, song *models.Song) error {
	coll := sdc.cl.Database(constants.DB_NAME).Collection(constants.COLLECTION_NAME)

	_, err := coll.ReplaceOne(ctx, bson.M{"videoId": song.VideoId}, song)
	if err != nil {
		return err
	}

	return nil
}

// Delete is not required
