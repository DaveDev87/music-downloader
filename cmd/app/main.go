package main

import (
	R "sqlite/test/api/v1/router"
)

// Ya no es sobre sqlite ahora es sobre IDv3

// func DownloadAudio(videoID string) {
// 	client := youtube.Client{}

// 	video, err := client.GetVideo(videoID)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	formats := video.Formats.WithAudioChannels()
// 	stream, _, err := client.GetStream(video, &formats[0])
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	file, err := os.Create("video.mp4")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer file.Close()

// 	_, err = io.Copy(file, stream)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

func main() {
	R.Serve()

}
