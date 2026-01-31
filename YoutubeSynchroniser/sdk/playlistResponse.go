package sdk

import (
	"errors"

	"github.com/Synchro/YoutubeSynchroniser/constants"
	"google.golang.org/api/youtube/v3"
)

type playlistResponse struct {
	chId        string
	data        *youtube.PlaylistListResponse
	cl          *youtube.PlaylistsService
	totalPages  int
	currentIter int
}

func newPlaylistReponse(channelId string, cl *youtube.PlaylistsService) *playlistResponse {
	return &playlistResponse{
		chId:        channelId,
		cl:          cl,
		currentIter: 1,
	}
}

func (pr *playlistResponse) Do() error {
	if pr.totalPages != 0 && pr.currentIter >= pr.totalPages {
		return errors.New(constants.PAGE_ENDED)
	}
	return pr.request()
}

func (pr *playlistResponse) Content() *youtube.PlaylistListResponse {
	return pr.data
}

func (pr *playlistResponse) request() error {
	req := pr.cl.List(constants.YT_PART).ChannelId(pr.chId).MaxResults(100)
	if pr.data != nil {
		// TODO: Check whether NextPageToken is empty in case of last page response
		req.PageToken(pr.data.NextPageToken)
	}
	resp, err := req.Do()
	if err != nil {
		return err
	}

	pr.data = resp
	pr.totalPages = int(pr.data.PageInfo.TotalResults) / int(pr.data.PageInfo.ResultsPerPage)
	if int(pr.data.PageInfo.TotalResults)%int(pr.data.PageInfo.ResultsPerPage) != 0 {
		pr.totalPages++
	}
	pr.currentIter++

	return nil
}
