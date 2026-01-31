package ipod

import (
	"context"
	"io"
	"log/slog"
	"net"
	"os"
	"path"
	"strings"
	"time"

	"github.com/Synchro/Agent/constants"
	sftp "github.com/pkg/sftp/v2"
	"golang.org/x/crypto/ssh"
)

type IpodConfig struct {
	IP   string // IP address of Ipod
	Dir  string // Directory where songs will be placed within playlist named folders
	User string // iPod SSH User
	Key  string // iPod SSH Key
}

type ipod struct {
	*IpodConfig
	logger *slog.Logger
}

func New(logger *slog.Logger, config *IpodConfig) *ipod {
	if config == nil {
		config = &IpodConfig{
			IP:   constants.DEFAULT_IPOD_IP,
			Dir:  constants.DEFAULT_IPOD_DIR,
			User: constants.DEFAULT_IPOD_SSH_USER,
			Key:  constants.DEFAULT_IPOD_SSH_KEY,
		}
	}
	return &ipod{
		IpodConfig: config,
		logger:     logger,
	}
}

// Checks whether ipod is in network or not
func (i *ipod) IsInNetwork() bool {
	addr := net.JoinHostPort(i.IP, "22")
	conn, err := net.DialTimeout("tcp", addr, 2*time.Second)
	if err != nil {
		i.logger.Error("error while connecting with ipod", "err", err.Error())
		return false
	}

	conn.SetReadDeadline(time.Now().Add(2 * time.Second))
	buf := make([]byte, 256)
	_, err = conn.Read(buf)
	if err != nil {
		i.logger.Error("error while reading buffer", "err", err.Error())
		return false
	}
	return true
}

// This takes in a map, where keys are playlist names and values are slices of song paths on disk,
// returns slice of song locations which succeeded in transferring.
//
// Errors out only when a blocking error is encountered, else tries to transfer all files
func (i *ipod) SendFiles(swp map[string][]string) ([]string, error) {
	// Set up SSH config
	succeeded := make([]string, 0)
	config := &ssh.ClientConfig{
		User: i.User,
		Auth: []ssh.AuthMethod{
			ssh.Password(i.Key),
		},
		HostKeyAlgorithms: []string{ssh.KeyAlgoRSA, ssh.KeyAlgoRSASHA256},
		HostKeyCallback:   ssh.InsecureIgnoreHostKey(),
		Timeout:           5 * time.Second,
	}

	addr := net.JoinHostPort(i.IP, "22")

	sshClient, err := ssh.Dial("tcp", addr, config)
	if err != nil {
		i.logger.Error("failed to create ssh client with iPod", "err", err.Error())
		return succeeded, err
	}
	defer sshClient.Close()

	// Create an SFTP client over that SSH connection
	sc, err := sftp.NewClient(context.Background(), sshClient, sftp.CopyStderrTo(os.Stdout))
	if err != nil {
		i.logger.Error("failed to create sftp client with iPod", "err", err.Error())
		return succeeded, err
	}
	defer sc.Close()

	for pl, songs := range swp {
		plp := path.Join(i.Dir, pl)
		if _, err = sc.Stat(plp); err != nil {
			// create directory
			err = sc.Mkdir(plp, os.ModeDir)
			if err != nil {
				return succeeded, err
			}
		}

		for _, s := range songs {
			func() {
				t := strings.Split(s, "/")
				sn := t[len(t)-1]
				fp, err := sc.Create(path.Join(plp, sn))
				if err != nil {
					i.logger.Error("error while creating file on ipod, will try for next song", "filename", sn, "err", err.Error())
					return
				}
				defer fp.Close()
				rsfp, err := os.Open(s)
				if err != nil {
					i.logger.Error("error while opening song file on disk, will try for next song", "song location", s, "err", err.Error())
					return
				}
				defer rsfp.Close()

				wb, err := io.Copy(fp, rsfp)
				if err != nil {
					i.logger.Error("error while copying song on ipod", "err", err.Error())
				}
				succeeded = append(succeeded, s)
				i.logger.Info("successfully copied song", "song location on disk", s, "written bytes", wb)
			}()
		}
	}
	return succeeded, nil
}
