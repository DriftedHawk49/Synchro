package ytdlp

import (
	"fmt"
	"os/exec"
	"path/filepath"
)

type ytDownloader struct {
	rootDir string // base directory where all songs will reside
}

func New(rootDir string) *ytDownloader {
	return &ytDownloader{
		rootDir: rootDir,
	}
}

// This function downloads the song and places it in the folder named playlistTitle
func (d *ytDownloader) Download(url string, playlistTitle string) error {
	output, err := exec.Command("yt-dlp", url, "--paths", filepath.Join(d.rootDir, playlistTitle), "-o", "%(id)s.%(ext)s", "-f", "ba", "-x", "--audio-format", "mp3", "--embed-thumbnail", "--embed-metadata").CombinedOutput()
	fmt.Println(string(output))
	return err
}
