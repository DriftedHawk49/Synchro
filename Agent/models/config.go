package models

type Config struct {
	MongoDBURI  string
	Interval    int
	BaseDir     string
	IpodAddr    string
	IpodDir     string
	IpodSSHUser string
	IpodSSHKey  string
}
