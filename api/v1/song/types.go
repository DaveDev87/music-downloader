package song

import (
	"encoding/json"
	"net/http"
)

// JSON type
type JSON map[string]interface{}

func (j JSON) toJson(w http.ResponseWriter) []byte {
	data, err := json.Marshal(j)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	return data
}
