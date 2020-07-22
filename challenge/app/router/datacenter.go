package router

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"echotom.dev/hellofresh/database"
)

func getAllDataCentersHandler(w http.ResponseWriter, r *http.Request) {
	dcs := database.GetAllDataCenters()
	if len(dcs) < 1 {
		json.NewEncoder(w).Encode(make([]string, 0))
		return
	}

	json.NewEncoder(w).Encode(dcs)
}

func addDataCenterHandler(w http.ResponseWriter, r *http.Request) {
	payload, _ := ioutil.ReadAll(r.Body)
	dc := database.DataCenter{}
	err := json.Unmarshal([]byte(payload), &dc)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	_, err = database.AddDataCenter(&dc)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	json.NewEncoder(w).Encode(dc)
}
