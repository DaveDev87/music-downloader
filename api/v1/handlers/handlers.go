package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/kkdai/youtube/v2"
)

// Represents a json like type, convenient to easy marshal any structure before serving it.
type JSON map[string]interface{}

// Marshals a givin type JSON into []bute.
func (j JSON) toJson(w http.ResponseWriter) []byte {
	data, err := json.Marshal(j)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	return data
}

type Progress struct {
	TotalSize int64
	BytesRead int64
}

func (pr *Progress) Write(p []byte) (n int, err error) {
	n, err = len(p), nil
	pr.BytesRead += int64(n)
	pr.Print()
	return
}

func (pr *Progress) Print() {
	if pr.BytesRead == pr.TotalSize {
		fmt.Println("DONE!")
		return
	}

	fmt.Printf("File upload in progress: %d\n", pr.BytesRead)
}

func TestVideo(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	client := youtube.Client{Debug: true}
	video, err := client.GetVideo(id)
	if err != nil {
		log.Fatal(err)
	}

	videoData := JSON{
		"Duration": video.Duration.String(),
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
	stream, _, err := client.GetStream(video, &formats[2])
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

	_, err = io.Copy(file, stream)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Disposition", "attachment; filename=\"test\"")
	w.Header().Set("Content-Type", r.Header.Get("Content-Type"))

	http.ServeFile(w, r, fmt.Sprintf("./tmp/%d.mp4", uName))
}
