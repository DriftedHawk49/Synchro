package main

import (
	"context"
	"log/slog"
	"strconv"
	"time"

	"github.com/Synchro/Agent/constants"
	"github.com/Synchro/Agent/sdk/ipod"
	gc "github.com/Synchro/common/constants"
	"github.com/Synchro/common/database"
	"github.com/Synchro/common/log"

	"fmt"
	"os"
)

func init() {

}

/*
 * Logic:
 * 1. In a certain interval, check if ipod is in network and reachable. (check by SSH)
 * 2. If ipod is reachable, get all the songs that are not yet synced.
 * 3. One by one, scp those into the ipod predetermined directory.
 * 4. As soon as a song is synced, update that song in database
 */

func init() {
	muri := os.Getenv(gc.ENV_MONGODB_URI)
	if muri == "" {
		panic(fmt.Sprintf(gc.MONGO_URI_UNSET, gc.ENV_MONGODB_URI))
	}

	baseDir := os.Getenv(constants.ENV_BASE_DIR)
	if baseDir == "" {
		panic(fmt.Sprintf(constants.BASE_DIR_UNSET, constants.ENV_BASE_DIR))
	}

	in := os.Getenv(constants.ENV_INTERVAL)
	inc, _ := strconv.Atoi(in)
	if inc != 0 {
		constants.GC.Interval = inc
	}

	ipodAddr := os.Getenv(constants.ENV_IPOD_ADDR)
	if ipodAddr != "" {
		constants.GC.IpodAddr = ipodAddr
	}

	ipodSSHUser := os.Getenv(constants.ENV_IPOD_SSH_USER)
	if ipodSSHUser != "" {
		constants.GC.IpodSSHUser = ipodSSHUser
	}

	ipodSSHKey := os.Getenv(constants.ENV_IPOD_SSH_KEY)
	if ipodSSHKey != "" {
		constants.GC.IpodSSHKey = ipodSSHKey
	}

	ipodDir := os.Getenv(constants.ENV_IPOD_DIR)
	if ipodSSHUser != "" {
		constants.GC.IpodSSHUser = ipodDir
	}

	constants.GC.MongoDBURI = muri
	constants.GC.BaseDir = baseDir

}

func main() {
	logger := log.New(slog.LevelInfo)
	logger.Info("starting agent")
	ia := ipod.New(logger, &ipod.IpodConfig{
		IP:   constants.GC.IpodAddr,
		Dir:  constants.GC.IpodDir,
		User: constants.GC.IpodSSHUser,
		Key:  constants.GC.IpodSSHKey,
	})
	for {
		if !ia.IsInNetwork() {
			logger.Info("ipod is not in network currently, will try after timout", "timeout (in minutes)", constants.GC.Interval)
			time.Sleep(time.Minute * time.Duration(constants.GC.Interval))
			continue
		}

		logger.Info("ipod is in network, getting un synced songs from db")

		mcl, err := database.New(constants.GC.MongoDBURI)
		if err != nil {
			logger.Error("failed to create database connection, will try after sometime", "err", err.Error())
			time.Sleep(time.Minute * time.Duration(constants.GC.Interval))
			continue
		}

		sc := database.NewSongsDataController(mcl)
		filter := make(map[string]any)
		filter["synced"] = false
		songs, err := sc.GetSongs(context.Background(), filter)
		if err != nil {
			logger.Error("failed to get un-synced songs from db, will try after sometime", "err", err.Error())
			time.Sleep(time.Minute * time.Duration(constants.GC.Interval))
			continue
		}

		logger.Info("successfully got un synced songs", "no of unsynced songs", len(songs))

		if len(songs) == 0 {
			logger.Info("nothing to sync right now, will try again after timeout")
			time.Sleep(time.Minute * time.Duration(constants.GC.Interval))
			continue
		}

		payload := make(map[string][]string)
		for _, s := range songs {
			if _, ok := payload[s.PlaylistName]; !ok {
				payload[s.PlaylistName] = make([]string, 0)
			}
			payload[s.PlaylistName] = append(payload[s.PlaylistName], s.LocationOnDisk)
		}

		succeeded, err := ia.SendFiles(payload)
		if err != nil {
			logger.Error("failed to send files to ipod; will retry after some time", "err", err.Error())
			time.Sleep(time.Minute * time.Duration(constants.GC.Interval))
			continue
		}

		for _, suc := range succeeded {
			err = sc.SetSyncStatus(context.Background(), suc)
			if err != nil {
				logger.Error("error while updating sync status for song", "song location", suc, "err", err.Error())
			}
		}

	}

}
