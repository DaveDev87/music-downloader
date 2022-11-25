package song

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/kkdai/youtube/v2"

	progression "sqlite/test/internal"
)

func TestVideo(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	client := youtube.Client{Debug: true}
	video, err := client.GetVideo(id)
	if err != nil {
		log.Fatal(err)
		return
	}
	formats := video.Formats.WithAudioChannels()

	_, size, err := client.GetStream(video, &formats[2])
	if err != nil {
		log.Fatal(err)
	}

	videoData := JSON{
		"size": size,
	}

	w.Write(videoData.toJson(w))
}

func GetVideoData(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	client := youtube.Client{Debug: true}
	video, err := client.GetVideo(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	videoData := JSON{
		"Id":         video.ID,
		"Title":      video.Title,
		"Author":     video.Author,
		"Duration":   video.Duration.String(),
		"Thumbnails": video.Thumbnails,
	}

	w.Write(videoData.toJson(w))
}

func DownloadAudio(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")
	uName := time.Now().UnixNano()

	client := youtube.Client{}
	video, err := client.GetVideo(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	formats := video.Formats.WithAudioChannels()
	stream, size, err := client.GetStream(video, &formats[2])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	file, err := os.Create(fmt.Sprintf("./tmp/%d.mp4", uName))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	done := make(chan int64)

	go progression.ProgressFile(done, fmt.Sprintf("./tmp/%d.mp4", uName), size)

	n, err := io.Copy(file, stream)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	done <- n

	w.Header().Set("Content-Disposition", "attachment; filename=\"test\"")
	w.Header().Set("Content-Type", r.Header.Get("Content-Type"))

	http.ServeFile(w, r, fmt.Sprintf("./tmp/%d.mp4", uName))
}
