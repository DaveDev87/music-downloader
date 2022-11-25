package song

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	"github.com/olahol/melody"
)

type GopherInfo struct {
	ID, X, Y string
}

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
		ss, _ := m.Sessions()

		for _, o := range ss {
			value, exists := o.Get("info")

			if !exists {
				continue
			}

			info := value.(*GopherInfo)

			s.Write([]byte("set " + info.ID + " " + info.X + " " + info.Y))
		}

		id := uuid.NewString()
		s.Set("info", &GopherInfo{id, "0", "0"})

		s.Write([]byte("iam" + id))
	})

	m.HandleDisconnect(func(s *melody.Session) {
		value, exists := s.Get("info")

		if !exists {
			return
		}

		info := value.(*GopherInfo)

		m.BroadcastOthers([]byte("dis "+info.ID), s)
	})

	m.HandleMessage(func(s *melody.Session, msg []byte) {
		p := strings.Split(string(msg), " ")
		value, exists := s.Get("info")

		if len(p) != 2 || !exists {
			return
		}

		info := value.(*GopherInfo)
		info.X = p[0]
		info.Y = p[1]

		m.BroadcastOthers([]byte("set "+info.ID+" "+info.X+" "+info.Y), s)
	})

	// #endregion

	// #region API routes
	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		resp := JSON{
			"msg": "Hola mundo",
		}
		w.Write(resp.toJson(w))
	})
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
