package util

import (
	"encoding/json"
	"log"
	"net/http"
	"reflect"
	"strconv"

	"github.com/ImranZahoor/blog-api/internal/models"
)

func JsonResponse(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(int(code))
	_, _ = w.Write(response)
}
func ToJSONResponse(w http.ResponseWriter, payload interface{}) {
	statusCode := 200

	if reflect.ValueOf(payload).Kind() == reflect.Struct {
		statusCode = int(reflect.ValueOf(payload).Field(2).Int())
	}
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(int(statusCode))
	byteWritten, err := w.Write(response)
	if err != nil {
		http.Error(w, "server error", http.StatusUnprocessableEntity)
	}
	log.Printf("Write %d bytes ", byteWritten)
}

func ToUUID(id string) (models.Uuid, error) {

	intId, err := strconv.Atoi(id)
	if err != nil {
		return 0, err
	}
	return models.Uuid(intId), nil
}
