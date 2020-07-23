package router

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"echotom.dev/hellofresh/database"
	"echotom.dev/hellofresh/logger"

	"github.com/gorilla/mux"
)

func getMealByNameHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	n := vars["name"]
	m, err := database.GetMealByName(n)

	if err != nil {
		logger.Err(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(fmtResponse(err.Error()))
		return
	}

	if m.Name == "" {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(fmtResponse(n + " not found"))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(m)
}

func addMealHandler(w http.ResponseWriter, r *http.Request) {
	payload, _ := ioutil.ReadAll(r.Body)
	m := database.Meal{}

	err := json.Unmarshal([]byte(payload), &m)

	if err != nil {
		logger.Err(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(fmtResponse(err.Error()))
		return
	}

	_, err = database.AddMeal(&m)

	if err != nil {
		if strings.Contains(err.Error(), "already exists") {
			w.WriteHeader(http.StatusAlreadyReported)
			json.NewEncoder(w).Encode(fmtResponse(err.Error()))
			return
		}
		logger.Err(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(fmtResponse(err.Error()))
		return
	}

	json.NewEncoder(w).Encode(m)

}

func updateMealHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	n := vars["name"]
	m, err := database.GetMealByName(n)

	if err != nil {
		logger.Err(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(fmtResponse(err.Error()))
		return
	}

	if m.Name == "" {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(fmtResponse(n + " not found"))
		return
	}

	payload, _ := ioutil.ReadAll(r.Body)
	updt := database.Meal{}
	err = json.Unmarshal([]byte(payload), &updt)

	if err != nil {
		logger.Err(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(fmtResponse(err.Error()))
		return
	}

	m.Name = updt.Name
	m.Metadata = updt.Metadata
	err = database.UpdateMeal(&m)

	if err != nil {
		logger.Err(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(fmtResponse(err.Error()))
		return
	}

	json.NewEncoder(w).Encode(m)
}

func deleteMealHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	n := vars["name"]

	err := database.DeleteMeal(n)

	if err != nil {
		if strings.Contains(err.Error(), "not exists") {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(fmtResponse(err.Error()))
			return
		}
		logger.Err(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(fmtResponse(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(fmtResponse(n + " was successfuly deleted"))
}
