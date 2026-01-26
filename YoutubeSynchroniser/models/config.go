package models

type Config struct {
	APIKey        string
	MongoDBURI    string
	ChannelId     string
	CheckInterval int // In minutes
}
