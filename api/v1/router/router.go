package router

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sqlite/test/api/v1/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
)

type JSON map[string]interface{}

func (j *JSON) toJson() []byte {
	resp, err := json.Marshal(j)
	if err != nil {
		log.Fatal("Error while parsing json")
	}
	return resp
}

func Init() *chi.Mux {

	r := chi.NewRouter()

	//Middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(render.SetContentType(render.ContentTypeJSON))
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
	}))

	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		resp := JSON{
			"msg": "Hola mundo",
		}
		w.Write(resp.toJson())
	})

	//https://www.youtube.com/watch?v=JVzERlysvts&list=RDJVzERlysvts&start_radio=1&ab_channel=GiselleChuwen

	r.Get("/video/{id}", handlers.GetVideoData)
	r.Post("/video/{id}", handlers.DownloadAudio)
	r.Get("/test/{id}", handlers.TestVideo)

	return r
}

func Serve() {

	r := Init()

	fmt.Println("Listenning on http://localhost:8080/")
	log.Fatal(http.ListenAndServe(":8080", r))
}
