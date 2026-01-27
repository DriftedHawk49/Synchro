package dispatcher

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"path"
	"time"

	"github.com/Synchro/YoutubeDownloader/constants"
	ytdlp "github.com/Synchro/YoutubeDownloader/sdk/ytDlp"
	"github.com/Synchro/common/database"
	"github.com/Synchro/common/models"
)

type jobDispatcher struct {
	noOfWorkers int
	jobChan     chan *models.Song
	mcl         *database.SongsDataController
	logger      *slog.Logger
}

// n : Number of workers to be dispatched;
// buffer : Number of jobs that can stay in buffer when all the workers are busy
func New(n int, buffer int, logger *slog.Logger) (*jobDispatcher, error) {
	mcl, err := database.New(constants.GC.MongodbURI)
	if err != nil {
		return nil, errors.Join(errors.New("error while creating db client"), err)
	}
	return &jobDispatcher{
		noOfWorkers: n,
		jobChan:     make(chan *models.Song, buffer),
		mcl:         database.NewSongsDataController(mcl),
		logger:      logger,
	}, nil
}

// This dispatches the workers to start accepting jobs
func (jd *jobDispatcher) Initialize() {
	for range jd.noOfWorkers {
		go jd.perform()
	}
}

// This closes all dispatcher, kills all workers gracefully
func (jd *jobDispatcher) Stop() {
	close(jd.jobChan)
}

func (jd *jobDispatcher) AddJob(s models.Song) {
	jd.jobChan <- &s
}

func (jd *jobDispatcher) perform() {
	for s := range jd.jobChan {
		err := jd.downloadAndUpdateSong(s)
		if err != nil {
			jd.logger.Error(err.Error())
		}
	}
}

func (jd *jobDispatcher) downloadAndUpdateSong(s *models.Song) error {
	ytd := ytdlp.New(constants.GC.RootDir)
	err := ytd.Download(fmt.Sprintf(constants.YT_VIDEO_URL_FORMAT, s.VideoId), s.PlaylistName)
	if err != nil {
		return fmt.Errorf("error while downloading song, check logs for reason of failure, song title : %s, playlist title : %s", s.Name, s.PlaylistName)
	}

	// updating song
	s.UpdatedAt = time.Now()
	s.Downloaded = true
	s.LocationOnDisk = path.Join(s.PlaylistName, s.Name)

	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Minute)
	defer cancelFunc()
	err = jd.mcl.UpdateSong(ctx, s)
	if err != nil {
		return fmt.Errorf("error while updating db after downloading song, err : %s", err.Error())
	}
	return nil
}
