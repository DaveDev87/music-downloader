package song

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	"github.com/olahol/melody"
)

func Init() *chi.Mux {

	r := chi.NewRouter()
	m := melody.New()

	// #region Middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(render.SetContentType(render.ContentTypeJSON))
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
	}))
	// #endregion

	// #region WebSockets
	r.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		m.HandleRequest(w, r)
	})

	m.HandleConnect(func(s *melody.Session) {

	})

	m.HandleDisconnect(func(s *melody.Session) {
		fmt.Println("closed")
	})

	m.HandleMessage(DownloadWithProgression(m))

	// #endregion

	// #region API routes
	r.Get("/video/{id}", GetVideoData)
	r.Post("/video/{id}", DownloadAudio)
	r.Get("/test/{id}", TestVideo)
	// #endregion

	return r
}

func Serve() {

	r := Init()

	fmt.Println("Listenning on http://localhost:8080/")
	log.Fatal(http.ListenAndServe(":8080", r))
}
