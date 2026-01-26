package sdk

import (
	"errors"

	"github.com/Synchro/YoutubeSynchroniser/constants"
	"google.golang.org/api/youtube/v3"
)

type playlistItemResponse struct {
	plId        string
	data        *youtube.PlaylistItemListResponse
	cl          *youtube.PlaylistItemsService
	totalPages  int
	currentIter int
}

func newPlaylistItemReponse(playlistId string, cl *youtube.PlaylistItemsService) *playlistItemResponse {
	return &playlistItemResponse{
		plId:        playlistId,
		cl:          cl,
		currentIter: 1,
	}
}

func (pr *playlistItemResponse) Do() error {
	if pr.totalPages != 0 && pr.currentIter >= pr.totalPages {
		return errors.New(constants.PAGE_ENDED)
	}
	return pr.request()
}

func (pr *playlistItemResponse) Content() *youtube.PlaylistItemListResponse {
	return pr.data
}

func (pr *playlistItemResponse) request() error {
	req := pr.cl.List(constants.YT_PART).PlaylistId(pr.plId)
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
