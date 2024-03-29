package router

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"echotom.dev/hellofresh/database"
	"echotom.dev/hellofresh/logger"
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
		logger.Err(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(fmtResponse(err.Error()))
		return
	}

	_, err = database.AddDataCenter(&dc)
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

	json.NewEncoder(w).Encode(dc)
}
