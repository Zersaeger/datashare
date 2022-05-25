package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type message struct {
	Message string
}
type id struct {
	Id string
}

var messages = make(map[string]message)

func main() {
	router := chi.NewRouter()
	router.Post("/message", createMessage)
	router.Get("/message/{id}", getMessage)
	http.ListenAndServe(":8080", router)

}

func createMessage(writer http.ResponseWriter, request *http.Request) {
	var msg message
	json.NewDecoder(request.Body).Decode(&msg)
	stringId := uuid.NewString()
	messages[stringId] = msg
	fmt.Println(messages)
	id := id{
		Id: stringId,
	}
	text, _ := json.Marshal(id)
	writer.Header().Set("Content-type", "application/json")
	writer.Write(text)

}

func getMessage(writer http.ResponseWriter, request *http.Request) {
	id := chi.URLParam(request, "id")
	message := messages[id]
	text, _ := json.Marshal(message)
	writer.Header().Set("Content-type", "application/json")
	writer.Write(text)
}
