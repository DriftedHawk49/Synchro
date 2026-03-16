package models

type Config struct {
	MongoDBURI    string
	ChannelName   string
	CheckInterval int // In minutes
}
