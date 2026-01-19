# Synchro
This contains a collection of microservices : 
1. Youtube Synchroniser : This will get periodic updates from youtube API and save details in mongodb
2. Youtube Downloader : This will download all the songs from youtube, and save them on disk for synchronisation
3. Agent : This Agent will keep looking for iPod to be in network, and as soon as it is within reach, it will simply sync music library

All three services depend on database for their assigned work and their functionality. They have no direct interdependencies.

## Project Directory Structure
1. common : This will contain all the common packages like constants, logger and database communication essentials
2. YoutubeSynchroniser : This will contain code for youtube synchronisation
3. Agent : This will sync music with iPod.
4. YoutubeDownloader : This will get data from database and download songs on disk and update their records
