package router

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"echotom.dev/hellofresh/database"

	"github.com/gorilla/mux"
)

func getMealByNameHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	n := vars["name"]
	m, err := database.GetMealByName(n)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(fmtResponse(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(m)
}

func createMealHandler(w http.ResponseWriter, r *http.Request) {
	payload, _ := ioutil.ReadAll(r.Body)
	m := database.Meal{}

	err := json.Unmarshal([]byte(payload), &m)

	if err != nil {
		fmt.Println(err, string(payload))
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(fmtResponse(err.Error()))
		return
	}

	_, err = database.AddMeal(&m)

	if err != nil {
		fmt.Println(err)
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
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(fmtResponse(err.Error()))
		return
	}

	if m.Name == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(fmtResponse(n + " not found"))
		return
	}

	payload, _ := ioutil.ReadAll(r.Body)
	updt := database.Meal{}
	err = json.Unmarshal([]byte(payload), &updt)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(fmtResponse(err.Error()))
		return
	}

	m.Name = updt.Name
	m.Metadata = updt.Metadata
	database.UpdateMeal(&m)
	json.NewEncoder(w).Encode(m)
}

func deleteMealHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	n := vars["name"]

	err := database.DeleteMeal(n)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(fmtResponse(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(fmtResponse(err.Error()))
}
