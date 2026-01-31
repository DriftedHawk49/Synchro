package constants

import "github.com/Synchro/Agent/models"

const (
	ENV_BASE_DIR      string = "BASE_DIR" // directory where songs are present
	ENV_INTERVAL      string = "INTERVAL"
	ENV_IPOD_DIR      string = "IPOD_DIR" // Directory on ipod where to scp songs
	ENV_IPOD_ADDR     string = "IPOD_ADDR"
	ENV_IPOD_SSH_USER string = "IPOD_SSH_USER"
	ENV_IPOD_SSH_KEY  string = "IPOD_SSH_KEY"

	DEFAULT_INTERVAL      int    = 1 // check to sync every minute
	DEFAULT_IPOD_DIR      string = "/var/mobile/Applications/DDA5940E-925A-4917-9362-714AB8A206D4/Documents/TuneShell/Music/Downloads"
	DEFAULT_IPOD_IP       string = "192.168.1.200"
	DEFAULT_IPOD_SSH_KEY  string = "alpine"
	DEFAULT_IPOD_SSH_USER string = "root"
)

var (
	GC models.Config = models.Config{
		Interval:    DEFAULT_INTERVAL,
		IpodDir:     DEFAULT_IPOD_DIR,
		IpodAddr:    DEFAULT_IPOD_IP,
		IpodSSHUser: DEFAULT_IPOD_SSH_USER,
		IpodSSHKey:  DEFAULT_IPOD_SSH_KEY,
	}
)

const (
	IPOD_ADDR_UNSET string = "ipod ip address unset, set %s as iPod IP address"
	BASE_DIR_UNSET  string = "base directory unset, set %s as root directory where songs are present"
)
