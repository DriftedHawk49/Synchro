package sdk

import (
	"context"
	"log/slog"

	"github.com/Synchro/YoutubeSynchroniser/constants"
	"github.com/Synchro/YoutubeSynchroniser/models"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

type YoutubeSDK struct {
	svc    *youtube.Service
	logger *slog.Logger
}

func New(apiKey string, logger *slog.Logger) (*YoutubeSDK, error) {
	svc, err := youtube.NewService(context.Background(), option.WithAPIKey(apiKey))
	if err != nil {
		return nil, err
	}

	return &YoutubeSDK{
		svc:    svc,
		logger: logger,
	}, nil
}

func (ys *YoutubeSDK) GetPlaylists(channelId string) ([]*models.Playlist, error) {

	pr := newPlaylistReponse(channelId, youtube.NewPlaylistsService(ys.svc))
	result := make([]*models.Playlist, 0)

	var err error
	for err == nil {
		err = pr.Do()
		if pr.data == nil {
			break
		}
		for _, res := range pr.data.Items {
			result = append(result, &models.Playlist{
				Title: res.Snippet.Title,
				Id:    res.Id,
			})
		}
	}

	if err.Error() != constants.PAGE_ENDED {
		return result, err
	}

	return result, nil

}

func (ys *YoutubeSDK) GetPlaylistItems(playlistId string) ([]*models.Song, error) {
	pr := newPlaylistItemReponse(playlistId, youtube.NewPlaylistItemsService(ys.svc))
	result := make([]*models.Song, 0)

	var err error
	for err == nil {
		err = pr.Do()
		for _, res := range pr.data.Items {
			result = append(result, &models.Song{
				Title:   res.Snippet.Title,
				VideoId: res.Snippet.ResourceId.VideoId,
			})
		}
	}

	if err.Error() != constants.PAGE_ENDED {
		return result, err
	}

	return result, nil
}
