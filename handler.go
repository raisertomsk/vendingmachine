package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Handler struct {
	app *App
}

func NewHandler(app *App) *Handler {
	handler := new(Handler)
	handler.app = app
	return handler
}

func (h *Handler) List(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	data, err := json.Marshal(h.app.List())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error marshal products in JSON for List: " + err.Error())
		return
	}
	w.Write(data)
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
}

func (h *Handler) Add(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error reading body data from request: " + err.Error())
		return
	}
	result, err := h.app.Add(ps.ByName("name"), data).AsJSON()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error marshal Error to JSON on Add: " + err.Error())
	}
	w.Write(result)
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
}
