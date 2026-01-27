package models

type Config struct {
	MongodbURI    string
	CheckInterval int
	RootDir       string // This is the base directory where all songs will be downloaded
}
